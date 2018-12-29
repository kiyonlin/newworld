package main

import (
    "net"
    "fmt"
    "log"
    "flag"
    "time"
    randomdata "github.com/pallinder/go-randomdata"
    pb "github.com/kiyonlin/newworld/mobsync"
    "io"
    "crypto/md5"
    "encoding/hex"
    "os"
    "github.com/golang/protobuf/proto"
)

var (
    n        = flag.Int("n", 1, "the node id")
    nodeId   int32
    sessions = make(map[string]time.Time)
)

func main() {
    flag.Parse()
    nodeId = int32(*n)

    tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:5000")
    checkError(err, "ResolveTCPAddr")
    conn, err := net.DialTCP("tcp", nil, tcpAddr)

    if err != nil {
        panic(err)
    }

    go start(conn)

    accept(conn)
}

// start build a connection to master and wait for messages to be sent to master
func start(conn *net.TCPConn) {
    session := getSession()
    message := &pb.Message{
        Action:        pb.Message_CONNECT,
        NodeId:        nodeId,
        SiteId:        1,
        Session:       session,
        IpBlackList:   []string{randomdata.IpV4Address(), randomdata.IpV4Address()},
        UuidBlackList: []string{randomdata.MacAddress(), randomdata.MacAddress()},
        UuidWhiteList: []string{randomdata.MacAddress(), randomdata.MacAddress()},
    }

    sessions[session] = time.Now()

    buffer, _ := proto.Marshal(message)
    _, err := conn.Write(buffer)
    //fmt.Println(lens)
    if err != nil {
        fmt.Println(err.Error())
        conn.Close()
        return
    }

    for {
        ticker := time.NewTicker(time.Millisecond * 500)
        select {
        case <-ticker.C:
            session := getSession()
            message := &pb.Message{
                Action:        pb.Message_REPORT,
                NodeId:        nodeId,
                SiteId:        1,
                Session:       session,
                IpBlackList:   []string{randomdata.IpV4Address(), randomdata.IpV4Address()},
                UuidBlackList: []string{randomdata.MacAddress(), randomdata.MacAddress()},
                UuidWhiteList: []string{randomdata.MacAddress(), randomdata.MacAddress()},
            }
            sessions[session] = time.Now()

            if send(conn, message) == false {
                break
            }
        }
    }
}

// accept wait to accept the task dispatched from master
func accept(conn *net.TCPConn) {
    for {
        //开始客户端轮训
        buf := make([]byte, 1024)
        for {
            length, err := conn.Read(buf)
            if checkError(err, "Connection") == false {
                conn.Close()
                fmt.Println("Server is dead ...ByeBye")
                os.Exit(0)
            }
            message := &pb.Message{}

            err = proto.Unmarshal(buf[0:length], message)

            if err == io.EOF {
                log.Println("⚠️ 收到服务端的结束信号")
                break //如果收到结束信号，则退出“接收循环”，结束客户端程序
            }

            if err != nil {
                log.Println("接收数据出错:", err)
            }

            switch message.Action {
            case pb.Message_REPLY:
                log.Println("收到主控回复", message.Session, "消耗", time.Since(sessions[message.Session]))
                delete(sessions, message.Session)
            case pb.Message_DISPATCH:
                log.Println("收到主控下发需要同步的任务")
                reply := &pb.Message{Action: pb.Message_REPLY, NodeId: nodeId, Session: message.Session}
                send(conn, reply)
            default:
                log.Println("invalid action")
            }
        }
    }
}

func checkError(err error, info string) (res bool) {
    if err != nil {
        fmt.Println(info + "  " + err.Error())
        return false
    }
    return true
}

// getSession returns a random session id for a message
func getSession() string {
    signByte := []byte(time.Now().String())
    hash := md5.New()
    hash.Write(signByte)
    return hex.EncodeToString(hash.Sum(nil))
}

func send(conn *net.TCPConn, message *pb.Message) bool {
    buffer, _ := proto.Marshal(message)
    _, err := conn.Write(buffer)
    if err != nil {
        fmt.Println(err.Error())
        conn.Close()
        return false
    }
    return true
}
