// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: shared/proto/pipeline_reporter.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PipelineStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineId int64                  `protobuf:"varint,1,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	StartTime  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Cmd        string                 `protobuf:"bytes,3,opt,name=cmd,proto3" json:"cmd,omitempty"`
	ExitCode   int64                  `protobuf:"varint,4,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *PipelineStartRequest) Reset() {
	*x = PipelineStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_pipeline_reporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineStartRequest) ProtoMessage() {}

func (x *PipelineStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_pipeline_reporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineStartRequest.ProtoReflect.Descriptor instead.
func (*PipelineStartRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_pipeline_reporter_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineStartRequest) GetPipelineId() int64 {
	if x != nil {
		return x.PipelineId
	}
	return 0
}

func (x *PipelineStartRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *PipelineStartRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *PipelineStartRequest) GetExitCode() int64 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_pipeline_reporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_pipeline_reporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_shared_proto_pipeline_reporter_proto_rawDescGZIP(), []int{1}
}

var File_shared_proto_pipeline_reporter_proto protoreflect.FileDescriptor

var file_shared_proto_pipeline_reporter_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x14, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x56,
	0x6f, 0x69, 0x64, 0x32, 0x43, 0x0a, 0x10, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0d, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x69, 0x6c, 0x69, 0x70, 0x53, 0x6f, 0x6c, 0x69,
	0x63, 0x68, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x6b, 0x2d, 0x63, 0x69, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_proto_pipeline_reporter_proto_rawDescOnce sync.Once
	file_shared_proto_pipeline_reporter_proto_rawDescData = file_shared_proto_pipeline_reporter_proto_rawDesc
)

func file_shared_proto_pipeline_reporter_proto_rawDescGZIP() []byte {
	file_shared_proto_pipeline_reporter_proto_rawDescOnce.Do(func() {
		file_shared_proto_pipeline_reporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_proto_pipeline_reporter_proto_rawDescData)
	})
	return file_shared_proto_pipeline_reporter_proto_rawDescData
}

var file_shared_proto_pipeline_reporter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shared_proto_pipeline_reporter_proto_goTypes = []interface{}{
	(*PipelineStartRequest)(nil),  // 0: PipelineStartRequest
	(*Void)(nil),                  // 1: Void
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_shared_proto_pipeline_reporter_proto_depIdxs = []int32{
	2, // 0: PipelineStartRequest.start_time:type_name -> google.protobuf.Timestamp
	0, // 1: PipelineReporter.PipelineStart:input_type -> PipelineStartRequest
	1, // 2: PipelineReporter.PipelineStart:output_type -> Void
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shared_proto_pipeline_reporter_proto_init() }
func file_shared_proto_pipeline_reporter_proto_init() {
	if File_shared_proto_pipeline_reporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_proto_pipeline_reporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineStartRequest); i {
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
		file_shared_proto_pipeline_reporter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_shared_proto_pipeline_reporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shared_proto_pipeline_reporter_proto_goTypes,
		DependencyIndexes: file_shared_proto_pipeline_reporter_proto_depIdxs,
		MessageInfos:      file_shared_proto_pipeline_reporter_proto_msgTypes,
	}.Build()
	File_shared_proto_pipeline_reporter_proto = out.File
	file_shared_proto_pipeline_reporter_proto_rawDesc = nil
	file_shared_proto_pipeline_reporter_proto_goTypes = nil
	file_shared_proto_pipeline_reporter_proto_depIdxs = nil
}
