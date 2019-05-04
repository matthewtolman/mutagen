// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session/selection.proto

package session

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Selection encodes a selection mechanism that can be used to select a
// collection of sessions. It should have exactly one member set.
type Selection struct {
	// All, if true, indicates that all sessions should be selected.
	All bool `protobuf:"varint,1,opt,name=all,proto3" json:"all,omitempty"`
	// Specifications is a list of standard Mutagen specifications (identifiers
	// and/or fragments). If non-empty, it indicates that these specifications
	// should be used to select sessions.
	Specifications []string `protobuf:"bytes,2,rep,name=specifications,proto3" json:"specifications,omitempty"`
	// LabelSelector is a label selector specification. If present (non-empty),
	// it indicates that this selector should be used to select sessions.
	LabelSelector        string   `protobuf:"bytes,3,opt,name=labelSelector,proto3" json:"labelSelector,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Selection) Reset()         { *m = Selection{} }
func (m *Selection) String() string { return proto.CompactTextString(m) }
func (*Selection) ProtoMessage()    {}
func (*Selection) Descriptor() ([]byte, []int) {
	return fileDescriptor_602d53804a0e1465, []int{0}
}

func (m *Selection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selection.Unmarshal(m, b)
}
func (m *Selection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selection.Marshal(b, m, deterministic)
}
func (m *Selection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selection.Merge(m, src)
}
func (m *Selection) XXX_Size() int {
	return xxx_messageInfo_Selection.Size(m)
}
func (m *Selection) XXX_DiscardUnknown() {
	xxx_messageInfo_Selection.DiscardUnknown(m)
}

var xxx_messageInfo_Selection proto.InternalMessageInfo

func (m *Selection) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func (m *Selection) GetSpecifications() []string {
	if m != nil {
		return m.Specifications
	}
	return nil
}

func (m *Selection) GetLabelSelector() string {
	if m != nil {
		return m.LabelSelector
	}
	return ""
}

func init() {
	proto.RegisterType((*Selection)(nil), "session.Selection")
}

func init() { proto.RegisterFile("session/selection.proto", fileDescriptor_602d53804a0e1465) }

var fileDescriptor_602d53804a0e1465 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x4a, 0x28, 0x65, 0x73, 0x71, 0x06, 0xc3, 0xe4, 0x84, 0x04,
	0xb8, 0x98, 0x13, 0x73, 0x72, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0x40, 0x4c, 0x21, 0x35,
	0x2e, 0xbe, 0xe2, 0x82, 0xd4, 0xe4, 0xcc, 0xb4, 0xcc, 0xe4, 0x44, 0x90, 0x92, 0x62, 0x09, 0x26,
	0x05, 0x66, 0x0d, 0xce, 0x20, 0x34, 0x51, 0x21, 0x15, 0x2e, 0xde, 0x9c, 0xc4, 0xa4, 0xd4, 0x1c,
	0x88, 0x59, 0xf9, 0x45, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xa8, 0x82, 0x4e, 0x9a, 0x51,
	0xea, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x19, 0x89, 0x65, 0xf9,
	0xc9, 0xba, 0x99, 0xf9, 0xfa, 0xb9, 0xa5, 0x25, 0x89, 0xe9, 0xa9, 0x79, 0xfa, 0x05, 0xd9, 0xe9,
	0xfa, 0x50, 0x77, 0x25, 0xb1, 0x81, 0xdd, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x52, 0x93,
	0xf1, 0x92, 0xc2, 0x00, 0x00, 0x00,
}