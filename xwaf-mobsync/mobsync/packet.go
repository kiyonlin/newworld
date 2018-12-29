package mobsync

import (
    _ "github.com/kiyonlin/newworld/mobsync"
    "github.com/golang/protobuf/proto"
    "time"
    "crypto/md5"
    "encoding/hex"
)

type Node struct {
    Id      int32
    Replies chan *Message
}

type Session struct {
    Id     string
    Action string
}

func BufToMessage(in []byte) *Message {
    message := &Message{}
    proto.Unmarshal(in, message)
    return message
}

func MessageToBuf(message *Message) []byte {
    if out, err := proto.Marshal(message); err != nil {
        return nil
    } else {
        return out
    }
}

// GetSessionId returns a random Session id for a message
func GetSessionId() string {
    signByte := []byte(time.Now().String())
    hash := md5.New()
    hash.Write(signByte)
    return hex.EncodeToString(hash.Sum(nil))
}