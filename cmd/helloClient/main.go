package main

// Use tcpdump to create a test file
// tcpdump -w test.pcap
// or use the example above for writing pcap files

import (
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "log"
    "github.com/google/gopacket/layers"
    "time"
    "net"
    "encoding/binary"
    "math"
    "sync"
    "sync/atomic"
    "flag"
    "fmt"
    "runtime"
    "os/exec"
    "io"
    "os"
    "os/signal"
)

type IpPack struct {
    arrivals []time.Time
}

type SimplePacket struct {
    ip            uint32
    contentType   byte
    handshakeType byte
    arrival       time.Time
}

type IpGroup struct {
    id            int
    ipPackMap     map[uint32]IpPack
    ipAppMap      map[uint32]bool
    ipBannedMap   map[uint32]bool
    simplePackets chan SimplePacket
    lock          sync.RWMutex
}

var (
    //pcapFile         = "/Users/kiyon/Downloads/pkt.pcap"
    filePath         string
    handle           *pcap.Handle
    err              error
    bannedLimitCount int
    bannedInterval   int
    limit            float64
    device           string
    ipsetName        string
    ipsetTimeout     int
    snapshot_len     int
    promiscuous      bool
    timeout          int
    cpuCount         int
    logFileName      string

    ethLayer        layers.Ethernet
    ipLayer         layers.IPv4
    tcpLayer        layers.TCP
    tlsLayer        layers.TLS
    parser          *gopacket.DecodingLayerParser
    foundLayerTypes []gopacket.LayerType

    handshake       byte = 22
    applicationData byte = 23
    helloClient     byte = 01

    ipGroups    []IpGroup
    ipGroupSize int
    groupSize   uint32

    wg sync.WaitGroup

    packetCount      uint64
    packetHandleTime uint64
    pps              uint64

    totalIpBanned uint32
)

func init() {
    parseFlag()

    runtime.GOMAXPROCS(cpuCount)

    initIpGroups()
}

func main() {
    start := time.Now()

    setupPcapHandle()
    defer handle.Close()

    setupLayerParser()

    // Loop through packets in file
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    packetSource.NoCopy = true
    packetSource.Lazy = true

    log.Println("开始处理...\n")

    ticker := time.NewTicker(time.Millisecond * time.Duration(bannedInterval))
    showPacketTicker := time.NewTicker(time.Second * 1)

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, os.Kill)

    for {
        select {
        case packet := <-packetSource.Packets():
            if packet != nil {
                atomic.AddUint64(&packetCount, 1)
                atomic.AddUint64(&pps, 1)
                now := time.Now()
                handlePacket(packet)
                interval := uint64(time.Since(now).Nanoseconds())
                atomic.AddUint64(&packetHandleTime, interval)
            } else {
                ticker.Stop()
                goto DONE
            }
        case <-ticker.C:
            doBanAndClean()
        case <-showPacketTicker.C:
            showPps()
        case <-c:
            goto DONE
        }
    }

DONE:

    for i := 0; i < ipGroupSize; i++ {
        close(ipGroups[i].simplePackets)
    }

    doBanAndClean()

    wg.Wait()

    showHandResult(start)
}

func initIpGroups() {
    setGroupSize()

    ipGroups = make([]IpGroup, ipGroupSize)
    for i := 0; i < ipGroupSize; i++ {
        ipGroups[i] = IpGroup{
            i,
            make(map[uint32]IpPack),
            make(map[uint32]bool),
            make(map[uint32]bool),
            make(chan SimplePacket, ipGroupSize),
            sync.RWMutex{},
        }
        wg.Add(1)
        go func(i int) {
            for simplePacket := range (ipGroups[i].simplePackets) {
                handleIp(simplePacket, ipGroups[i].id)
            }
            wg.Done()
        }(i)
    }
    groupSize = uint32(uint64(1<<32) / uint64(ipGroupSize))
}

func setupPcapHandle() {
    if filePath == "" {
        log.Println("抓取", device, "网卡设备")
        handle, err = pcap.OpenLive(device, int32(snapshot_len), promiscuous, time.Duration(timeout))
    } else {
        log.Println("打开", filePath, "文件获取数据源模拟网卡设备")
        handle, err = pcap.OpenOffline(filePath)
    }
    if err != nil {
        log.Fatal(err)
    }

    handle.SetDirection(pcap.DirectionIn)
}

func setupLayerParser() {
    parser = gopacket.NewDecodingLayerParser(
        layers.LayerTypeEthernet,
        &ethLayer,
        &ipLayer,
        &tcpLayer,
        &tlsLayer,
    )
    foundLayerTypes = []gopacket.LayerType{}
}

func parseFlag() {
    flag.IntVar(&bannedLimitCount, "t", 20, "threshold hello包封禁阈值，和limit组合使用")
    flag.IntVar(&bannedInterval, "i", 300, "IP封禁操作间隔时间，单位毫秒")
    flag.Float64Var(&limit, "l", 10, "limit hello包限制时间，和阈值组合使用，单位秒")
    flag.IntVar(&snapshot_len, "len", 1500, "snapshot_len 读取每个包的最大长度")
    flag.BoolVar(&promiscuous, "p", false, "promiscuous mode 网卡是否使用混合模式，混合模式获取所有经过网卡的包")
    flag.IntVar(&timeout, "read-timeout", 1, "read timeout 读取网卡数据超时时间，单位秒")
    flag.StringVar(&device, "d", "en1", "device 网卡名")
    flag.StringVar(&filePath, "f", "", "文件路径")
    flag.StringVar(&logFileName, "log", "/tmp/ssl_attack_mitigation.log", "日志文件路径")
    flag.StringVar(&ipsetName, "ipset", "", "ipset名称")
    flag.IntVar(&cpuCount, "cpu", 4, "设置cpu数量")
    flag.IntVar(&ipsetTimeout, "ipset-timeout", 86400, "ipset timeout 封禁时间，单位秒")

    flag.Parse()

    if ipsetName == "" {
        log.Fatal("请输入 ipset 名称")
    }
}

func showPps() {
    log.Println("pps:", pps)
    atomic.StoreUint64(&pps, 0)
}

func showHandResult(start time.Time) {
    fmt.Println()
    log.Println("总包数：", packetCount)
    log.Println("包处理总时间：", packetHandleTime)
    avgTime := float64(packetHandleTime) / float64(packetCount)
    pps := 1e9 / avgTime
    log.Printf("数据包平均处理时间%.2fns，pps：%.2f\n\n", avgTime, pps)

    log.Printf("%.0fs间隔，%d个hello包阈值限制条件下封禁ip数【%d】\n", limit, bannedLimitCount, totalIpBanned)

    elapsed := time.Since(start)
    log.Println("运行时间: ", elapsed)
}

// setGroupSize 设置使用的cpu数为2的幂
func setGroupSize() {
    numCpu := float64(cpuCount)
    maxPower := float64(int(math.Log2(numCpu)))

    ipGroupSize = int(math.Pow(2.0, maxPower))
    // IP分组数至少为2
    if ipGroupSize <= 1 {
        ipGroupSize = 2
    }
}

func handlePacket(packet gopacket.Packet) {
    err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
    if err == nil {
        var SrcIP net.IP
        needHandle := false
        var contentType byte
        var handshakeType byte
        for _, layerType := range foundLayerTypes {
            if layerType == layers.LayerTypeIPv4 {
                SrcIP = ipLayer.SrcIP
            }
            if layerType == layers.LayerTypeTLS {
                if len(tlsLayer.Contents) > 5 {
                    contentType = tlsLayer.Contents[0]
                    handshakeType = tlsLayer.Contents[5]
                    needHandle = true
                }
            }
        }
        if len(SrcIP) > 0 && needHandle {
            ip := ip2int(SrcIP)
            simplePacket := SimplePacket{
                ip,
                contentType,
                handshakeType,
                packet.Metadata().Timestamp,
            }
            ipGroups[ip/groupSize].simplePackets <- simplePacket
        }
    }
}

func handleIp(simplePacket SimplePacket, groupId int) {
    ipGroups[groupId].lock.Lock()
    defer ipGroups[groupId].lock.Unlock()

    if _, ok := ipGroups[groupId].ipBannedMap[simplePacket.ip]; ok {
        return
    }

    if simplePacket.contentType == handshake && simplePacket.handshakeType == helloClient {
        handleHelloClientHandshake(simplePacket.ip, simplePacket.arrival, groupId)
    }
    if simplePacket.contentType == applicationData {
        handleApplicationData(simplePacket.ip, groupId)
    }
}

func ip2int(ip net.IP) uint32 {
    if len(ip) == 16 {
        return binary.BigEndian.Uint32(ip[12:16])
    }
    return binary.BigEndian.Uint32(ip)
}

func handleHelloClientHandshake(ip uint32, arrival time.Time, groupId int) {
    if ipPack, ok := ipGroups[groupId].ipPackMap[ip]; ok {

        removedCount := getRemovedCount(ipPack, arrival)

        ipPack.arrivals = append(ipPack.arrivals[removedCount:], arrival)
        if len(ipPack.arrivals) >= bannedLimitCount {
            ipGroups[groupId].ipBannedMap[ip] = true
            delete(ipGroups[groupId].ipPackMap, ip)
        } else {
            ipGroups[groupId].ipPackMap[ip] = ipPack
        }
    } else {
        arrivals := make([]time.Time, bannedLimitCount*2)
        arrivals = append(arrivals, arrival)
        ipGroups[groupId].ipPackMap[ip] = IpPack{arrivals,}
    }
}

func getRemovedCount(ipPack IpPack, virtualNow time.Time) int {
    ipCount := len(ipPack.arrivals)
    removedCount := 0
    for i := 0; i < ipCount; i++ {
        oldArrival := ipPack.arrivals[i]
        timeInterval := virtualNow.Sub(oldArrival).Seconds()
        if timeInterval >= limit {
            removedCount++
        } else {
            break
        }
    }
    return removedCount
}

func handleApplicationData(ip uint32, groupId int) {
    delete(ipGroups[groupId].ipPackMap, ip)
}

func doBanAndClean() {
    for i := 0; i < ipGroupSize; i++ {
        ipGroups[i].lock.Lock()
        ipBannedCount := uint32(len(ipGroups[i].ipBannedMap))
        totalIpBanned += ipBannedCount

        j := 0
        ips := make([]string, ipBannedCount, ipBannedCount)
        for ip := range (ipGroups[i].ipBannedMap) {
            ips[j] = int2IpStr(ip)
            j++
        }
        for ip := range (ipGroups[i].ipBannedMap) {
            delete(ipGroups[i].ipBannedMap, ip)
        }

        wg.Add(1)
        go banIps(ips)

        for _, ipPack := range (ipGroups[i].ipPackMap) {
            removedCount := getRemovedCount(ipPack, time.Now())
            ipPack.arrivals = ipPack.arrivals[removedCount:]
        }
        ipGroups[i].lock.Unlock()
    }
}

func banIps(ips []string) {
    //log.Println("封禁IP", len(ips))
    ipset := exec.Command("ipset", "-!", "restore")

    r, w := io.Pipe()

    ipset.Stdin = r
    ipset.Start()

    for _, ip := range (ips) {
        fmt.Fprintf(w, "add %s %s timeout %d\n", ipsetName, ip, ipsetTimeout)
    }

    w.Close()

    ipset.Wait()

    f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
    defer f.Close()
    if err == nil && len(ips) > 0 {
        fmt.Fprintf(f, "%s: 封禁%d个ip\n", time.Now().Format("2006-01-02 15:04:05"), len(ips))
    }

    wg.Done()
}

func int2IpStr(ip uint32) string {
    return fmt.Sprintf("%d.%d.%d.%d",
        byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}
