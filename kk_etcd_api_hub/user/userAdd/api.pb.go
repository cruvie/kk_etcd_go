// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: kk_etcd_api_hub/user/userAdd/api.proto

package userAdd

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

// add user
type UserAdd struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserAdd) Reset() {
	*x = UserAdd{}
	mi := &file_kk_etcd_api_hub_user_userAdd_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdd) ProtoMessage() {}

func (x *UserAdd) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_user_userAdd_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdd.ProtoReflect.Descriptor instead.
func (*UserAdd) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_user_userAdd_api_proto_rawDescGZIP(), []int{0}
}

type UserAdd_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Roles         []string               `protobuf:"bytes,4,rep,name=Roles,proto3" json:"Roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserAdd_Input) Reset() {
	*x = UserAdd_Input{}
	mi := &file_kk_etcd_api_hub_user_userAdd_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAdd_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdd_Input) ProtoMessage() {}

func (x *UserAdd_Input) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_user_userAdd_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdd_Input.ProtoReflect.Descriptor instead.
func (*UserAdd_Input) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_user_userAdd_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UserAdd_Input) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserAdd_Input) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserAdd_Input) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UserAdd_Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserAdd_Output) Reset() {
	*x = UserAdd_Output{}
	mi := &file_kk_etcd_api_hub_user_userAdd_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAdd_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdd_Output) ProtoMessage() {}

func (x *UserAdd_Output) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_user_userAdd_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdd_Output.ProtoReflect.Descriptor instead.
func (*UserAdd_Output) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_user_userAdd_api_proto_rawDescGZIP(), []int{0, 1}
}

var File_kk_etcd_api_hub_user_userAdd_api_proto protoreflect.FileDescriptor

const file_kk_etcd_api_hub_user_userAdd_api_proto_rawDesc = "" +
	"\n" +
	"&kk_etcd_api_hub/user/userAdd/api.proto\x12\auserAdd\"j\n" +
	"\aUserAdd\x1aU\n" +
	"\x05Input\x12\x1a\n" +
	"\bUserName\x18\x02 \x01(\tR\bUserName\x12\x1a\n" +
	"\bPassword\x18\x03 \x01(\tR\bPassword\x12\x14\n" +
	"\x05Roles\x18\x04 \x03(\tR\x05Roles\x1a\b\n" +
	"\x06OutputB\n" +
	"Z\b/userAddb\x06proto3"

var (
	file_kk_etcd_api_hub_user_userAdd_api_proto_rawDescOnce sync.Once
	file_kk_etcd_api_hub_user_userAdd_api_proto_rawDescData []byte
)

func file_kk_etcd_api_hub_user_userAdd_api_proto_rawDescGZIP() []byte {
	file_kk_etcd_api_hub_user_userAdd_api_proto_rawDescOnce.Do(func() {
		file_kk_etcd_api_hub_user_userAdd_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_user_userAdd_api_proto_rawDesc), len(file_kk_etcd_api_hub_user_userAdd_api_proto_rawDesc)))
	})
	return file_kk_etcd_api_hub_user_userAdd_api_proto_rawDescData
}

var file_kk_etcd_api_hub_user_userAdd_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kk_etcd_api_hub_user_userAdd_api_proto_goTypes = []any{
	(*UserAdd)(nil),        // 0: userAdd.UserAdd
	(*UserAdd_Input)(nil),  // 1: userAdd.UserAdd.Input
	(*UserAdd_Output)(nil), // 2: userAdd.UserAdd.Output
}
var file_kk_etcd_api_hub_user_userAdd_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kk_etcd_api_hub_user_userAdd_api_proto_init() }
func file_kk_etcd_api_hub_user_userAdd_api_proto_init() {
	if File_kk_etcd_api_hub_user_userAdd_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_user_userAdd_api_proto_rawDesc), len(file_kk_etcd_api_hub_user_userAdd_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kk_etcd_api_hub_user_userAdd_api_proto_goTypes,
		DependencyIndexes: file_kk_etcd_api_hub_user_userAdd_api_proto_depIdxs,
		MessageInfos:      file_kk_etcd_api_hub_user_userAdd_api_proto_msgTypes,
	}.Build()
	File_kk_etcd_api_hub_user_userAdd_api_proto = out.File
	file_kk_etcd_api_hub_user_userAdd_api_proto_goTypes = nil
	file_kk_etcd_api_hub_user_userAdd_api_proto_depIdxs = nil
}
