// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: kk_etcd_api_hub/user/api_def/UserDelete.proto

package api_def

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// delete user
type UserDelete struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserDelete) Reset() {
	*x = UserDelete{}
	mi := &file_kk_etcd_api_hub_user_api_def_UserDelete_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDelete) ProtoMessage() {}

func (x *UserDelete) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_user_api_def_UserDelete_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDelete.ProtoReflect.Descriptor instead.
func (*UserDelete) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDescGZIP(), []int{0}
}

type UserDelete_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserDelete_Input) Reset() {
	*x = UserDelete_Input{}
	mi := &file_kk_etcd_api_hub_user_api_def_UserDelete_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDelete_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDelete_Input) ProtoMessage() {}

func (x *UserDelete_Input) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_user_api_def_UserDelete_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDelete_Input.ProtoReflect.Descriptor instead.
func (*UserDelete_Input) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UserDelete_Input) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type UserDelete_Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserDelete_Output) Reset() {
	*x = UserDelete_Output{}
	mi := &file_kk_etcd_api_hub_user_api_def_UserDelete_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDelete_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDelete_Output) ProtoMessage() {}

func (x *UserDelete_Output) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_user_api_def_UserDelete_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDelete_Output.ProtoReflect.Descriptor instead.
func (*UserDelete_Output) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDescGZIP(), []int{0, 1}
}

var File_kk_etcd_api_hub_user_api_def_UserDelete_proto protoreflect.FileDescriptor

const file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDesc = "" +
	"\n" +
	"-kk_etcd_api_hub/user/api_def/UserDelete.proto\x12\aapi_def\";\n" +
	"\n" +
	"UserDelete\x1a#\n" +
	"\x05Input\x12\x1a\n" +
	"\bUserName\x18\x01 \x01(\tR\bUserName\x1a\b\n" +
	"\x06OutputB\n" +
	"Z\b/api_defb\x06proto3"

var (
	file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDescOnce sync.Once
	file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDescData []byte
)

func file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDescGZIP() []byte {
	file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDescOnce.Do(func() {
		file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDesc), len(file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDesc)))
	})
	return file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDescData
}

var file_kk_etcd_api_hub_user_api_def_UserDelete_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kk_etcd_api_hub_user_api_def_UserDelete_proto_goTypes = []any{
	(*UserDelete)(nil),        // 0: api_def.UserDelete
	(*UserDelete_Input)(nil),  // 1: api_def.UserDelete.Input
	(*UserDelete_Output)(nil), // 2: api_def.UserDelete.Output
}
var file_kk_etcd_api_hub_user_api_def_UserDelete_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kk_etcd_api_hub_user_api_def_UserDelete_proto_init() }
func file_kk_etcd_api_hub_user_api_def_UserDelete_proto_init() {
	if File_kk_etcd_api_hub_user_api_def_UserDelete_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDesc), len(file_kk_etcd_api_hub_user_api_def_UserDelete_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kk_etcd_api_hub_user_api_def_UserDelete_proto_goTypes,
		DependencyIndexes: file_kk_etcd_api_hub_user_api_def_UserDelete_proto_depIdxs,
		MessageInfos:      file_kk_etcd_api_hub_user_api_def_UserDelete_proto_msgTypes,
	}.Build()
	File_kk_etcd_api_hub_user_api_def_UserDelete_proto = out.File
	file_kk_etcd_api_hub_user_api_def_UserDelete_proto_goTypes = nil
	file_kk_etcd_api_hub_user_api_def_UserDelete_proto_depIdxs = nil
}
