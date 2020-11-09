// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/remote_tensor_handle.proto

package core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
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

type ResourceDtypeAndShape struct {
	Dtype                types_go_proto.DataType                 `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ResourceDtypeAndShape) Reset()         { *m = ResourceDtypeAndShape{} }
func (m *ResourceDtypeAndShape) String() string { return proto.CompactTextString(m) }
func (*ResourceDtypeAndShape) ProtoMessage()    {}
func (*ResourceDtypeAndShape) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7497b828cba9b15, []int{0}
}

func (m *ResourceDtypeAndShape) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceDtypeAndShape.Unmarshal(m, b)
}
func (m *ResourceDtypeAndShape) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceDtypeAndShape.Marshal(b, m, deterministic)
}
func (m *ResourceDtypeAndShape) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDtypeAndShape.Merge(m, src)
}
func (m *ResourceDtypeAndShape) XXX_Size() int {
	return xxx_messageInfo_ResourceDtypeAndShape.Size(m)
}
func (m *ResourceDtypeAndShape) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDtypeAndShape.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDtypeAndShape proto.InternalMessageInfo

func (m *ResourceDtypeAndShape) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *ResourceDtypeAndShape) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

type RemoteTensorHandle struct {
	// The ID of the operation that produced this tensor.
	OpId int64 `protobuf:"varint,1,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	// The index into the outputs of the operation that produced this tensor.
	OutputNum int32 `protobuf:"varint,2,opt,name=output_num,json=outputNum,proto3" json:"output_num,omitempty"`
	// Device where the tensor is located. Cannot be empty.
	// For multi-device functions, it's the default device passed to placer.
	Device string `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	// Device of the operation producing this tensor. Can be empty if the
	// operation producing this tensor is a multi-device function.
	OpDevice string `protobuf:"bytes,4,opt,name=op_device,json=opDevice,proto3" json:"op_device,omitempty"`
	// Tensor type.
	Dtype types_go_proto.DataType `protobuf:"varint,5,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Optional data types and shapes of a remote resource variable.
	ResourceDtypesAndShapes []*ResourceDtypeAndShape `protobuf:"bytes,6,rep,name=resource_dtypes_and_shapes,json=resourceDtypesAndShapes,proto3" json:"resource_dtypes_and_shapes,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *RemoteTensorHandle) Reset()         { *m = RemoteTensorHandle{} }
func (m *RemoteTensorHandle) String() string { return proto.CompactTextString(m) }
func (*RemoteTensorHandle) ProtoMessage()    {}
func (*RemoteTensorHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7497b828cba9b15, []int{1}
}

func (m *RemoteTensorHandle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteTensorHandle.Unmarshal(m, b)
}
func (m *RemoteTensorHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteTensorHandle.Marshal(b, m, deterministic)
}
func (m *RemoteTensorHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteTensorHandle.Merge(m, src)
}
func (m *RemoteTensorHandle) XXX_Size() int {
	return xxx_messageInfo_RemoteTensorHandle.Size(m)
}
func (m *RemoteTensorHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteTensorHandle.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteTensorHandle proto.InternalMessageInfo

func (m *RemoteTensorHandle) GetOpId() int64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *RemoteTensorHandle) GetOutputNum() int32 {
	if m != nil {
		return m.OutputNum
	}
	return 0
}

func (m *RemoteTensorHandle) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *RemoteTensorHandle) GetOpDevice() string {
	if m != nil {
		return m.OpDevice
	}
	return ""
}

func (m *RemoteTensorHandle) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *RemoteTensorHandle) GetResourceDtypesAndShapes() []*ResourceDtypeAndShape {
	if m != nil {
		return m.ResourceDtypesAndShapes
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceDtypeAndShape)(nil), "tensorflow.eager.ResourceDtypeAndShape")
	proto.RegisterType((*RemoteTensorHandle)(nil), "tensorflow.eager.RemoteTensorHandle")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/remote_tensor_handle.proto", fileDescriptor_c7497b828cba9b15)
}

var fileDescriptor_c7497b828cba9b15 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0x45, 0x49, 0x1c, 0x36, 0x0a, 0x2c, 0x8b, 0x76, 0xb7, 0x35, 0x69, 0x0b, 0x26, 0x50, 0x6a,
	0x4a, 0xb1, 0xc1, 0xf9, 0x82, 0x86, 0x1c, 0xd2, 0x4b, 0x09, 0x6a, 0x4e, 0xbd, 0x08, 0xc7, 0x52,
	0x9c, 0xd0, 0xd8, 0x23, 0x24, 0xbb, 0x21, 0xe7, 0x7e, 0x40, 0x7f, 0xb7, 0xc7, 0x62, 0xc9, 0x4d,
	0x4d, 0x1a, 0xda, 0x8b, 0x19, 0xbf, 0xf7, 0xf4, 0x46, 0xf3, 0x46, 0x78, 0x54, 0x88, 0x5c, 0x83,
	0x5a, 0x6e, 0x60, 0x1b, 0x26, 0xa0, 0x44, 0x28, 0x15, 0x14, 0xb0, 0x28, 0x97, 0xa1, 0x12, 0x19,
	0x14, 0x82, 0x59, 0x9e, 0xad, 0xe2, 0x9c, 0x6f, 0x44, 0x60, 0x58, 0xf2, 0xe7, 0xf3, 0x50, 0x20,
	0xe2, 0x54, 0xa8, 0xc1, 0xcd, 0xa1, 0xcd, 0x52, 0xc5, 0x99, 0xd8, 0x82, 0x7a, 0x0a, 0x6b, 0x03,
	0xbd, 0x8a, 0x65, 0x7d, 0x7e, 0x70, 0xf9, 0x8d, 0x7a, 0x27, 0x85, 0xb6, 0xb2, 0xe1, 0x16, 0xff,
	0xa7, 0x42, 0x43, 0xa9, 0x12, 0x31, 0xa9, 0xf0, 0xdb, 0x9c, 0x3f, 0x54, 0x2e, 0xe4, 0x1a, 0x3b,
	0xbc, 0x02, 0x5c, 0xe4, 0x21, 0xff, 0x77, 0xf4, 0x2f, 0x68, 0xdc, 0x67, 0x12, 0x17, 0xf1, 0x7c,
	0x27, 0x05, 0xb5, 0x12, 0x12, 0x61, 0xc7, 0xb4, 0x76, 0x5b, 0x1e, 0xf2, 0xfb, 0xd1, 0x79, 0x53,
	0x3b, 0x37, 0xa5, 0xf1, 0x9c, 0x55, 0x1d, 0xa9, 0x95, 0x0e, 0x5f, 0x5b, 0x98, 0x50, 0x33, 0xbe,
	0x55, 0x4c, 0xcd, 0xf0, 0xe4, 0x2f, 0x76, 0x40, 0xb2, 0x35, 0x37, 0x6d, 0xdb, 0xb4, 0x03, 0xf2,
	0x8e, 0x93, 0x0b, 0x8c, 0xa1, 0x2c, 0x64, 0x59, 0xb0, 0xbc, 0xcc, 0x4c, 0x13, 0x87, 0xf6, 0x2c,
	0x72, 0x5f, 0x66, 0xe4, 0x04, 0x77, 0xb9, 0x78, 0x5e, 0x27, 0xc2, 0x6d, 0x7b, 0xc8, 0xef, 0xd1,
	0xfa, 0x8f, 0x9c, 0xe1, 0x1e, 0x48, 0x56, 0x53, 0x1d, 0x43, 0xfd, 0x02, 0x39, 0xb1, 0xe4, 0x7e,
	0x3e, 0xe7, 0xe7, 0xf9, 0x38, 0x1e, 0xa8, 0x3a, 0x24, 0x66, 0x10, 0xcd, 0xe2, 0x9c, 0xdb, 0xb8,
	0xb5, 0xdb, 0xf5, 0xda, 0x7e, 0x3f, 0xba, 0x0a, 0x0e, 0x17, 0x16, 0x1c, 0x0d, 0x96, 0x9e, 0xaa,
	0x26, 0xac, 0x3f, 0x70, 0x3d, 0x7e, 0x41, 0xd8, 0x05, 0x95, 0x36, 0x7d, 0xf6, 0x3b, 0x1b, 0xbb,
	0x5f, 0xb3, 0x32, 0x71, 0xea, 0x19, 0x7a, 0x9c, 0xa6, 0xeb, 0x62, 0x55, 0x2e, 0x82, 0x04, 0xb2,
	0xb0, 0xb1, 0xf5, 0xe3, 0x65, 0x0a, 0xf6, 0x39, 0x54, 0x1f, 0x66, 0xde, 0x80, 0x66, 0x29, 0xd8,
	0xea, 0x0d, 0xa1, 0x45, 0xd7, 0x54, 0xa3, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xd3, 0x55,
	0xd5, 0xb5, 0x02, 0x00, 0x00,
}
