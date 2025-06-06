// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: kk_etcd_api_hub/role/api_def/RoleDelete.proto

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

// delete role
type RoleDelete struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleDelete) Reset() {
	*x = RoleDelete{}
	mi := &file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDelete) ProtoMessage() {}

func (x *RoleDelete) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDelete.ProtoReflect.Descriptor instead.
func (*RoleDelete) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDescGZIP(), []int{0}
}

type RoleDelete_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleDelete_Input) Reset() {
	*x = RoleDelete_Input{}
	mi := &file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleDelete_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDelete_Input) ProtoMessage() {}

func (x *RoleDelete_Input) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDelete_Input.ProtoReflect.Descriptor instead.
func (*RoleDelete_Input) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RoleDelete_Input) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RoleDelete_Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleDelete_Output) Reset() {
	*x = RoleDelete_Output{}
	mi := &file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleDelete_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDelete_Output) ProtoMessage() {}

func (x *RoleDelete_Output) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDelete_Output.ProtoReflect.Descriptor instead.
func (*RoleDelete_Output) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDescGZIP(), []int{0, 1}
}

var File_kk_etcd_api_hub_role_api_def_RoleDelete_proto protoreflect.FileDescriptor

const file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDesc = "" +
	"\n" +
	"-kk_etcd_api_hub/role/api_def/RoleDelete.proto\x12\aapi_def\"3\n" +
	"\n" +
	"RoleDelete\x1a\x1b\n" +
	"\x05Input\x12\x12\n" +
	"\x04Name\x18\x01 \x01(\tR\x04Name\x1a\b\n" +
	"\x06OutputB\n" +
	"Z\b/api_defb\x06proto3"

var (
	file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDescOnce sync.Once
	file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDescData []byte
)

func file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDescGZIP() []byte {
	file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDescOnce.Do(func() {
		file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDesc), len(file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDesc)))
	})
	return file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDescData
}

var file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_goTypes = []any{
	(*RoleDelete)(nil),        // 0: api_def.RoleDelete
	(*RoleDelete_Input)(nil),  // 1: api_def.RoleDelete.Input
	(*RoleDelete_Output)(nil), // 2: api_def.RoleDelete.Output
}
var file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_init() }
func file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_init() {
	if File_kk_etcd_api_hub_role_api_def_RoleDelete_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDesc), len(file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_goTypes,
		DependencyIndexes: file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_depIdxs,
		MessageInfos:      file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_msgTypes,
	}.Build()
	File_kk_etcd_api_hub_role_api_def_RoleDelete_proto = out.File
	file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_goTypes = nil
	file_kk_etcd_api_hub_role_api_def_RoleDelete_proto_depIdxs = nil
}
