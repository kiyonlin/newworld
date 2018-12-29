package main

import (
    "net"
    "fmt"
    "os"
    "time"
    "github.com/gogather/com/log"
    "flag"
    "io"
    "github.com/kiyonlin/newworld/xwaf-mobsync/mobsync"
)

var (
    n            = flag.Int("n", 1, "the node id")
    nodeId       int32
    sessions     = make(map[string]time.Time)
    xwafMessages chan *mobsync.Message
)

func init() {
    xwafMessages = make(chan *mobsync.Message, 128)
}

func main() {
    flag.Parse()
    nodeId = int32(*n)

    tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:12345")
    checkError(err, "ResolveTCPAddr")
    c, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err, "DialTCP")
    go keepWithMaster(c)
    //启动客户端发送线程

    go keepWithXwaf()

    go onXwafMessageReceived(xwafMessages)
    onMasterMessageReceived(c)
}
func onMasterMessageReceived(conn *net.TCPConn) {
    //开始客户端轮训
    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        if checkError(err, "与主控失去连接") == false {
            conn.Close()
            os.Exit(0)
        }
        message := mobsync.BufToMessage(buf[0:n])

        switch message.Action {
        case mobsync.Message_REPLY:
            log.Println("收到主控回复", message.Session, "消耗", time.Since(sessions[message.Session]))
            delete(sessions, message.Session)
        case mobsync.Message_DISPATCH:
            log.Println("收到主控下发来自节点", message.NodeId, "需要同步的任务，进行回复")
            reply := &mobsync.Message{Action: mobsync.Message_REPLY, NodeId: nodeId, Session: message.Session}
            send(conn, reply)
        default:
            log.Println("错误动作:", message.Action, "session为", message.Session)
        }
    }
}

func keepWithMaster(conn *net.TCPConn) {
    session := mobsync.GetSessionId()
    connectToMaster := &mobsync.Message{
        Action:  mobsync.Message_CONNECT,
        NodeId:  nodeId,
        Session: session,
    }
    sessions[session] = time.Now()
    send(conn, connectToMaster)

    for {
        select {
        case message := <-xwafMessages:
            log.Println("上报消息给主控...")

            session := mobsync.GetSessionId()
            message.Session = session
            sessions[session] = time.Now()
            send(conn, message)
        }
    }
}

func keepWithXwaf() {
    sock := "/tmp/xwaf.sock"
    prepareUnixSock(sock)
    l, err := net.ListenUnix("unix", &net.UnixAddr{sock, "unix"})
    if err != nil {
        panic(err)
    }

    for {
        unixConn, err := l.AcceptUnix()
        if err != nil {
            panic(err)
        }

        go unixPipe(unixConn)
    }
}

func prepareUnixSock(sock string) {
    os.Remove(sock)
}

func unixPipe(c *net.UnixConn) {
    defer func() {
        fmt.Println("disconnected", c.RemoteAddr().String())
        c.Close()
    }()

    for {
        buf := make([]byte, 8192)
        nr, err := c.Read(buf)

        if err == io.EOF {
            log.Println("⚠️ 收到服务端的结束信号")
            break //如果收到结束信号，则退出“接收循环”，结束客户端程序
        }

        if err != nil {
            log.Println("read error:", err)
            break
        }

        data := buf[0:nr]
        log.Println("Server got:", string(data))
        _, err = c.Write(data)
        if err != nil {
            log.Fatal("Writing client error: ", err)
        }
    }
}

func onXwafMessageReceived(messages chan *mobsync.Message) {

}

func checkError(err error, info string) (res bool) {
    if err != nil {
        log.Println(info + "  " + err.Error())
        return false
    }
    return true
}

func send(conn *net.TCPConn, message *mobsync.Message) bool {
    _, err := conn.Write(mobsync.MessageToBuf(message))
    //fmt.Println(lens)
    if err != nil {
        fmt.Println(err.Error())
        conn.Close()
        return false
    }
    return true
}
