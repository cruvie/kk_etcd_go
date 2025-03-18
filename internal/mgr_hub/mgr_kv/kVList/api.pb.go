// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: internal/mgr_hub/mgr_kv/kVList/api.proto

package kVList

import (
	kk_etcd_models "github.com/cruvie/kk_etcd_go/kk_etcd_models"
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

// list kv
type KVList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KVList) Reset() {
	*x = KVList{}
	mi := &file_internal_mgr_hub_mgr_kv_kVList_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KVList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVList) ProtoMessage() {}

func (x *KVList) ProtoReflect() protoreflect.Message {
	mi := &file_internal_mgr_hub_mgr_kv_kVList_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVList.ProtoReflect.Descriptor instead.
func (*KVList) Descriptor() ([]byte, []int) {
	return file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDescGZIP(), []int{0}
}

type KVList_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Prefix        string                 `protobuf:"bytes,1,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KVList_Input) Reset() {
	*x = KVList_Input{}
	mi := &file_internal_mgr_hub_mgr_kv_kVList_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KVList_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVList_Input) ProtoMessage() {}

func (x *KVList_Input) ProtoReflect() protoreflect.Message {
	mi := &file_internal_mgr_hub_mgr_kv_kVList_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVList_Input.ProtoReflect.Descriptor instead.
func (*KVList_Input) Descriptor() ([]byte, []int) {
	return file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *KVList_Input) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type KVList_Output struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	KVList        *kk_etcd_models.PBListKV `protobuf:"bytes,1,opt,name=KVList,proto3" json:"KVList,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KVList_Output) Reset() {
	*x = KVList_Output{}
	mi := &file_internal_mgr_hub_mgr_kv_kVList_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KVList_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVList_Output) ProtoMessage() {}

func (x *KVList_Output) ProtoReflect() protoreflect.Message {
	mi := &file_internal_mgr_hub_mgr_kv_kVList_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVList_Output.ProtoReflect.Descriptor instead.
func (*KVList_Output) Descriptor() ([]byte, []int) {
	return file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDescGZIP(), []int{0, 1}
}

func (x *KVList_Output) GetKVList() *kk_etcd_models.PBListKV {
	if x != nil {
		return x.KVList
	}
	return nil
}

var File_internal_mgr_hub_mgr_kv_kVList_api_proto protoreflect.FileDescriptor

var file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x67, 0x72, 0x5f, 0x68,
	0x75, 0x62, 0x2f, 0x6d, 0x67, 0x72, 0x5f, 0x6b, 0x76, 0x2f, 0x6b, 0x56, 0x4c, 0x69, 0x73, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6b, 0x56, 0x4c, 0x69,
	0x73, 0x74, 0x1a, 0x22, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x70, 0x62, 0x5f, 0x6b, 0x76, 0x5f, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x06, 0x4b, 0x56, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x1f, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x1a, 0x3a, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x4b,
	0x56, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6b,
	0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x42, 0x4c,
	0x69, 0x73, 0x74, 0x4b, 0x56, 0x52, 0x06, 0x4b, 0x56, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x5a,
	0x07, 0x2f, 0x6b, 0x56, 0x4c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDescOnce sync.Once
	file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDescData []byte
)

func file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDescGZIP() []byte {
	file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDescOnce.Do(func() {
		file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDesc), len(file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDesc)))
	})
	return file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDescData
}

var file_internal_mgr_hub_mgr_kv_kVList_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_mgr_hub_mgr_kv_kVList_api_proto_goTypes = []any{
	(*KVList)(nil),                  // 0: kVList.KVList
	(*KVList_Input)(nil),            // 1: kVList.KVList.Input
	(*KVList_Output)(nil),           // 2: kVList.KVList.Output
	(*kk_etcd_models.PBListKV)(nil), // 3: kk_etcd_models.PBListKV
}
var file_internal_mgr_hub_mgr_kv_kVList_api_proto_depIdxs = []int32{
	3, // 0: kVList.KVList.Output.KVList:type_name -> kk_etcd_models.PBListKV
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_mgr_hub_mgr_kv_kVList_api_proto_init() }
func file_internal_mgr_hub_mgr_kv_kVList_api_proto_init() {
	if File_internal_mgr_hub_mgr_kv_kVList_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDesc), len(file_internal_mgr_hub_mgr_kv_kVList_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_mgr_hub_mgr_kv_kVList_api_proto_goTypes,
		DependencyIndexes: file_internal_mgr_hub_mgr_kv_kVList_api_proto_depIdxs,
		MessageInfos:      file_internal_mgr_hub_mgr_kv_kVList_api_proto_msgTypes,
	}.Build()
	File_internal_mgr_hub_mgr_kv_kVList_api_proto = out.File
	file_internal_mgr_hub_mgr_kv_kVList_api_proto_goTypes = nil
	file_internal_mgr_hub_mgr_kv_kVList_api_proto_depIdxs = nil
}
