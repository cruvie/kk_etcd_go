// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: internal/mgr_hub/mgr_backup/allKVsBackup/api.proto

package allKVsBackup

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

// all kvs backup
type AllKVsBackup struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllKVsBackup) Reset() {
	*x = AllKVsBackup{}
	mi := &file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllKVsBackup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllKVsBackup) ProtoMessage() {}

func (x *AllKVsBackup) ProtoReflect() protoreflect.Message {
	mi := &file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllKVsBackup.ProtoReflect.Descriptor instead.
func (*AllKVsBackup) Descriptor() ([]byte, []int) {
	return file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDescGZIP(), []int{0}
}

type AllKVsBackup_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllKVsBackup_Input) Reset() {
	*x = AllKVsBackup_Input{}
	mi := &file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllKVsBackup_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllKVsBackup_Input) ProtoMessage() {}

func (x *AllKVsBackup_Input) ProtoReflect() protoreflect.Message {
	mi := &file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllKVsBackup_Input.ProtoReflect.Descriptor instead.
func (*AllKVsBackup_Input) Descriptor() ([]byte, []int) {
	return file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDescGZIP(), []int{0, 0}
}

type AllKVsBackup_Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	File          []byte                 `protobuf:"bytes,2,opt,name=File,proto3" json:"File,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllKVsBackup_Output) Reset() {
	*x = AllKVsBackup_Output{}
	mi := &file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllKVsBackup_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllKVsBackup_Output) ProtoMessage() {}

func (x *AllKVsBackup_Output) ProtoReflect() protoreflect.Message {
	mi := &file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllKVsBackup_Output.ProtoReflect.Descriptor instead.
func (*AllKVsBackup_Output) Descriptor() ([]byte, []int) {
	return file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AllKVsBackup_Output) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AllKVsBackup_Output) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

var File_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto protoreflect.FileDescriptor

var file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDesc = string([]byte{
	0x0a, 0x32, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x67, 0x72, 0x5f, 0x68,
	0x75, 0x62, 0x2f, 0x6d, 0x67, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x61, 0x6c,
	0x6c, 0x4b, 0x56, 0x73, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x6c, 0x6c, 0x4b, 0x56, 0x73, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x22, 0x49, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x4b, 0x56, 0x73, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x1a, 0x07, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x30, 0x0a, 0x06, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x0f, 0x5a,
	0x0d, 0x2f, 0x61, 0x6c, 0x6c, 0x4b, 0x56, 0x73, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDescOnce sync.Once
	file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDescData []byte
)

func file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDescGZIP() []byte {
	file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDescOnce.Do(func() {
		file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDesc), len(file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDesc)))
	})
	return file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDescData
}

var file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_goTypes = []any{
	(*AllKVsBackup)(nil),        // 0: allKVsBackup.AllKVsBackup
	(*AllKVsBackup_Input)(nil),  // 1: allKVsBackup.AllKVsBackup.Input
	(*AllKVsBackup_Output)(nil), // 2: allKVsBackup.AllKVsBackup.Output
}
var file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_init() }
func file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_init() {
	if File_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDesc), len(file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_goTypes,
		DependencyIndexes: file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_depIdxs,
		MessageInfos:      file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_msgTypes,
	}.Build()
	File_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto = out.File
	file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_goTypes = nil
	file_internal_mgr_hub_mgr_backup_allKVsBackup_api_proto_depIdxs = nil
}
