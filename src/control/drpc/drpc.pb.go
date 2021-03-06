// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drpc.proto

package drpc

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

//*
// Status represents the valid values for a response status.
type Status int32

const (
	Status_SUCCESS                  Status = 0
	Status_SUBMITTED                Status = 1
	Status_FAILURE                  Status = 2
	Status_UNKNOWN_MODULE           Status = 3
	Status_UNKNOWN_METHOD           Status = 4
	Status_FAILED_UNMARSHAL_CALL    Status = 5
	Status_FAILED_UNMARSHAL_PAYLOAD Status = 6
	Status_FAILED_MARSHAL           Status = 7
)

var Status_name = map[int32]string{
	0: "SUCCESS",
	1: "SUBMITTED",
	2: "FAILURE",
	3: "UNKNOWN_MODULE",
	4: "UNKNOWN_METHOD",
	5: "FAILED_UNMARSHAL_CALL",
	6: "FAILED_UNMARSHAL_PAYLOAD",
	7: "FAILED_MARSHAL",
}

var Status_value = map[string]int32{
	"SUCCESS":                  0,
	"SUBMITTED":                1,
	"FAILURE":                  2,
	"UNKNOWN_MODULE":           3,
	"UNKNOWN_METHOD":           4,
	"FAILED_UNMARSHAL_CALL":    5,
	"FAILED_UNMARSHAL_PAYLOAD": 6,
	"FAILED_MARSHAL":           7,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e4c8b3f839895dc3, []int{0}
}

//*
// Call describes a function call to be executed over the dRPC channel.
type Call struct {
	Module               int32    `protobuf:"varint,1,opt,name=module,proto3" json:"module,omitempty"`
	Method               int32    `protobuf:"varint,2,opt,name=method,proto3" json:"method,omitempty"`
	Sequence             int64    `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}
func (*Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c8b3f839895dc3, []int{0}
}

func (m *Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call.Unmarshal(m, b)
}
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call.Marshal(b, m, deterministic)
}
func (m *Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call.Merge(m, src)
}
func (m *Call) XXX_Size() int {
	return xxx_messageInfo_Call.Size(m)
}
func (m *Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Call proto.InternalMessageInfo

func (m *Call) GetModule() int32 {
	if m != nil {
		return m.Module
	}
	return 0
}

func (m *Call) GetMethod() int32 {
	if m != nil {
		return m.Method
	}
	return 0
}

func (m *Call) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Call) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

//*
// Response describes the result of a dRPC call.
type Response struct {
	Sequence             int64    `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Status               Status   `protobuf:"varint,2,opt,name=status,proto3,enum=drpc.Status" json:"status,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c8b3f839895dc3, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Response) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_SUCCESS
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterEnum("drpc.Status", Status_name, Status_value)
	proto.RegisterType((*Call)(nil), "drpc.Call")
	proto.RegisterType((*Response)(nil), "drpc.Response")
}

func init() {
	proto.RegisterFile("drpc.proto", fileDescriptor_e4c8b3f839895dc3)
}

var fileDescriptor_e4c8b3f839895dc3 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x1c, 0xc4, 0x3f, 0x37, 0x69, 0xda, 0xef, 0x4f, 0xa9, 0x2c, 0x4b, 0xa0, 0x80, 0x18, 0xaa, 0x8a,
	0xa1, 0x62, 0xe8, 0x00, 0x4f, 0x60, 0x12, 0xa3, 0x56, 0x38, 0x09, 0xb2, 0x63, 0x21, 0xa6, 0xd0,
	0x36, 0x46, 0x0c, 0xa1, 0x0e, 0x4d, 0x32, 0xf0, 0x42, 0x3c, 0x27, 0x8a, 0x89, 0x8a, 0x0a, 0xdb,
	0xdd, 0xfd, 0x7c, 0x3e, 0xcb, 0x00, 0xf9, 0xae, 0xdc, 0xcc, 0xcb, 0x9d, 0xa9, 0x0d, 0x71, 0x5b,
	0x3d, 0x7d, 0x01, 0x37, 0x58, 0x15, 0x05, 0x39, 0x05, 0xef, 0xcd, 0xe4, 0x4d, 0xa1, 0x7d, 0x34,
	0x41, 0xb3, 0xbe, 0xe8, 0x9c, 0xcd, 0x75, 0xfd, 0x6a, 0x72, 0xbf, 0xd7, 0xe5, 0xd6, 0x91, 0x73,
	0x18, 0x56, 0xfa, 0xbd, 0xd1, 0xdb, 0x8d, 0xf6, 0x9d, 0x09, 0x9a, 0x39, 0x62, 0xef, 0x09, 0x01,
	0x77, 0x6d, 0xf2, 0x0f, 0xdf, 0x9d, 0xa0, 0xd9, 0x48, 0x58, 0x3d, 0x7d, 0x86, 0xa1, 0xd0, 0x55,
	0x69, 0xb6, 0x95, 0x3e, 0xe8, 0xa2, 0x5f, 0xdd, 0x4b, 0xf0, 0xaa, 0x7a, 0x55, 0x37, 0x95, 0xdd,
	0x1b, 0x5f, 0x8f, 0xe6, 0xf6, 0xc9, 0xd2, 0x66, 0xa2, 0x63, 0xfb, 0x05, 0xe7, 0x67, 0xe1, 0xea,
	0x13, 0x81, 0xf7, 0x7d, 0x8c, 0x1c, 0xc1, 0x40, 0xaa, 0x20, 0x60, 0x52, 0xe2, 0x7f, 0xe4, 0x18,
	0xfe, 0x4b, 0x75, 0x1b, 0x2d, 0xd3, 0x94, 0x85, 0x18, 0xb5, 0xec, 0x8e, 0x2e, 0xb9, 0x12, 0x0c,
	0xf7, 0x08, 0x81, 0xb1, 0x8a, 0xef, 0xe3, 0xe4, 0x31, 0xce, 0xa2, 0x24, 0x54, 0x9c, 0x61, 0xe7,
	0x20, 0x63, 0xe9, 0x22, 0x09, 0xb1, 0x4b, 0xce, 0xe0, 0xa4, 0x2d, 0xb1, 0x30, 0x53, 0x71, 0x44,
	0x85, 0x5c, 0x50, 0x9e, 0x05, 0x94, 0x73, 0xdc, 0x27, 0x17, 0xe0, 0xff, 0x41, 0x0f, 0xf4, 0x89,
	0x27, 0x34, 0xc4, 0x5e, 0x7b, 0x59, 0x47, 0x3b, 0x86, 0x07, 0x6b, 0xcf, 0xfe, 0xff, 0xcd, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xba, 0x2f, 0xe4, 0x8d, 0x01, 0x00, 0x00,
}
