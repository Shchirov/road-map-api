// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user-service/post.proto

package api

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

type PostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PostResponse) Reset() {
	*x = PostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponse) ProtoMessage() {}

func (x *PostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponse.ProtoReflect.Descriptor instead.
func (*PostResponse) Descriptor() ([]byte, []int) {
	return file_user_service_post_proto_rawDescGZIP(), []int{0}
}

func (x *PostResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PostResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_user_service_post_proto protoreflect.FileDescriptor

var file_user_service_post_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22,
	0x54, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x12, 0x5a, 0x10, 0x72, 0x6f, 0x61, 0x64, 0x2d, 0x6d, 0x61,
	0x70, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_service_post_proto_rawDescOnce sync.Once
	file_user_service_post_proto_rawDescData = file_user_service_post_proto_rawDesc
)

func file_user_service_post_proto_rawDescGZIP() []byte {
	file_user_service_post_proto_rawDescOnce.Do(func() {
		file_user_service_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_post_proto_rawDescData)
	})
	return file_user_service_post_proto_rawDescData
}

var file_user_service_post_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_service_post_proto_goTypes = []interface{}{
	(*PostResponse)(nil), // 0: post.PostResponse
}
var file_user_service_post_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_post_proto_init() }
func file_user_service_post_proto_init() {
	if File_user_service_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_service_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponse); i {
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
			RawDescriptor: file_user_service_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_service_post_proto_goTypes,
		DependencyIndexes: file_user_service_post_proto_depIdxs,
		MessageInfos:      file_user_service_post_proto_msgTypes,
	}.Build()
	File_user_service_post_proto = out.File
	file_user_service_post_proto_rawDesc = nil
	file_user_service_post_proto_goTypes = nil
	file_user_service_post_proto_depIdxs = nil
}
