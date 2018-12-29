package main

import (
    "log"
    "github.com/tidwall/evio"
    "github.com/golang/protobuf/proto"
    pb "github.com/kiyonlin/newworld/mobsync"
    "time"
    "crypto/md5"
    "encoding/hex"
)

type node struct {
    id      int32
    replies chan *pb.Message
}

type Session struct {
    id     string
    action string
}

var (
    entering    = make(chan node)
    leaving     = make(chan node)
    messages    = make(chan *pb.Message, 512)
    sessions    = make(map[string]time.Time)
    sessionChan = make(chan Session, 64)
)

func init() {
    go broadcaster()
    go handleSession()
}

func main() {
    var events evio.Events
    events.Opened = func(c evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
        n := node{0, make(chan *pb.Message)}
        //log.Println("连接打开：", n)
        c.SetContext(n)
        return
    }

    events.Closed = func(c evio.Conn, err error) (action evio.Action) {
        n := c.Context().(node)
        //log.Println("连接关闭：", n)
        leaving <- n
        return evio.Close
    }

    events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
        c.Wake()
        n := c.Context().(node)
        if in == nil {
            //log.Printf("wake from %s\n", c.RemoteAddr())
            return reply(n), evio.None
        }

        go accept(in, n)

        return
    }
    if err := evio.Serve(events, "tcp://localhost:12345"); err != nil {
        panic(err.Error())
    }
}

func broadcaster() {
    nodes := make(map[node]bool) //all connected nodes
    for {
        select {
        case message := <-messages:
            for node := range nodes {
                if node.id != message.NodeId {
                    task := &pb.Message{
                        Action:        pb.Message_DISPATCH,
                        NodeId:        message.NodeId,
                        IpBlackList:   message.IpBlackList,
                        UuidBlackList: message.UuidBlackList,
                        UuidWhiteList: message.UuidWhiteList,
                    }
                    session := Session{getSessionId(), "add"}
                    task.Session = session.id
                    log.Println("下发来自节点", message.NodeId, "上报情况", session.id, "给节点", node.id)
                    //sessions[session.id] = time.Now()
                    sessionChan <- session

                    node.replies <- task
                }
            }
        case node := <-entering:
            nodes[node] = true
        case node := <-leaving:
            delete(nodes, node)
            close(node.replies)
        }
    }
}

func accept(in []byte, n node) {
    message := &pb.Message{}
    proto.Unmarshal(in, message)

    switch message.Action {
    case pb.Message_CONNECT:
        n.id = message.NodeId
        entering <- n

        reply := &pb.Message{
            Action:  pb.Message_REPLY,
            NodeId:  message.NodeId,
            Session: message.Session,
        }
        log.Println("回复节点", message.NodeId, "已收到连接消息", message.Session)

        n.replies <- reply
    case pb.Message_REPORT:
        reply := &pb.Message{
            Action:  pb.Message_REPLY,
            NodeId:  message.NodeId,
            Session: message.Session,
        }
        log.Println("回复节点", message.NodeId, "已收到上报消息", message.Session)
        n.replies <- reply

        log.Println("准备广播节点", message.NodeId, "上报的消息")
        messages <- message

        go save(message)
    case pb.Message_REPLY:
        sessionChan <- Session{message.Session, "del"}
    default:
        log.Println("invalid action")
    }
}

// save message info to database
func save(message *pb.Message) {
    log.Println("异步保存消息数据到数据库...")
}

func reply(n node) []byte {
    select {
    case reply := <-n.replies:
        out, _ := proto.Marshal(reply);
        return out
    default:
        return nil
    }
}

// getSession returns a random session id for a message
func getSessionId() string {
    signByte := []byte(time.Now().String())
    hash := md5.New()
    hash.Write(signByte)
    return hex.EncodeToString(hash.Sum(nil))
}

func handleSession() {
    for session := range (sessionChan) {
        switch session.action {
        case "add":
            sessions[session.id] = time.Now()
        case "del":
            log.Println("任务", session.id, "耗时", time.Since(sessions[session.id]))
            delete(sessions, session.id)
        }
    }
}
