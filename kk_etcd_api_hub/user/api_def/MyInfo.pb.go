// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: kk_etcd_api_hub/user/api_def/MyInfo.proto

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

// get my info
type MyInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MyInfo) Reset() {
	*x = MyInfo{}
	mi := &file_kk_etcd_api_hub_user_api_def_MyInfo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyInfo) ProtoMessage() {}

func (x *MyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_user_api_def_MyInfo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyInfo.ProtoReflect.Descriptor instead.
func (*MyInfo) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDescGZIP(), []int{0}
}

type MyInfo_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MyInfo_Input) Reset() {
	*x = MyInfo_Input{}
	mi := &file_kk_etcd_api_hub_user_api_def_MyInfo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MyInfo_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyInfo_Input) ProtoMessage() {}

func (x *MyInfo_Input) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_user_api_def_MyInfo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyInfo_Input.ProtoReflect.Descriptor instead.
func (*MyInfo_Input) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDescGZIP(), []int{0, 0}
}

type MyInfo_Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Roles         []string               `protobuf:"bytes,4,rep,name=Roles,proto3" json:"Roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MyInfo_Output) Reset() {
	*x = MyInfo_Output{}
	mi := &file_kk_etcd_api_hub_user_api_def_MyInfo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MyInfo_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyInfo_Output) ProtoMessage() {}

func (x *MyInfo_Output) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_user_api_def_MyInfo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyInfo_Output.ProtoReflect.Descriptor instead.
func (*MyInfo_Output) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MyInfo_Output) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *MyInfo_Output) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_kk_etcd_api_hub_user_api_def_MyInfo_proto protoreflect.FileDescriptor

const file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDesc = "" +
	"\n" +
	")kk_etcd_api_hub/user/api_def/MyInfo.proto\x12\aapi_def\"M\n" +
	"\x06MyInfo\x1a\a\n" +
	"\x05Input\x1a:\n" +
	"\x06Output\x12\x1a\n" +
	"\bUserName\x18\x02 \x01(\tR\bUserName\x12\x14\n" +
	"\x05Roles\x18\x04 \x03(\tR\x05RolesB\n" +
	"Z\b/api_defb\x06proto3"

var (
	file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDescOnce sync.Once
	file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDescData []byte
)

func file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDescGZIP() []byte {
	file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDescOnce.Do(func() {
		file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDesc), len(file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDesc)))
	})
	return file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDescData
}

var file_kk_etcd_api_hub_user_api_def_MyInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kk_etcd_api_hub_user_api_def_MyInfo_proto_goTypes = []any{
	(*MyInfo)(nil),        // 0: api_def.MyInfo
	(*MyInfo_Input)(nil),  // 1: api_def.MyInfo.Input
	(*MyInfo_Output)(nil), // 2: api_def.MyInfo.Output
}
var file_kk_etcd_api_hub_user_api_def_MyInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kk_etcd_api_hub_user_api_def_MyInfo_proto_init() }
func file_kk_etcd_api_hub_user_api_def_MyInfo_proto_init() {
	if File_kk_etcd_api_hub_user_api_def_MyInfo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDesc), len(file_kk_etcd_api_hub_user_api_def_MyInfo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kk_etcd_api_hub_user_api_def_MyInfo_proto_goTypes,
		DependencyIndexes: file_kk_etcd_api_hub_user_api_def_MyInfo_proto_depIdxs,
		MessageInfos:      file_kk_etcd_api_hub_user_api_def_MyInfo_proto_msgTypes,
	}.Build()
	File_kk_etcd_api_hub_user_api_def_MyInfo_proto = out.File
	file_kk_etcd_api_hub_user_api_def_MyInfo_proto_goTypes = nil
	file_kk_etcd_api_hub_user_api_def_MyInfo_proto_depIdxs = nil
}
