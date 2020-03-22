// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/CLR.proto

package common

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

type CLRMetric struct {
	Time                 int64      `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Cpu                  *CPU       `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gc                   *ClrGC     `protobuf:"bytes,3,opt,name=gc,proto3" json:"gc,omitempty"`
	Thread               *ClrThread `protobuf:"bytes,4,opt,name=thread,proto3" json:"thread,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CLRMetric) Reset()         { *m = CLRMetric{} }
func (m *CLRMetric) String() string { return proto.CompactTextString(m) }
func (*CLRMetric) ProtoMessage()    {}
func (*CLRMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{0}
}

func (m *CLRMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CLRMetric.Unmarshal(m, b)
}
func (m *CLRMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CLRMetric.Marshal(b, m, deterministic)
}
func (m *CLRMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CLRMetric.Merge(m, src)
}
func (m *CLRMetric) XXX_Size() int {
	return xxx_messageInfo_CLRMetric.Size(m)
}
func (m *CLRMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_CLRMetric.DiscardUnknown(m)
}

var xxx_messageInfo_CLRMetric proto.InternalMessageInfo

func (m *CLRMetric) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *CLRMetric) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *CLRMetric) GetGc() *ClrGC {
	if m != nil {
		return m.Gc
	}
	return nil
}

func (m *CLRMetric) GetThread() *ClrThread {
	if m != nil {
		return m.Thread
	}
	return nil
}

type ClrGC struct {
	Gen0CollectCount     int64    `protobuf:"varint,1,opt,name=Gen0CollectCount,proto3" json:"Gen0CollectCount,omitempty"`
	Gen1CollectCount     int64    `protobuf:"varint,2,opt,name=Gen1CollectCount,proto3" json:"Gen1CollectCount,omitempty"`
	Gen2CollectCount     int64    `protobuf:"varint,3,opt,name=Gen2CollectCount,proto3" json:"Gen2CollectCount,omitempty"`
	HeapMemory           int64    `protobuf:"varint,4,opt,name=HeapMemory,proto3" json:"HeapMemory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClrGC) Reset()         { *m = ClrGC{} }
func (m *ClrGC) String() string { return proto.CompactTextString(m) }
func (*ClrGC) ProtoMessage()    {}
func (*ClrGC) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{1}
}

func (m *ClrGC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClrGC.Unmarshal(m, b)
}
func (m *ClrGC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClrGC.Marshal(b, m, deterministic)
}
func (m *ClrGC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClrGC.Merge(m, src)
}
func (m *ClrGC) XXX_Size() int {
	return xxx_messageInfo_ClrGC.Size(m)
}
func (m *ClrGC) XXX_DiscardUnknown() {
	xxx_messageInfo_ClrGC.DiscardUnknown(m)
}

var xxx_messageInfo_ClrGC proto.InternalMessageInfo

func (m *ClrGC) GetGen0CollectCount() int64 {
	if m != nil {
		return m.Gen0CollectCount
	}
	return 0
}

func (m *ClrGC) GetGen1CollectCount() int64 {
	if m != nil {
		return m.Gen1CollectCount
	}
	return 0
}

func (m *ClrGC) GetGen2CollectCount() int64 {
	if m != nil {
		return m.Gen2CollectCount
	}
	return 0
}

func (m *ClrGC) GetHeapMemory() int64 {
	if m != nil {
		return m.HeapMemory
	}
	return 0
}

type ClrThread struct {
	AvailableCompletionPortThreads int32    `protobuf:"varint,1,opt,name=AvailableCompletionPortThreads,proto3" json:"AvailableCompletionPortThreads,omitempty"`
	AvailableWorkerThreads         int32    `protobuf:"varint,2,opt,name=AvailableWorkerThreads,proto3" json:"AvailableWorkerThreads,omitempty"`
	MaxCompletionPortThreads       int32    `protobuf:"varint,3,opt,name=MaxCompletionPortThreads,proto3" json:"MaxCompletionPortThreads,omitempty"`
	MaxWorkerThreads               int32    `protobuf:"varint,4,opt,name=MaxWorkerThreads,proto3" json:"MaxWorkerThreads,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *ClrThread) Reset()         { *m = ClrThread{} }
func (m *ClrThread) String() string { return proto.CompactTextString(m) }
func (*ClrThread) ProtoMessage()    {}
func (*ClrThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{2}
}

func (m *ClrThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClrThread.Unmarshal(m, b)
}
func (m *ClrThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClrThread.Marshal(b, m, deterministic)
}
func (m *ClrThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClrThread.Merge(m, src)
}
func (m *ClrThread) XXX_Size() int {
	return xxx_messageInfo_ClrThread.Size(m)
}
func (m *ClrThread) XXX_DiscardUnknown() {
	xxx_messageInfo_ClrThread.DiscardUnknown(m)
}

var xxx_messageInfo_ClrThread proto.InternalMessageInfo

func (m *ClrThread) GetAvailableCompletionPortThreads() int32 {
	if m != nil {
		return m.AvailableCompletionPortThreads
	}
	return 0
}

func (m *ClrThread) GetAvailableWorkerThreads() int32 {
	if m != nil {
		return m.AvailableWorkerThreads
	}
	return 0
}

func (m *ClrThread) GetMaxCompletionPortThreads() int32 {
	if m != nil {
		return m.MaxCompletionPortThreads
	}
	return 0
}

func (m *ClrThread) GetMaxWorkerThreads() int32 {
	if m != nil {
		return m.MaxWorkerThreads
	}
	return 0
}

func init() {
	proto.RegisterType((*CLRMetric)(nil), "CLRMetric")
	proto.RegisterType((*ClrGC)(nil), "ClrGC")
	proto.RegisterType((*ClrThread)(nil), "ClrThread")
}

func init() { proto.RegisterFile("common/CLR.proto", fileDescriptor_a10d56830892247a) }

var fileDescriptor_a10d56830892247a = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xb1, 0x9d, 0x06, 0xfa, 0x76, 0x29, 0x1a, 0x04, 0xd3, 0x43, 0x29, 0x39, 0x95, 0xc1,
	0xe4, 0x36, 0x83, 0x1d, 0x76, 0xeb, 0x04, 0xeb, 0x0e, 0xf5, 0x30, 0xea, 0x46, 0x61, 0x37, 0x45,
	0x7b, 0x28, 0xc6, 0xb2, 0x9e, 0x51, 0x94, 0xb5, 0xbe, 0xed, 0xf3, 0x6c, 0x1f, 0x6f, 0x5f, 0x60,
	0x58, 0xf1, 0xc2, 0x42, 0x1b, 0x76, 0xb2, 0xf9, 0xfd, 0x7f, 0x7f, 0x3d, 0xd9, 0x3c, 0x38, 0xd1,
	0xd4, 0xb6, 0xe4, 0x0a, 0x71, 0x2b, 0x79, 0xe7, 0x29, 0xd0, 0xe9, 0xcb, 0x91, 0x6c, 0x1f, 0x5b,
	0x38, 0x5f, 0xc3, 0xb1, 0xb8, 0x95, 0x25, 0x06, 0x5f, 0x6b, 0xc6, 0x60, 0x12, 0xea, 0x16, 0xf3,
	0xe4, 0x3c, 0xb9, 0xc8, 0x64, 0x7c, 0x67, 0x33, 0xc8, 0x74, 0xb7, 0xc9, 0xd3, 0xf3, 0xe4, 0xe2,
	0xc5, 0x62, 0xc2, 0x45, 0xf5, 0x45, 0x0e, 0x80, 0xcd, 0x20, 0x35, 0x3a, 0xcf, 0x22, 0x9e, 0x72,
	0x61, 0xfd, 0x8d, 0x90, 0xa9, 0xd1, 0x6c, 0x0e, 0xd3, 0xb0, 0xf2, 0xa8, 0xbe, 0xe5, 0x93, 0x98,
	0xc1, 0x90, 0x7d, 0x8e, 0x44, 0x8e, 0xc9, 0xfc, 0x57, 0x02, 0x47, 0xb1, 0xc1, 0x5e, 0xc1, 0xc9,
	0x0d, 0xba, 0x4b, 0x41, 0xd6, 0xa2, 0x0e, 0x82, 0x36, 0x2e, 0x8c, 0xd3, 0x9f, 0xf0, 0xd1, 0xbd,
	0xda, 0x73, 0xd3, 0x9d, 0x7b, 0xf5, 0x8c, 0xbb, 0xd8, 0x73, 0xb3, 0x9d, 0xbb, 0xc7, 0xd9, 0x19,
	0xc0, 0x47, 0x54, 0x5d, 0x89, 0x2d, 0xf9, 0x3e, 0xde, 0x3a, 0x93, 0xff, 0x90, 0xf9, 0xef, 0x04,
	0x8e, 0x77, 0xdf, 0xc0, 0x3e, 0xc0, 0xd9, 0xf5, 0x77, 0x55, 0x5b, 0xb5, 0xb4, 0x28, 0xa8, 0xed,
	0x2c, 0x86, 0x9a, 0x5c, 0x45, 0x3e, 0x6c, 0x85, 0x75, 0xbc, 0xff, 0x91, 0xfc, 0x8f, 0xc5, 0xde,
	0xc2, 0x6c, 0x67, 0xdc, 0x93, 0x6f, 0xd0, 0xff, 0xed, 0xa7, 0xb1, 0x7f, 0x20, 0x65, 0xef, 0x20,
	0x2f, 0xd5, 0xe3, 0xf3, 0x93, 0xb3, 0xd8, 0x3c, 0x98, 0x0f, 0x7f, 0xa5, 0x54, 0x8f, 0xfb, 0xd3,
	0x26, 0xb1, 0xf3, 0x84, 0xbf, 0xff, 0x91, 0xc0, 0x25, 0x79, 0xc3, 0x55, 0xa7, 0xf4, 0x0a, 0xf9,
	0xba, 0xe9, 0x1f, 0x94, 0x6d, 0x6a, 0x37, 0x90, 0x96, 0x3b, 0x0c, 0x0f, 0xe4, 0x1b, 0x6e, 0x95,
	0x33, 0x1b, 0x65, 0x90, 0x2b, 0x83, 0x2e, 0x54, 0xc9, 0xd7, 0xd7, 0xa6, 0x0e, 0xab, 0xcd, 0x92,
	0x6b, 0x6a, 0x8b, 0xbb, 0xa6, 0xbf, 0xae, 0xca, 0xc2, 0xd0, 0x62, 0xdd, 0xf4, 0x85, 0xc7, 0x8e,
	0x7c, 0x40, 0x5f, 0x18, 0xdf, 0xe9, 0x71, 0x09, 0x7f, 0xa6, 0xa7, 0x77, 0x4d, 0x7f, 0x3f, 0x1e,
	0xfc, 0x69, 0x7b, 0x68, 0x35, 0x2c, 0xa6, 0x26, 0xbb, 0x9c, 0xc6, 0x15, 0x7d, 0xf3, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x7a, 0x6e, 0x4b, 0xd9, 0xcb, 0x02, 0x00, 0x00,
}
