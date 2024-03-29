// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: solpipe.proto

package solpipe

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solpipe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_solpipe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_solpipe_proto_rawDescGZIP(), []int{0}
}

type CapacityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capacity float32 `protobuf:"fixed32,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
}

func (x *CapacityResponse) Reset() {
	*x = CapacityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solpipe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapacityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapacityResponse) ProtoMessage() {}

func (x *CapacityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_solpipe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapacityResponse.ProtoReflect.Descriptor instead.
func (*CapacityResponse) Descriptor() ([]byte, []int) {
	return file_solpipe_proto_rawDescGZIP(), []int{1}
}

func (x *CapacityResponse) GetCapacity() float32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

var File_solpipe_proto protoreflect.FileDescriptor

var file_solpipe_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x6f, 0x6c, 0x70, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x6f, 0x6c, 0x70, 0x69, 0x70, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x2e, 0x0a, 0x10, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x32, 0x47, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x37, 0x0a, 0x08, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x2e,
	0x73, 0x6f, 0x6c, 0x70, 0x69, 0x70, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e,
	0x73, 0x6f, 0x6c, 0x70, 0x69, 0x70, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x61,
	0x64, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x70, 0x69, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_solpipe_proto_rawDescOnce sync.Once
	file_solpipe_proto_rawDescData = file_solpipe_proto_rawDesc
)

func file_solpipe_proto_rawDescGZIP() []byte {
	file_solpipe_proto_rawDescOnce.Do(func() {
		file_solpipe_proto_rawDescData = protoimpl.X.CompressGZIP(file_solpipe_proto_rawDescData)
	})
	return file_solpipe_proto_rawDescData
}

var file_solpipe_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_solpipe_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: solpipe.Empty
	(*CapacityResponse)(nil), // 1: solpipe.CapacityResponse
}
var file_solpipe_proto_depIdxs = []int32{
	0, // 0: solpipe.WorkerStatus.OnStatus:input_type -> solpipe.Empty
	1, // 1: solpipe.WorkerStatus.OnStatus:output_type -> solpipe.CapacityResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_solpipe_proto_init() }
func file_solpipe_proto_init() {
	if File_solpipe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_solpipe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_solpipe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapacityResponse); i {
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
			RawDescriptor: file_solpipe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_solpipe_proto_goTypes,
		DependencyIndexes: file_solpipe_proto_depIdxs,
		MessageInfos:      file_solpipe_proto_msgTypes,
	}.Build()
	File_solpipe_proto = out.File
	file_solpipe_proto_rawDesc = nil
	file_solpipe_proto_goTypes = nil
	file_solpipe_proto_depIdxs = nil
}
