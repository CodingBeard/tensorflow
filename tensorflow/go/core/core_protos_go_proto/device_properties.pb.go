// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/device_properties.proto

package core_protos_go_proto

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

type DeviceProperties struct {
	// Device type (CPU, GPU, ...)
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Vendor (Intel, nvidia, ...)
	Vendor string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Model (Haswell, K40, ...)
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	// Core Frequency in Mhz
	Frequency int64 `protobuf:"varint,4,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Number of cores
	NumCores int64 `protobuf:"varint,5,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`
	// Version of the tools and libraries used with this device (e.g. gcc 4.9,
	// cudnn 5.1)
	Environment map[string]string `protobuf:"bytes,6,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Number of registers per core.
	NumRegisters int64 `protobuf:"varint,7,opt,name=num_registers,json=numRegisters,proto3" json:"num_registers,omitempty"`
	// L1 cache size in bytes
	L1CacheSize int64 `protobuf:"varint,8,opt,name=l1_cache_size,json=l1CacheSize,proto3" json:"l1_cache_size,omitempty"`
	// L2 cache size in bytes
	L2CacheSize int64 `protobuf:"varint,9,opt,name=l2_cache_size,json=l2CacheSize,proto3" json:"l2_cache_size,omitempty"`
	// L3 cache size in bytes
	L3CacheSize int64 `protobuf:"varint,10,opt,name=l3_cache_size,json=l3CacheSize,proto3" json:"l3_cache_size,omitempty"`
	// Shared memory size per multiprocessor in bytes. This field is
	// applicable to GPUs only.
	SharedMemorySizePerMultiprocessor int64 `protobuf:"varint,11,opt,name=shared_memory_size_per_multiprocessor,json=sharedMemorySizePerMultiprocessor,proto3" json:"shared_memory_size_per_multiprocessor,omitempty"`
	// Memory size in bytes
	MemorySize int64 `protobuf:"varint,12,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	// Memory bandwidth in KB/s
	Bandwidth            int64    `protobuf:"varint,13,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProperties) Reset()         { *m = DeviceProperties{} }
func (m *DeviceProperties) String() string { return proto.CompactTextString(m) }
func (*DeviceProperties) ProtoMessage()    {}
func (*DeviceProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c4fdb3c62f9732, []int{0}
}

func (m *DeviceProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProperties.Unmarshal(m, b)
}
func (m *DeviceProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProperties.Marshal(b, m, deterministic)
}
func (m *DeviceProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProperties.Merge(m, src)
}
func (m *DeviceProperties) XXX_Size() int {
	return xxx_messageInfo_DeviceProperties.Size(m)
}
func (m *DeviceProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProperties.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProperties proto.InternalMessageInfo

func (m *DeviceProperties) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DeviceProperties) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *DeviceProperties) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *DeviceProperties) GetFrequency() int64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DeviceProperties) GetNumCores() int64 {
	if m != nil {
		return m.NumCores
	}
	return 0
}

func (m *DeviceProperties) GetEnvironment() map[string]string {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *DeviceProperties) GetNumRegisters() int64 {
	if m != nil {
		return m.NumRegisters
	}
	return 0
}

func (m *DeviceProperties) GetL1CacheSize() int64 {
	if m != nil {
		return m.L1CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetL2CacheSize() int64 {
	if m != nil {
		return m.L2CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetL3CacheSize() int64 {
	if m != nil {
		return m.L3CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetSharedMemorySizePerMultiprocessor() int64 {
	if m != nil {
		return m.SharedMemorySizePerMultiprocessor
	}
	return 0
}

func (m *DeviceProperties) GetMemorySize() int64 {
	if m != nil {
		return m.MemorySize
	}
	return 0
}

func (m *DeviceProperties) GetBandwidth() int64 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

type NamedDevice struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Properties           *DeviceProperties `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NamedDevice) Reset()         { *m = NamedDevice{} }
func (m *NamedDevice) String() string { return proto.CompactTextString(m) }
func (*NamedDevice) ProtoMessage()    {}
func (*NamedDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c4fdb3c62f9732, []int{1}
}

func (m *NamedDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedDevice.Unmarshal(m, b)
}
func (m *NamedDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedDevice.Marshal(b, m, deterministic)
}
func (m *NamedDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedDevice.Merge(m, src)
}
func (m *NamedDevice) XXX_Size() int {
	return xxx_messageInfo_NamedDevice.Size(m)
}
func (m *NamedDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedDevice.DiscardUnknown(m)
}

var xxx_messageInfo_NamedDevice proto.InternalMessageInfo

func (m *NamedDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedDevice) GetProperties() *DeviceProperties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceProperties)(nil), "tensorflow.DeviceProperties")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.DeviceProperties.EnvironmentEntry")
	proto.RegisterType((*NamedDevice)(nil), "tensorflow.NamedDevice")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/device_properties.proto", fileDescriptor_07c4fdb3c62f9732)
}

var fileDescriptor_07c4fdb3c62f9732 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0xb5, 0x2b, 0xeb, 0xcb, 0x2a, 0x55, 0x16, 0x9a, 0x2c, 0x98, 0x44, 0x29, 0x42,
	0xea, 0x85, 0x94, 0xb5, 0x17, 0x84, 0x10, 0x87, 0x8d, 0x49, 0x5c, 0x06, 0x55, 0xb8, 0x71, 0xb1,
	0xd2, 0xe4, 0xb5, 0xb5, 0x88, 0xed, 0x60, 0x3b, 0x9d, 0xb2, 0x8f, 0xc7, 0xa7, 0xe2, 0x88, 0xec,
	0x84, 0x26, 0xab, 0x10, 0x97, 0xe8, 0xbd, 0xbf, 0x7f, 0xef, 0xd9, 0xca, 0xfb, 0x3f, 0x78, 0x6b,
	0x51, 0x1a, 0xa5, 0x37, 0xb9, 0xba, 0x9f, 0xa7, 0x4a, 0xe3, 0xbc, 0xd0, 0xca, 0xaa, 0x75, 0xb9,
	0x99, 0x67, 0xb8, 0xe7, 0x29, 0xb2, 0x42, 0xab, 0x02, 0xb5, 0xe5, 0x68, 0x22, 0x7f, 0x44, 0xa0,
	0xad, 0x98, 0xfe, 0xea, 0xc3, 0xf8, 0x93, 0xe7, 0x56, 0x07, 0x8c, 0x10, 0xe8, 0xdb, 0xaa, 0x40,
	0x1a, 0x4c, 0x82, 0xd9, 0x30, 0xf6, 0x31, 0xb9, 0x80, 0xc1, 0x1e, 0x65, 0xa6, 0x34, 0x3d, 0xf1,
	0x6a, 0x93, 0x91, 0xa7, 0x70, 0x2a, 0x54, 0x86, 0x39, 0xed, 0x79, 0xb9, 0x4e, 0xc8, 0x25, 0x0c,
	0x37, 0x1a, 0x7f, 0x96, 0x28, 0xd3, 0x8a, 0xf6, 0x27, 0xc1, 0xac, 0x17, 0xb7, 0x02, 0x79, 0x0e,
	0x43, 0x59, 0x0a, 0xe6, 0x5e, 0x6b, 0xe8, 0xa9, 0x3f, 0x3d, 0x93, 0xa5, 0xb8, 0x71, 0x39, 0xf9,
	0x0a, 0x21, 0xca, 0x3d, 0xd7, 0x4a, 0x0a, 0x94, 0x96, 0x0e, 0x26, 0xbd, 0x59, 0xb8, 0x78, 0x13,
	0xb5, 0x6f, 0x8e, 0x8e, 0xdf, 0x1b, 0xdd, 0xb6, 0xfc, 0xad, 0xb4, 0xba, 0x8a, 0xbb, 0x1d, 0xc8,
	0x2b, 0x18, 0xb9, 0xdb, 0x34, 0x6e, 0xb9, 0xb1, 0xa8, 0x0d, 0x7d, 0xe2, 0x6f, 0x3c, 0x97, 0xa5,
	0x88, 0xff, 0x6a, 0x64, 0x0a, 0xa3, 0xfc, 0x8a, 0xa5, 0x49, 0xba, 0x43, 0x66, 0xf8, 0x03, 0xd2,
	0x33, 0x0f, 0x85, 0xf9, 0xd5, 0x8d, 0xd3, 0xbe, 0xf1, 0x07, 0xf4, 0xcc, 0xa2, 0xcb, 0x0c, 0x1b,
	0x66, 0xf1, 0x98, 0x59, 0x76, 0x19, 0x68, 0x98, 0x65, 0xcb, 0xac, 0xe0, 0xb5, 0xd9, 0x25, 0x1a,
	0x33, 0x26, 0x50, 0x28, 0x5d, 0x79, 0x90, 0x15, 0xa8, 0x99, 0x28, 0x73, 0xcb, 0x0b, 0xad, 0x52,
	0x34, 0x46, 0x69, 0x1a, 0xfa, 0xda, 0x97, 0x35, 0x7c, 0xe7, 0x59, 0xd7, 0x60, 0x85, 0xfa, 0xee,
	0x11, 0x48, 0x5e, 0x40, 0xd8, 0x69, 0x45, 0xcf, 0x7d, 0x1d, 0x88, 0x43, 0x85, 0x9b, 0xc7, 0x3a,
	0x91, 0xd9, 0x3d, 0xcf, 0xec, 0x8e, 0x8e, 0xea, 0x79, 0x1c, 0x84, 0x67, 0x1f, 0x61, 0x7c, 0xfc,
	0x0b, 0xc9, 0x18, 0x7a, 0x3f, 0xb0, 0x6a, 0x2c, 0xe0, 0x42, 0x37, 0xe9, 0x7d, 0x92, 0x97, 0xd8,
	0x18, 0xa0, 0x4e, 0xde, 0x9f, 0xbc, 0x0b, 0xa6, 0x0c, 0xc2, 0x2f, 0x89, 0xc0, 0xac, 0x1e, 0x8c,
	0xb3, 0x8f, 0x4c, 0xc4, 0xc1, 0x3e, 0x2e, 0x26, 0x1f, 0x00, 0x5a, 0x1f, 0xfa, 0x0e, 0xe1, 0xe2,
	0xf2, 0x7f, 0x43, 0x8d, 0x3b, 0xfc, 0x35, 0x5e, 0x5f, 0x1c, 0x9f, 0xaf, 0x9c, 0x95, 0xcd, 0xf7,
	0xcf, 0x5b, 0x6e, 0x77, 0xe5, 0x3a, 0x4a, 0x95, 0x98, 0x77, 0x16, 0xe1, 0xdf, 0xe1, 0x56, 0xd5,
	0x1b, 0xe2, 0x3e, 0xcc, 0xef, 0x82, 0x61, 0x5b, 0x55, 0x47, 0xbf, 0x83, 0x60, 0x3d, 0xf0, 0xd1,
	0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xfd, 0x8a, 0x09, 0x53, 0x03, 0x00, 0x00,
}
