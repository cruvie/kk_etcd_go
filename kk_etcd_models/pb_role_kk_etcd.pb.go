// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: kk_etcd_models/pb_role_kk_etcd.proto

package kk_etcd_models

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

type PBRole struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Perms         []*PBPermission        `protobuf:"bytes,2,rep,name=Perms,proto3" json:"Perms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PBRole) Reset() {
	*x = PBRole{}
	mi := &file_kk_etcd_models_pb_role_kk_etcd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PBRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBRole) ProtoMessage() {}

func (x *PBRole) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_models_pb_role_kk_etcd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBRole.ProtoReflect.Descriptor instead.
func (*PBRole) Descriptor() ([]byte, []int) {
	return file_kk_etcd_models_pb_role_kk_etcd_proto_rawDescGZIP(), []int{0}
}

func (x *PBRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PBRole) GetPerms() []*PBPermission {
	if x != nil {
		return x.Perms
	}
	return nil
}

type PBPermission struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Key      string                 `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	RangeEnd string                 `protobuf:"bytes,3,opt,name=RangeEnd,proto3" json:"RangeEnd,omitempty"`
	// authpb.READ 0
	// authpb.WRITE 1
	// authpb.READWRITE 2
	PermissionType int32 `protobuf:"varint,4,opt,name=PermissionType,proto3" json:"PermissionType,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PBPermission) Reset() {
	*x = PBPermission{}
	mi := &file_kk_etcd_models_pb_role_kk_etcd_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PBPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBPermission) ProtoMessage() {}

func (x *PBPermission) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_models_pb_role_kk_etcd_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBPermission.ProtoReflect.Descriptor instead.
func (*PBPermission) Descriptor() ([]byte, []int) {
	return file_kk_etcd_models_pb_role_kk_etcd_proto_rawDescGZIP(), []int{1}
}

func (x *PBPermission) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PBPermission) GetRangeEnd() string {
	if x != nil {
		return x.RangeEnd
	}
	return ""
}

func (x *PBPermission) GetPermissionType() int32 {
	if x != nil {
		return x.PermissionType
	}
	return 0
}

type PBListRole struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*PBRole              `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PBListRole) Reset() {
	*x = PBListRole{}
	mi := &file_kk_etcd_models_pb_role_kk_etcd_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PBListRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBListRole) ProtoMessage() {}

func (x *PBListRole) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_models_pb_role_kk_etcd_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBListRole.ProtoReflect.Descriptor instead.
func (*PBListRole) Descriptor() ([]byte, []int) {
	return file_kk_etcd_models_pb_role_kk_etcd_proto_rawDescGZIP(), []int{2}
}

func (x *PBListRole) GetList() []*PBRole {
	if x != nil {
		return x.List
	}
	return nil
}

var File_kk_etcd_models_pb_role_kk_etcd_proto protoreflect.FileDescriptor

var file_kk_etcd_models_pb_role_kk_etcd_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x70, 0x62, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x50, 0x0a, 0x06, 0x50, 0x42, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x50, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x42, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x50, 0x65, 0x72, 0x6d, 0x73, 0x22, 0x64, 0x0a, 0x0c, 0x50, 0x42, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x38,
	0x0a, 0x0a, 0x50, 0x42, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x6b, 0x5f,
	0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x42, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x75, 0x76, 0x69, 0x65, 0x2f, 0x6b, 0x6b,
	0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x67, 0x6f, 0x2f, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_kk_etcd_models_pb_role_kk_etcd_proto_rawDescOnce sync.Once
	file_kk_etcd_models_pb_role_kk_etcd_proto_rawDescData []byte
)

func file_kk_etcd_models_pb_role_kk_etcd_proto_rawDescGZIP() []byte {
	file_kk_etcd_models_pb_role_kk_etcd_proto_rawDescOnce.Do(func() {
		file_kk_etcd_models_pb_role_kk_etcd_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kk_etcd_models_pb_role_kk_etcd_proto_rawDesc), len(file_kk_etcd_models_pb_role_kk_etcd_proto_rawDesc)))
	})
	return file_kk_etcd_models_pb_role_kk_etcd_proto_rawDescData
}

var file_kk_etcd_models_pb_role_kk_etcd_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kk_etcd_models_pb_role_kk_etcd_proto_goTypes = []any{
	(*PBRole)(nil),       // 0: kk_etcd_models.PBRole
	(*PBPermission)(nil), // 1: kk_etcd_models.PBPermission
	(*PBListRole)(nil),   // 2: kk_etcd_models.PBListRole
}
var file_kk_etcd_models_pb_role_kk_etcd_proto_depIdxs = []int32{
	1, // 0: kk_etcd_models.PBRole.Perms:type_name -> kk_etcd_models.PBPermission
	0, // 1: kk_etcd_models.PBListRole.List:type_name -> kk_etcd_models.PBRole
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kk_etcd_models_pb_role_kk_etcd_proto_init() }
func file_kk_etcd_models_pb_role_kk_etcd_proto_init() {
	if File_kk_etcd_models_pb_role_kk_etcd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kk_etcd_models_pb_role_kk_etcd_proto_rawDesc), len(file_kk_etcd_models_pb_role_kk_etcd_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kk_etcd_models_pb_role_kk_etcd_proto_goTypes,
		DependencyIndexes: file_kk_etcd_models_pb_role_kk_etcd_proto_depIdxs,
		MessageInfos:      file_kk_etcd_models_pb_role_kk_etcd_proto_msgTypes,
	}.Build()
	File_kk_etcd_models_pb_role_kk_etcd_proto = out.File
	file_kk_etcd_models_pb_role_kk_etcd_proto_goTypes = nil
	file_kk_etcd_models_pb_role_kk_etcd_proto_depIdxs = nil
}
