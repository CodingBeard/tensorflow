// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: tensorflow/core/framework/cost_graph.proto

package cost_graph_go_proto

import (
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CostGraphDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node []*CostGraphDef_Node           `protobuf:"bytes,1,rep,name=node,proto3" json:"node,omitempty"`
	Cost []*CostGraphDef_AggregatedCost `protobuf:"bytes,2,rep,name=cost,proto3" json:"cost,omitempty"`
}

func (x *CostGraphDef) Reset() {
	*x = CostGraphDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostGraphDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostGraphDef) ProtoMessage() {}

func (x *CostGraphDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostGraphDef.ProtoReflect.Descriptor instead.
func (*CostGraphDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_cost_graph_proto_rawDescGZIP(), []int{0}
}

func (x *CostGraphDef) GetNode() []*CostGraphDef_Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *CostGraphDef) GetCost() []*CostGraphDef_AggregatedCost {
	if x != nil {
		return x.Cost
	}
	return nil
}

type CostGraphDef_Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the node. Names are globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The device of the node. Can be empty if the node is mapped to the
	// default partition or partitioning hasn't been run yet.
	Device string `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	// The id of the node. Node ids are only unique inside a partition.
	Id         int32                           `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	InputInfo  []*CostGraphDef_Node_InputInfo  `protobuf:"bytes,4,rep,name=input_info,json=inputInfo,proto3" json:"input_info,omitempty"`
	OutputInfo []*CostGraphDef_Node_OutputInfo `protobuf:"bytes,5,rep,name=output_info,json=outputInfo,proto3" json:"output_info,omitempty"`
	// Temporary memory used by this node.
	TemporaryMemorySize int64 `protobuf:"varint,6,opt,name=temporary_memory_size,json=temporaryMemorySize,proto3" json:"temporary_memory_size,omitempty"`
	// Persistent memory used by this node.
	PersistentMemorySize int64 `protobuf:"varint,12,opt,name=persistent_memory_size,json=persistentMemorySize,proto3" json:"persistent_memory_size,omitempty"`
	// Deprecated: Do not use.
	HostTempMemorySize int64 `protobuf:"varint,10,opt,name=host_temp_memory_size,json=hostTempMemorySize,proto3" json:"host_temp_memory_size,omitempty"`
	// Deprecated: Do not use.
	DeviceTempMemorySize int64 `protobuf:"varint,11,opt,name=device_temp_memory_size,json=deviceTempMemorySize,proto3" json:"device_temp_memory_size,omitempty"`
	// Deprecated: Do not use.
	DevicePersistentMemorySize int64 `protobuf:"varint,16,opt,name=device_persistent_memory_size,json=devicePersistentMemorySize,proto3" json:"device_persistent_memory_size,omitempty"`
	// Estimate of the computational cost of this node, in microseconds.
	ComputeCost int64 `protobuf:"varint,9,opt,name=compute_cost,json=computeCost,proto3" json:"compute_cost,omitempty"`
	// Analytical estimate of the computational cost of this node, in
	// microseconds.
	ComputeTime int64 `protobuf:"varint,14,opt,name=compute_time,json=computeTime,proto3" json:"compute_time,omitempty"`
	// Analytical estimate of the memory access cost of this node, in
	// microseconds.
	MemoryTime int64 `protobuf:"varint,15,opt,name=memory_time,json=memoryTime,proto3" json:"memory_time,omitempty"`
	// If true, the output is permanent: it can't be discarded, because this
	// node is part of the "final output". Nodes may depend on final nodes.
	IsFinal bool `protobuf:"varint,7,opt,name=is_final,json=isFinal,proto3" json:"is_final,omitempty"`
	// Ids of the control inputs for this node.
	ControlInput []int32 `protobuf:"varint,8,rep,packed,name=control_input,json=controlInput,proto3" json:"control_input,omitempty"`
	// Are the costs inaccurate?
	Inaccurate bool `protobuf:"varint,17,opt,name=inaccurate,proto3" json:"inaccurate,omitempty"`
}

func (x *CostGraphDef_Node) Reset() {
	*x = CostGraphDef_Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostGraphDef_Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostGraphDef_Node) ProtoMessage() {}

func (x *CostGraphDef_Node) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostGraphDef_Node.ProtoReflect.Descriptor instead.
func (*CostGraphDef_Node) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_cost_graph_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CostGraphDef_Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CostGraphDef_Node) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *CostGraphDef_Node) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CostGraphDef_Node) GetInputInfo() []*CostGraphDef_Node_InputInfo {
	if x != nil {
		return x.InputInfo
	}
	return nil
}

func (x *CostGraphDef_Node) GetOutputInfo() []*CostGraphDef_Node_OutputInfo {
	if x != nil {
		return x.OutputInfo
	}
	return nil
}

func (x *CostGraphDef_Node) GetTemporaryMemorySize() int64 {
	if x != nil {
		return x.TemporaryMemorySize
	}
	return 0
}

func (x *CostGraphDef_Node) GetPersistentMemorySize() int64 {
	if x != nil {
		return x.PersistentMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (x *CostGraphDef_Node) GetHostTempMemorySize() int64 {
	if x != nil {
		return x.HostTempMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (x *CostGraphDef_Node) GetDeviceTempMemorySize() int64 {
	if x != nil {
		return x.DeviceTempMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (x *CostGraphDef_Node) GetDevicePersistentMemorySize() int64 {
	if x != nil {
		return x.DevicePersistentMemorySize
	}
	return 0
}

func (x *CostGraphDef_Node) GetComputeCost() int64 {
	if x != nil {
		return x.ComputeCost
	}
	return 0
}

func (x *CostGraphDef_Node) GetComputeTime() int64 {
	if x != nil {
		return x.ComputeTime
	}
	return 0
}

func (x *CostGraphDef_Node) GetMemoryTime() int64 {
	if x != nil {
		return x.MemoryTime
	}
	return 0
}

func (x *CostGraphDef_Node) GetIsFinal() bool {
	if x != nil {
		return x.IsFinal
	}
	return false
}

func (x *CostGraphDef_Node) GetControlInput() []int32 {
	if x != nil {
		return x.ControlInput
	}
	return nil
}

func (x *CostGraphDef_Node) GetInaccurate() bool {
	if x != nil {
		return x.Inaccurate
	}
	return false
}

// Total cost of this graph, typically used for balancing decisions.
type CostGraphDef_AggregatedCost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Aggregated cost value.
	Cost float32 `protobuf:"fixed32,1,opt,name=cost,proto3" json:"cost,omitempty"`
	// Aggregated cost dimension (e.g. 'memory', 'compute', 'network').
	Dimension string `protobuf:"bytes,2,opt,name=dimension,proto3" json:"dimension,omitempty"`
}

func (x *CostGraphDef_AggregatedCost) Reset() {
	*x = CostGraphDef_AggregatedCost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostGraphDef_AggregatedCost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostGraphDef_AggregatedCost) ProtoMessage() {}

func (x *CostGraphDef_AggregatedCost) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostGraphDef_AggregatedCost.ProtoReflect.Descriptor instead.
func (*CostGraphDef_AggregatedCost) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_cost_graph_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CostGraphDef_AggregatedCost) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *CostGraphDef_AggregatedCost) GetDimension() string {
	if x != nil {
		return x.Dimension
	}
	return ""
}

// Inputs of this node. They must be executed before this node can be
// executed. An input is a particular output of another node, specified
// by the node id and the output index.
type CostGraphDef_Node_InputInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrecedingNode int32 `protobuf:"varint,1,opt,name=preceding_node,json=precedingNode,proto3" json:"preceding_node,omitempty"`
	PrecedingPort int32 `protobuf:"varint,2,opt,name=preceding_port,json=precedingPort,proto3" json:"preceding_port,omitempty"`
}

func (x *CostGraphDef_Node_InputInfo) Reset() {
	*x = CostGraphDef_Node_InputInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostGraphDef_Node_InputInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostGraphDef_Node_InputInfo) ProtoMessage() {}

func (x *CostGraphDef_Node_InputInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostGraphDef_Node_InputInfo.ProtoReflect.Descriptor instead.
func (*CostGraphDef_Node_InputInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_cost_graph_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *CostGraphDef_Node_InputInfo) GetPrecedingNode() int32 {
	if x != nil {
		return x.PrecedingNode
	}
	return 0
}

func (x *CostGraphDef_Node_InputInfo) GetPrecedingPort() int32 {
	if x != nil {
		return x.PrecedingPort
	}
	return 0
}

// Outputs of this node.
type CostGraphDef_Node_OutputInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// If >= 0, the output is an alias of an input. Note that an alias input
	// may itself be an alias. The algorithm will therefore need to follow
	// those pointers.
	AliasInputPort int64                                   `protobuf:"varint,2,opt,name=alias_input_port,json=aliasInputPort,proto3" json:"alias_input_port,omitempty"`
	Shape          *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,3,opt,name=shape,proto3" json:"shape,omitempty"`
	Dtype          types_go_proto.DataType                 `protobuf:"varint,4,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
}

func (x *CostGraphDef_Node_OutputInfo) Reset() {
	*x = CostGraphDef_Node_OutputInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostGraphDef_Node_OutputInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostGraphDef_Node_OutputInfo) ProtoMessage() {}

func (x *CostGraphDef_Node_OutputInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_cost_graph_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostGraphDef_Node_OutputInfo.ProtoReflect.Descriptor instead.
func (*CostGraphDef_Node_OutputInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_cost_graph_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *CostGraphDef_Node_OutputInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CostGraphDef_Node_OutputInfo) GetAliasInputPort() int64 {
	if x != nil {
		return x.AliasInputPort
	}
	return 0
}

func (x *CostGraphDef_Node_OutputInfo) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *CostGraphDef_Node_OutputInfo) GetDtype() types_go_proto.DataType {
	if x != nil {
		return x.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

var File_tensorflow_core_framework_cost_graph_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_cost_graph_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x09,
	0x0a, 0x0c, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x66, 0x12, 0x31,
	0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x44, 0x65, 0x66, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x12, 0x3b, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x73,
	0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x66, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x1a, 0xc7,
	0x07, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x46, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65,
	0x66, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x0b, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f,
	0x73, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x66, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x65,
	0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x35, 0x0a, 0x15, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x12, 0x68, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x14, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x45, 0x0a, 0x1d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x1a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x74, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x6e, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x74, 0x65,
	0x1a, 0x59, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x72, 0x65, 0x63, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x65, 0x64, 0x69, 0x6e, 0x67,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x63, 0x65, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72,
	0x65, 0x63, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x1a, 0xaa, 0x01, 0x0a, 0x0a,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x05,
	0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x42, 0x0a, 0x0e, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x83, 0x01, 0x0a,
	0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0f, 0x43, 0x6f, 0x73, 0x74, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_cost_graph_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_cost_graph_proto_rawDescData = file_tensorflow_core_framework_cost_graph_proto_rawDesc
)

func file_tensorflow_core_framework_cost_graph_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_cost_graph_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_cost_graph_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_cost_graph_proto_rawDescData)
	})
	return file_tensorflow_core_framework_cost_graph_proto_rawDescData
}

var file_tensorflow_core_framework_cost_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tensorflow_core_framework_cost_graph_proto_goTypes = []interface{}{
	(*CostGraphDef)(nil),                           // 0: tensorflow.CostGraphDef
	(*CostGraphDef_Node)(nil),                      // 1: tensorflow.CostGraphDef.Node
	(*CostGraphDef_AggregatedCost)(nil),            // 2: tensorflow.CostGraphDef.AggregatedCost
	(*CostGraphDef_Node_InputInfo)(nil),            // 3: tensorflow.CostGraphDef.Node.InputInfo
	(*CostGraphDef_Node_OutputInfo)(nil),           // 4: tensorflow.CostGraphDef.Node.OutputInfo
	(*tensor_shape_go_proto.TensorShapeProto)(nil), // 5: tensorflow.TensorShapeProto
	(types_go_proto.DataType)(0),                   // 6: tensorflow.DataType
}
var file_tensorflow_core_framework_cost_graph_proto_depIdxs = []int32{
	1, // 0: tensorflow.CostGraphDef.node:type_name -> tensorflow.CostGraphDef.Node
	2, // 1: tensorflow.CostGraphDef.cost:type_name -> tensorflow.CostGraphDef.AggregatedCost
	3, // 2: tensorflow.CostGraphDef.Node.input_info:type_name -> tensorflow.CostGraphDef.Node.InputInfo
	4, // 3: tensorflow.CostGraphDef.Node.output_info:type_name -> tensorflow.CostGraphDef.Node.OutputInfo
	5, // 4: tensorflow.CostGraphDef.Node.OutputInfo.shape:type_name -> tensorflow.TensorShapeProto
	6, // 5: tensorflow.CostGraphDef.Node.OutputInfo.dtype:type_name -> tensorflow.DataType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_cost_graph_proto_init() }
func file_tensorflow_core_framework_cost_graph_proto_init() {
	if File_tensorflow_core_framework_cost_graph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_cost_graph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostGraphDef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_framework_cost_graph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostGraphDef_Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_framework_cost_graph_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostGraphDef_AggregatedCost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_framework_cost_graph_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostGraphDef_Node_InputInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_framework_cost_graph_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostGraphDef_Node_OutputInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_framework_cost_graph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_cost_graph_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_cost_graph_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_cost_graph_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_cost_graph_proto = out.File
	file_tensorflow_core_framework_cost_graph_proto_rawDesc = nil
	file_tensorflow_core_framework_cost_graph_proto_goTypes = nil
	file_tensorflow_core_framework_cost_graph_proto_depIdxs = nil
}
