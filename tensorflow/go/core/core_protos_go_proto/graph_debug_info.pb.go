// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/graph_debug_info.proto

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

type GraphDebugInfo struct {
	// This stores all the source code file names and can be indexed by the
	// `file_index`.
	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// This maps a node name to a stack trace in the source code.
	// The map key is a mangling of the containing function and op name with
	// syntax:
	//   op.name '@' func_name
	// For ops in the top-level graph, the func_name is the empty string.
	// Note that op names are restricted to a small number of characters which
	// exclude '@', making it impossible to collide keys of this form. Function
	// names accept a much wider set of characters.
	// It would be preferable to avoid mangling and use a tuple key of (op.name,
	// func_name), but this is not supported with protocol buffers.
	Traces               map[string]*GraphDebugInfo_StackTrace `protobuf:"bytes,2,rep,name=traces,proto3" json:"traces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *GraphDebugInfo) Reset()         { *m = GraphDebugInfo{} }
func (m *GraphDebugInfo) String() string { return proto.CompactTextString(m) }
func (*GraphDebugInfo) ProtoMessage()    {}
func (*GraphDebugInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d49d5c184d173e1, []int{0}
}

func (m *GraphDebugInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphDebugInfo.Unmarshal(m, b)
}
func (m *GraphDebugInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphDebugInfo.Marshal(b, m, deterministic)
}
func (m *GraphDebugInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphDebugInfo.Merge(m, src)
}
func (m *GraphDebugInfo) XXX_Size() int {
	return xxx_messageInfo_GraphDebugInfo.Size(m)
}
func (m *GraphDebugInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphDebugInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GraphDebugInfo proto.InternalMessageInfo

func (m *GraphDebugInfo) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *GraphDebugInfo) GetTraces() map[string]*GraphDebugInfo_StackTrace {
	if m != nil {
		return m.Traces
	}
	return nil
}

// This represents a file/line location in the source code.
type GraphDebugInfo_FileLineCol struct {
	// File name index, which can be used to retrieve the file name string from
	// `files`. The value should be between 0 and (len(files)-1)
	FileIndex int32 `protobuf:"varint,1,opt,name=file_index,json=fileIndex,proto3" json:"file_index,omitempty"`
	// Line number in the file.
	Line int32 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Col number in the file line.
	Col int32 `protobuf:"varint,3,opt,name=col,proto3" json:"col,omitempty"`
	// Name of function contains the file line.
	Func string `protobuf:"bytes,4,opt,name=func,proto3" json:"func,omitempty"`
	// Source code contained in this file line.
	Code                 string   `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphDebugInfo_FileLineCol) Reset()         { *m = GraphDebugInfo_FileLineCol{} }
func (m *GraphDebugInfo_FileLineCol) String() string { return proto.CompactTextString(m) }
func (*GraphDebugInfo_FileLineCol) ProtoMessage()    {}
func (*GraphDebugInfo_FileLineCol) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d49d5c184d173e1, []int{0, 0}
}

func (m *GraphDebugInfo_FileLineCol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphDebugInfo_FileLineCol.Unmarshal(m, b)
}
func (m *GraphDebugInfo_FileLineCol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphDebugInfo_FileLineCol.Marshal(b, m, deterministic)
}
func (m *GraphDebugInfo_FileLineCol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphDebugInfo_FileLineCol.Merge(m, src)
}
func (m *GraphDebugInfo_FileLineCol) XXX_Size() int {
	return xxx_messageInfo_GraphDebugInfo_FileLineCol.Size(m)
}
func (m *GraphDebugInfo_FileLineCol) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphDebugInfo_FileLineCol.DiscardUnknown(m)
}

var xxx_messageInfo_GraphDebugInfo_FileLineCol proto.InternalMessageInfo

func (m *GraphDebugInfo_FileLineCol) GetFileIndex() int32 {
	if m != nil {
		return m.FileIndex
	}
	return 0
}

func (m *GraphDebugInfo_FileLineCol) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *GraphDebugInfo_FileLineCol) GetCol() int32 {
	if m != nil {
		return m.Col
	}
	return 0
}

func (m *GraphDebugInfo_FileLineCol) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *GraphDebugInfo_FileLineCol) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// This represents a stack trace which is a ordered list of `FileLineCol`.
type GraphDebugInfo_StackTrace struct {
	// Each line in the stack trace.
	FileLineCols         []*GraphDebugInfo_FileLineCol `protobuf:"bytes,1,rep,name=file_line_cols,json=fileLineCols,proto3" json:"file_line_cols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GraphDebugInfo_StackTrace) Reset()         { *m = GraphDebugInfo_StackTrace{} }
func (m *GraphDebugInfo_StackTrace) String() string { return proto.CompactTextString(m) }
func (*GraphDebugInfo_StackTrace) ProtoMessage()    {}
func (*GraphDebugInfo_StackTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d49d5c184d173e1, []int{0, 1}
}

func (m *GraphDebugInfo_StackTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphDebugInfo_StackTrace.Unmarshal(m, b)
}
func (m *GraphDebugInfo_StackTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphDebugInfo_StackTrace.Marshal(b, m, deterministic)
}
func (m *GraphDebugInfo_StackTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphDebugInfo_StackTrace.Merge(m, src)
}
func (m *GraphDebugInfo_StackTrace) XXX_Size() int {
	return xxx_messageInfo_GraphDebugInfo_StackTrace.Size(m)
}
func (m *GraphDebugInfo_StackTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphDebugInfo_StackTrace.DiscardUnknown(m)
}

var xxx_messageInfo_GraphDebugInfo_StackTrace proto.InternalMessageInfo

func (m *GraphDebugInfo_StackTrace) GetFileLineCols() []*GraphDebugInfo_FileLineCol {
	if m != nil {
		return m.FileLineCols
	}
	return nil
}

func init() {
	proto.RegisterType((*GraphDebugInfo)(nil), "tensorflow.GraphDebugInfo")
	proto.RegisterMapType((map[string]*GraphDebugInfo_StackTrace)(nil), "tensorflow.GraphDebugInfo.TracesEntry")
	proto.RegisterType((*GraphDebugInfo_FileLineCol)(nil), "tensorflow.GraphDebugInfo.FileLineCol")
	proto.RegisterType((*GraphDebugInfo_StackTrace)(nil), "tensorflow.GraphDebugInfo.StackTrace")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/graph_debug_info.proto", fileDescriptor_2d49d5c184d173e1)
}

var fileDescriptor_2d49d5c184d173e1 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0xcb, 0xdb, 0x30,
	0x10, 0xc6, 0x51, 0xfc, 0x3a, 0xe0, 0x73, 0x09, 0x45, 0x64, 0x10, 0x81, 0x82, 0x29, 0xb4, 0x78,
	0xb2, 0x21, 0x5d, 0x4a, 0x0b, 0x1d, 0xd2, 0xbf, 0x81, 0x0c, 0xc1, 0xed, 0x94, 0xc5, 0xb5, 0x15,
	0xc9, 0x31, 0x51, 0x74, 0x41, 0xb6, 0x9b, 0xa6, 0x4b, 0xbf, 0x4a, 0x3f, 0x66, 0xc7, 0x22, 0xb9,
	0x60, 0x07, 0x5e, 0xb2, 0x98, 0x9f, 0x1e, 0xdd, 0x3d, 0xba, 0x07, 0x1f, 0xa4, 0xad, 0xd0, 0x0d,
	0x1a, 0xa9, 0xf0, 0x92, 0x72, 0x34, 0x22, 0x3d, 0x1b, 0x6c, 0xb1, 0xec, 0x64, 0x5a, 0x99, 0xe2,
	0x7c, 0xc8, 0xf7, 0xa2, 0xec, 0xaa, 0xbc, 0xd6, 0x12, 0x13, 0x77, 0x43, 0x61, 0x68, 0x78, 0xfe,
	0xc7, 0x83, 0xd9, 0x67, 0x5b, 0xf6, 0xc1, 0x56, 0xad, 0xb5, 0x44, 0x3a, 0x07, 0x5f, 0xd6, 0x4a,
	0x34, 0x8c, 0x44, 0x5e, 0x1c, 0x64, 0xfd, 0x81, 0xbe, 0x83, 0x69, 0x6b, 0x0a, 0x2e, 0x1a, 0x36,
	0x89, 0xbc, 0x38, 0x5c, 0xbe, 0x4c, 0x06, 0x97, 0xe4, 0xd6, 0x21, 0xf9, 0xe6, 0x0a, 0x3f, 0xea,
	0xd6, 0x5c, 0xb3, 0xff, 0x5d, 0x8b, 0x5f, 0x10, 0x7e, 0xaa, 0x95, 0xd8, 0xd4, 0x5a, 0xbc, 0x47,
	0x45, 0x9f, 0x01, 0x58, 0xdf, 0xbc, 0xd6, 0x7b, 0xf1, 0x93, 0x91, 0x88, 0xc4, 0x7e, 0x16, 0x58,
	0x65, 0x6d, 0x05, 0x4a, 0xe1, 0x41, 0xd5, 0x5a, 0xb0, 0x89, 0xbb, 0x70, 0x4c, 0x9f, 0x82, 0xc7,
	0x51, 0x31, 0xcf, 0x49, 0x16, 0x6d, 0x95, 0xec, 0x34, 0x67, 0x0f, 0x11, 0x89, 0x83, 0xcc, 0xb1,
	0xd5, 0x38, 0xee, 0x05, 0xf3, 0x7b, 0xcd, 0xf2, 0x62, 0x07, 0xf0, 0xb5, 0x2d, 0xf8, 0xd1, 0xcd,
	0x45, 0x37, 0x30, 0x73, 0x4f, 0x5b, 0xd3, 0x9c, 0xa3, 0xea, 0x83, 0xde, 0x4f, 0x34, 0x1a, 0x3d,
	0x7b, 0x22, 0x87, 0x43, 0xb3, 0xf8, 0x0e, 0xe1, 0x28, 0xae, 0x1d, 0xf2, 0x28, 0xae, 0x2e, 0x50,
	0x90, 0x59, 0xa4, 0x6f, 0xc1, 0xff, 0x51, 0xa8, 0xae, 0xcf, 0x12, 0x2e, 0x5f, 0xdc, 0x79, 0x65,
	0x18, 0x32, 0xeb, 0x7b, 0xde, 0x4c, 0x5e, 0x93, 0xd5, 0x6f, 0x60, 0x68, 0xaa, 0x71, 0x9b, 0x34,
	0xc5, 0x49, 0x5c, 0xd0, 0x1c, 0x57, 0xf3, 0x5b, 0x87, 0xad, 0xfd, 0xbf, 0xcd, 0x96, 0xec, 0xbe,
	0x54, 0x75, 0x7b, 0xe8, 0xca, 0x84, 0xe3, 0x69, 0xbc, 0x1e, 0x8f, 0x63, 0x85, 0xfd, 0xde, 0xd8,
	0x4f, 0xee, 0x56, 0xa4, 0xc9, 0x2b, 0xec, 0xe9, 0x2f, 0x21, 0xe5, 0xd4, 0xd1, 0xab, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc5, 0x35, 0x54, 0x3e, 0x69, 0x02, 0x00, 0x00,
}
