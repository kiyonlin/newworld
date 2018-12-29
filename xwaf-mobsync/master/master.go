package main

import (
    "net"
    "log"
    "time"
    "github.com/kiyonlin/newworld/xwaf-mobsync/mobsync"
)

var (
    entering    = make(chan mobsync.Node)
    leaving     = make(chan mobsync.Node)
    messages    = make(chan *mobsync.Message, 512)
    sessions    = make(map[string]time.Time)
    sessionChan = make(chan mobsync.Session, 64)
)

func main() {
    log.Println("开启服务...")
    listener, err := net.Listen("tcp", "0.0.0.0:12345")
    if err != nil {
        log.Fatal(err)
    }

    go broadcaster()

    go handleSession()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("you got something wrong %v", err)
            continue
        }
        go handleConn(conn)
    }
}

func broadcaster() {
    nodes := make(map[mobsync.Node]bool) //all connected nodes
    for {
        select {
        case message := <-messages:
            for node := range nodes {
                if node.Id != message.NodeId {
                    task := &mobsync.Message{
                        Action:        mobsync.Message_DISPATCH,
                        NodeId:        message.NodeId,
                        IpBlackList:   message.IpBlackList,
                        UuidBlackList: message.UuidBlackList,
                        UuidWhiteList: message.UuidWhiteList,
                    }
                    session := mobsync.Session{mobsync.GetSessionId(), "add"}
                    task.Session = session.Id
                    log.Println("下发来自节点", message.NodeId, "上报情况", session.Id, "给节点", node.Id)
                    //sessions[session.Id] = time.Now()
                    sessionChan <- session

                    node.Replies <- task
                }
            }
        case node := <-entering:
            nodes[node] = true
        case node := <-leaving:
            delete(nodes, node)
            close(node.Replies)
        }
    }
}

func handleConn(conn net.Conn) {
    node := mobsync.Node{0, make(chan *mobsync.Message)}
    go reply(conn, node)

    buf := make([]byte, 4096)
    for {
        //node.lock.Lock()
        c, err := conn.Read(buf)
        //node.lock.Unlock()
        if err != nil {
            leaving <- node
            conn.Close()
            break
        }
        if c > 0 {
            buf[c] = 0
        }

        message := mobsync.BufToMessage(buf[0:c])

        switch message.Action {
        case mobsync.Message_CONNECT:
            node.Id = message.NodeId
            entering <- node

            reply := &mobsync.Message{
                Action:  mobsync.Message_REPLY,
                NodeId:  message.NodeId,
                Session: message.Session,
            }
            log.Println("回复节点", message.NodeId, "已收到连接消息", message.Session)

            node.Replies <- reply
        case mobsync.Message_REPORT:
            reply := &mobsync.Message{
                Action:  mobsync.Message_REPLY,
                NodeId:  message.NodeId,
                Session: message.Session,
            }
            log.Println("回复节点", message.NodeId, "已收到上报消息", message.Session)
            node.Replies <- reply

            log.Println("准备广播节点", message.NodeId, "上报的消息")
            messages <- message

            go save(message)
        case mobsync.Message_REPLY:
            //if t := sessions[message.Session]; !time.Time.IsZero(t) {
            //log.Println("节点", message.NodeId, "回复了任务", message.Session, "耗时", time.Since(sessions[message.Session]))
            //}
            sessionChan <- mobsync.Session{message.Session, "del"}
        default:
            log.Println("invalid action")
        }
    }
}

func handleSession() {
    for session := range (sessionChan) {
        switch session.Action {
        case "add":
            sessions[session.Id] = time.Now()
        case "del":
            log.Println("任务", session.Id, "耗时", time.Since(sessions[session.Id]))
            delete(sessions, session.Id)
        }
    }
}

// save message info to database
func save(message *mobsync.Message) {
    log.Println("异步保存消息数据到数据库...")
}

func reply(conn net.Conn, node mobsync.Node) {
    for reply := range node.Replies {
        conn.Write(mobsync.MessageToBuf(reply)) // NOTE: ignoring network errors
    }
}
