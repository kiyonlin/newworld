// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mobsync.proto

package mobsync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message_Action int32

const (
	Message_REPLY    Message_Action = 0
	Message_CONNECT  Message_Action = 1
	Message_REPORT   Message_Action = 2
	Message_DISPATCH Message_Action = 3
)

var Message_Action_name = map[int32]string{
	0: "REPLY",
	1: "CONNECT",
	2: "REPORT",
	3: "DISPATCH",
}
var Message_Action_value = map[string]int32{
	"REPLY":    0,
	"CONNECT":  1,
	"REPORT":   2,
	"DISPATCH": 3,
}

func (x Message_Action) String() string {
	return proto.EnumName(Message_Action_name, int32(x))
}
func (Message_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mobsync_6ca28e6aafd3530a, []int{0, 0}
}

type Message struct {
	Action               Message_Action `protobuf:"varint,1,opt,name=action,proto3,enum=mobsync.Message_Action" json:"action,omitempty"`
	Session              string         `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	NodeId               int32          `protobuf:"varint,3,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	SiteId               int32          `protobuf:"varint,4,opt,name=siteId,proto3" json:"siteId,omitempty"`
	IpBlackList          []string       `protobuf:"bytes,5,rep,name=ipBlackList,proto3" json:"ipBlackList,omitempty"`
	UuidBlackList        []string       `protobuf:"bytes,6,rep,name=uuidBlackList,proto3" json:"uuidBlackList,omitempty"`
	UuidWhiteList        []string       `protobuf:"bytes,7,rep,name=uuidWhiteList,proto3" json:"uuidWhiteList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobsync_6ca28e6aafd3530a, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetAction() Message_Action {
	if m != nil {
		return m.Action
	}
	return Message_REPLY
}

func (m *Message) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *Message) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *Message) GetSiteId() int32 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *Message) GetIpBlackList() []string {
	if m != nil {
		return m.IpBlackList
	}
	return nil
}

func (m *Message) GetUuidBlackList() []string {
	if m != nil {
		return m.UuidBlackList
	}
	return nil
}

func (m *Message) GetUuidWhiteList() []string {
	if m != nil {
		return m.UuidWhiteList
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "mobsync.Message")
	proto.RegisterEnum("mobsync.Message_Action", Message_Action_name, Message_Action_value)
}

func init() { proto.RegisterFile("mobsync.proto", fileDescriptor_mobsync_6ca28e6aafd3530a) }

var fileDescriptor_mobsync_6ca28e6aafd3530a = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4b, 0x4f, 0xb3, 0x40,
	0x14, 0x40, 0x3b, 0xed, 0xc7, 0xcc, 0xc7, 0xad, 0x35, 0xe4, 0x2e, 0x94, 0xb8, 0x22, 0xc4, 0x05,
	0x2b, 0x34, 0x35, 0x71, 0xe1, 0x0e, 0x91, 0xc4, 0x26, 0x7d, 0x90, 0x81, 0xc4, 0xb8, 0xa4, 0x30,
	0xd1, 0x89, 0x16, 0x9a, 0xce, 0x74, 0xd1, 0x3f, 0xe5, 0x6f, 0x34, 0xbc, 0xac, 0x8f, 0xe5, 0x39,
	0x73, 0xee, 0x62, 0xee, 0x85, 0xc9, 0xa6, 0x5a, 0xab, 0x43, 0x99, 0xfb, 0xdb, 0x5d, 0xa5, 0x2b,
	0x64, 0x1d, 0xba, 0x1f, 0x43, 0x60, 0x0b, 0xa1, 0x54, 0xf6, 0x22, 0xf0, 0x0a, 0x68, 0x96, 0x6b,
	0x59, 0x95, 0x36, 0x71, 0x88, 0x77, 0x3a, 0x3d, 0xf7, 0xfb, 0xa1, 0xae, 0xf0, 0x83, 0xe6, 0x99,
	0x77, 0x19, 0xda, 0xc0, 0x94, 0x50, 0xaa, 0x9e, 0x18, 0x3a, 0xc4, 0x33, 0x79, 0x8f, 0x78, 0x06,
	0xb4, 0xac, 0x0a, 0x31, 0x2b, 0xec, 0x91, 0x43, 0x3c, 0x83, 0x77, 0x54, 0x7b, 0x25, 0x75, 0xed,
	0xff, 0xb5, 0xbe, 0x25, 0x74, 0x60, 0x2c, 0xb7, 0xf7, 0xef, 0x59, 0xfe, 0x36, 0x97, 0x4a, 0xdb,
	0x86, 0x33, 0xf2, 0x4c, 0xfe, 0x5d, 0xe1, 0x25, 0x4c, 0xf6, 0x7b, 0x59, 0x1c, 0x1b, 0xda, 0x34,
	0x3f, 0x65, 0x5f, 0x3d, 0xbd, 0x4a, 0x2d, 0x9a, 0x8a, 0x1d, 0xab, 0x2f, 0xe9, 0xde, 0x01, 0x6d,
	0x7f, 0x82, 0x26, 0x18, 0x3c, 0x8a, 0xe7, 0xcf, 0xd6, 0x00, 0xc7, 0xc0, 0xc2, 0xd5, 0x72, 0x19,
	0x85, 0xa9, 0x45, 0x10, 0x80, 0xf2, 0x28, 0x5e, 0xf1, 0xd4, 0x1a, 0xe2, 0x09, 0xfc, 0x7f, 0x98,
	0x25, 0x71, 0x90, 0x86, 0x8f, 0xd6, 0x68, 0x1a, 0x00, 0x5b, 0xb4, 0x5b, 0xc1, 0x5b, 0x80, 0xe4,
	0x50, 0xe6, 0x89, 0xde, 0x89, 0x6c, 0x83, 0xd6, 0xef, 0x6d, 0x5d, 0xfc, 0x31, 0xee, 0xc0, 0x23,
	0xd7, 0x64, 0x4d, 0x9b, 0x1b, 0xdc, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x80, 0x6b, 0x68, 0x1e,
	0x94, 0x01, 0x00, 0x00,
}
