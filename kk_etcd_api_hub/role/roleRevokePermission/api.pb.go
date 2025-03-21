// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: kk_etcd_api_hub/role/roleRevokePermission/api.proto

package roleRevokePermission

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

// revoke permission
type RoleRevokePermission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleRevokePermission) Reset() {
	*x = RoleRevokePermission{}
	mi := &file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleRevokePermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleRevokePermission) ProtoMessage() {}

func (x *RoleRevokePermission) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleRevokePermission.ProtoReflect.Descriptor instead.
func (*RoleRevokePermission) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDescGZIP(), []int{0}
}

type RoleRevokePermission_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Key           string                 `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	RangeEnd      string                 `protobuf:"bytes,3,opt,name=RangeEnd,proto3" json:"RangeEnd,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleRevokePermission_Input) Reset() {
	*x = RoleRevokePermission_Input{}
	mi := &file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleRevokePermission_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleRevokePermission_Input) ProtoMessage() {}

func (x *RoleRevokePermission_Input) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleRevokePermission_Input.ProtoReflect.Descriptor instead.
func (*RoleRevokePermission_Input) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RoleRevokePermission_Input) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleRevokePermission_Input) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RoleRevokePermission_Input) GetRangeEnd() string {
	if x != nil {
		return x.RangeEnd
	}
	return ""
}

type RoleRevokePermission_Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleRevokePermission_Output) Reset() {
	*x = RoleRevokePermission_Output{}
	mi := &file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleRevokePermission_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleRevokePermission_Output) ProtoMessage() {}

func (x *RoleRevokePermission_Output) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleRevokePermission_Output.ProtoReflect.Descriptor instead.
func (*RoleRevokePermission_Output) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDescGZIP(), []int{0, 1}
}

var File_kk_etcd_api_hub_role_roleRevokePermission_api_proto protoreflect.FileDescriptor

var file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDesc = string([]byte{
	0x0a, 0x33, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x75,
	0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6b, 0x0a, 0x14, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x49, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b,
	0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x1a, 0x08,
	0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x17, 0x5a, 0x15, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDescOnce sync.Once
	file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDescData []byte
)

func file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDescGZIP() []byte {
	file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDescOnce.Do(func() {
		file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDesc), len(file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDesc)))
	})
	return file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDescData
}

var file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_goTypes = []any{
	(*RoleRevokePermission)(nil),        // 0: roleRevokePermission.RoleRevokePermission
	(*RoleRevokePermission_Input)(nil),  // 1: roleRevokePermission.RoleRevokePermission.Input
	(*RoleRevokePermission_Output)(nil), // 2: roleRevokePermission.RoleRevokePermission.Output
}
var file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_init() }
func file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_init() {
	if File_kk_etcd_api_hub_role_roleRevokePermission_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDesc), len(file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_goTypes,
		DependencyIndexes: file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_depIdxs,
		MessageInfos:      file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_msgTypes,
	}.Build()
	File_kk_etcd_api_hub_role_roleRevokePermission_api_proto = out.File
	file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_goTypes = nil
	file_kk_etcd_api_hub_role_roleRevokePermission_api_proto_depIdxs = nil
}
