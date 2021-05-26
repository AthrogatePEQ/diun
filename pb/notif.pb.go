// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: notif.proto

package pb

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

type NotifTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifTestRequest) Reset() {
	*x = NotifTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifTestRequest) ProtoMessage() {}

func (x *NotifTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notif_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifTestRequest.ProtoReflect.Descriptor instead.
func (*NotifTestRequest) Descriptor() ([]byte, []int) {
	return file_notif_proto_rawDescGZIP(), []int{0}
}

type NotifTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotifTestResponse) Reset() {
	*x = NotifTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifTestResponse) ProtoMessage() {}

func (x *NotifTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notif_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifTestResponse.ProtoReflect.Descriptor instead.
func (*NotifTestResponse) Descriptor() ([]byte, []int) {
	return file_notif_proto_rawDescGZIP(), []int{1}
}

func (x *NotifTestResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_notif_proto protoreflect.FileDescriptor

var file_notif_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x12, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x4a, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x54, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x72, 0x61, 0x7a, 0x79, 0x2d, 0x6d, 0x61, 0x78, 0x2f, 0x64, 0x69, 0x75, 0x6e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notif_proto_rawDescOnce sync.Once
	file_notif_proto_rawDescData = file_notif_proto_rawDesc
)

func file_notif_proto_rawDescGZIP() []byte {
	file_notif_proto_rawDescOnce.Do(func() {
		file_notif_proto_rawDescData = protoimpl.X.CompressGZIP(file_notif_proto_rawDescData)
	})
	return file_notif_proto_rawDescData
}

var file_notif_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notif_proto_goTypes = []interface{}{
	(*NotifTestRequest)(nil),  // 0: pb.NotifTestRequest
	(*NotifTestResponse)(nil), // 1: pb.NotifTestResponse
}
var file_notif_proto_depIdxs = []int32{
	0, // 0: pb.NotifService.NotifTest:input_type -> pb.NotifTestRequest
	1, // 1: pb.NotifService.NotifTest:output_type -> pb.NotifTestResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notif_proto_init() }
func file_notif_proto_init() {
	if File_notif_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notif_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifTestRequest); i {
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
		file_notif_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifTestResponse); i {
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
			RawDescriptor: file_notif_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notif_proto_goTypes,
		DependencyIndexes: file_notif_proto_depIdxs,
		MessageInfos:      file_notif_proto_msgTypes,
	}.Build()
	File_notif_proto = out.File
	file_notif_proto_rawDesc = nil
	file_notif_proto_goTypes = nil
	file_notif_proto_depIdxs = nil
}
