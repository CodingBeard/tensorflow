// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: tensorflow/core/protobuf/service_config.proto

package tensorflow_data_experimental

import (
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

// Configuration for a tf.data service DispatchServer.
type DispatcherConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The port for the dispatcher to bind to. A value of 0 indicates that the
	// dispatcher may bind to any available port.
	Port int64 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// The protocol for the dispatcher to use when connecting to workers.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// A work directory to use for storing dispatcher state, and for recovering
	// during restarts. The empty string indicates not to use any work directory.
	WorkDir string `protobuf:"bytes,3,opt,name=work_dir,json=workDir,proto3" json:"work_dir,omitempty"`
	// Whether to run in fault tolerant mode, where dispatcher state is saved
	// across restarts. Requires that `work_dir` is nonempty.
	FaultTolerantMode bool `protobuf:"varint,4,opt,name=fault_tolerant_mode,json=faultTolerantMode,proto3" json:"fault_tolerant_mode,omitempty"`
	// How often the dispatcher should scan through to delete old and unused jobs.
	JobGcCheckIntervalMs int64 `protobuf:"varint,5,opt,name=job_gc_check_interval_ms,json=jobGcCheckIntervalMs,proto3" json:"job_gc_check_interval_ms,omitempty"`
	// How long a job needs to be unused before it becomes a candidate for garbage
	// collection.
	JobGcTimeoutMs int64 `protobuf:"varint,6,opt,name=job_gc_timeout_ms,json=jobGcTimeoutMs,proto3" json:"job_gc_timeout_ms,omitempty"`
}

func (x *DispatcherConfig) Reset() {
	*x = DispatcherConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_service_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatcherConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatcherConfig) ProtoMessage() {}

func (x *DispatcherConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_service_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatcherConfig.ProtoReflect.Descriptor instead.
func (*DispatcherConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_service_config_proto_rawDescGZIP(), []int{0}
}

func (x *DispatcherConfig) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *DispatcherConfig) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *DispatcherConfig) GetWorkDir() string {
	if x != nil {
		return x.WorkDir
	}
	return ""
}

func (x *DispatcherConfig) GetFaultTolerantMode() bool {
	if x != nil {
		return x.FaultTolerantMode
	}
	return false
}

func (x *DispatcherConfig) GetJobGcCheckIntervalMs() int64 {
	if x != nil {
		return x.JobGcCheckIntervalMs
	}
	return 0
}

func (x *DispatcherConfig) GetJobGcTimeoutMs() int64 {
	if x != nil {
		return x.JobGcTimeoutMs
	}
	return 0
}

// Configuration for a tf.data service WorkerServer.
type WorkerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The port for the worker to bind to. A value of 0 indicates that the
	// worker may bind to any available port.
	Port int64 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// The protocol for the worker to use when connecting to the dispatcher.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// The address of the dispatcher to register with.
	DispatcherAddress string `protobuf:"bytes,3,opt,name=dispatcher_address,json=dispatcherAddress,proto3" json:"dispatcher_address,omitempty"`
	// The address of the worker server. The substring "%port%", if specified,
	// will be replaced with the worker's bound port. This is useful when the port
	// is set to `0`.
	WorkerAddress string `protobuf:"bytes,4,opt,name=worker_address,json=workerAddress,proto3" json:"worker_address,omitempty"`
	// How often the worker should heartbeat to the master.
	HeartbeatIntervalMs int64 `protobuf:"varint,5,opt,name=heartbeat_interval_ms,json=heartbeatIntervalMs,proto3" json:"heartbeat_interval_ms,omitempty"`
	// How long to retry requests to the dispatcher before giving up and reporting
	// an error.
	DispatcherTimeoutMs int64 `protobuf:"varint,6,opt,name=dispatcher_timeout_ms,json=dispatcherTimeoutMs,proto3" json:"dispatcher_timeout_ms,omitempty"`
	// The protocol for the worker to use when transferring data to clients.
	DataTransferProtocol string `protobuf:"bytes,7,opt,name=data_transfer_protocol,json=dataTransferProtocol,proto3" json:"data_transfer_protocol,omitempty"`
}

func (x *WorkerConfig) Reset() {
	*x = WorkerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_service_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerConfig) ProtoMessage() {}

func (x *WorkerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_service_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerConfig.ProtoReflect.Descriptor instead.
func (*WorkerConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_service_config_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerConfig) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *WorkerConfig) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *WorkerConfig) GetDispatcherAddress() string {
	if x != nil {
		return x.DispatcherAddress
	}
	return ""
}

func (x *WorkerConfig) GetWorkerAddress() string {
	if x != nil {
		return x.WorkerAddress
	}
	return ""
}

func (x *WorkerConfig) GetHeartbeatIntervalMs() int64 {
	if x != nil {
		return x.HeartbeatIntervalMs
	}
	return 0
}

func (x *WorkerConfig) GetDispatcherTimeoutMs() int64 {
	if x != nil {
		return x.DispatcherTimeoutMs
	}
	return 0
}

func (x *WorkerConfig) GetDataTransferProtocol() string {
	if x != nil {
		return x.DataTransferProtocol
	}
	return ""
}

var File_tensorflow_core_protobuf_service_config_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_service_config_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x22, 0xf0, 0x01,
	0x0a, 0x10, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x69, 0x72, 0x12, 0x2e, 0x0a,
	0x13, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x6e, 0x74, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x54, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a,
	0x18, 0x6a, 0x6f, 0x62, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x14, 0x6a, 0x6f, 0x62, 0x47, 0x63, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x4d, 0x73, 0x12, 0x29, 0x0a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x67, 0x63, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x6a, 0x6f, 0x62, 0x47, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x73,
	0x22, 0xb2, 0x02, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x5f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x73, 0x12,
	0x34, 0x0a, 0x16, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x64, 0x61, 0x74, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_service_config_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_service_config_proto_rawDescData = file_tensorflow_core_protobuf_service_config_proto_rawDesc
)

func file_tensorflow_core_protobuf_service_config_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_service_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_service_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_service_config_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_service_config_proto_rawDescData
}

var file_tensorflow_core_protobuf_service_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_protobuf_service_config_proto_goTypes = []interface{}{
	(*DispatcherConfig)(nil), // 0: tensorflow.data.experimental.DispatcherConfig
	(*WorkerConfig)(nil),     // 1: tensorflow.data.experimental.WorkerConfig
}
var file_tensorflow_core_protobuf_service_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_service_config_proto_init() }
func file_tensorflow_core_protobuf_service_config_proto_init() {
	if File_tensorflow_core_protobuf_service_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_service_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatcherConfig); i {
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
		file_tensorflow_core_protobuf_service_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerConfig); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_service_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_service_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_service_config_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_service_config_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_service_config_proto = out.File
	file_tensorflow_core_protobuf_service_config_proto_rawDesc = nil
	file_tensorflow_core_protobuf_service_config_proto_goTypes = nil
	file_tensorflow_core_protobuf_service_config_proto_depIdxs = nil
}
