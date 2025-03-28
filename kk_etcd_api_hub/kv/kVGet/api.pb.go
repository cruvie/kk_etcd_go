// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: kk_etcd_api_hub/kv/kVGet/api.proto

package kVGet

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

// get kv
type KVGet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KVGet) Reset() {
	*x = KVGet{}
	mi := &file_kk_etcd_api_hub_kv_kVGet_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KVGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVGet) ProtoMessage() {}

func (x *KVGet) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_kv_kVGet_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVGet.ProtoReflect.Descriptor instead.
func (*KVGet) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDescGZIP(), []int{0}
}

type KVGet_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KVGet_Input) Reset() {
	*x = KVGet_Input{}
	mi := &file_kk_etcd_api_hub_kv_kVGet_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KVGet_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVGet_Input) ProtoMessage() {}

func (x *KVGet_Input) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_kv_kVGet_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVGet_Input.ProtoReflect.Descriptor instead.
func (*KVGet_Input) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *KVGet_Input) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type KVGet_Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	KV            *kk_etcd_models.PBKV   `protobuf:"bytes,1,opt,name=KV,proto3" json:"KV,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KVGet_Output) Reset() {
	*x = KVGet_Output{}
	mi := &file_kk_etcd_api_hub_kv_kVGet_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KVGet_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVGet_Output) ProtoMessage() {}

func (x *KVGet_Output) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_kv_kVGet_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVGet_Output.ProtoReflect.Descriptor instead.
func (*KVGet_Output) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDescGZIP(), []int{0, 1}
}

func (x *KVGet_Output) GetKV() *kk_etcd_models.PBKV {
	if x != nil {
		return x.KV
	}
	return nil
}

var File_kk_etcd_api_hub_kv_kVGet_api_proto protoreflect.FileDescriptor

const file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDesc = "" +
	"\n" +
	"\"kk_etcd_api_hub/kv/kVGet/api.proto\x12\x05kVGet\x1a\"kk_etcd_models/pb_kv_kk_etcd.proto\"R\n" +
	"\x05KVGet\x1a\x19\n" +
	"\x05Input\x12\x10\n" +
	"\x03Key\x18\x01 \x01(\tR\x03Key\x1a.\n" +
	"\x06Output\x12$\n" +
	"\x02KV\x18\x01 \x01(\v2\x14.kk_etcd_models.PBKVR\x02KVB\bZ\x06/kVGetb\x06proto3"

var (
	file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDescOnce sync.Once
	file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDescData []byte
)

func file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDescGZIP() []byte {
	file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDescOnce.Do(func() {
		file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDesc), len(file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDesc)))
	})
	return file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDescData
}

var file_kk_etcd_api_hub_kv_kVGet_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kk_etcd_api_hub_kv_kVGet_api_proto_goTypes = []any{
	(*KVGet)(nil),               // 0: kVGet.KVGet
	(*KVGet_Input)(nil),         // 1: kVGet.KVGet.Input
	(*KVGet_Output)(nil),        // 2: kVGet.KVGet.Output
	(*kk_etcd_models.PBKV)(nil), // 3: kk_etcd_models.PBKV
}
var file_kk_etcd_api_hub_kv_kVGet_api_proto_depIdxs = []int32{
	3, // 0: kVGet.KVGet.Output.KV:type_name -> kk_etcd_models.PBKV
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kk_etcd_api_hub_kv_kVGet_api_proto_init() }
func file_kk_etcd_api_hub_kv_kVGet_api_proto_init() {
	if File_kk_etcd_api_hub_kv_kVGet_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDesc), len(file_kk_etcd_api_hub_kv_kVGet_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kk_etcd_api_hub_kv_kVGet_api_proto_goTypes,
		DependencyIndexes: file_kk_etcd_api_hub_kv_kVGet_api_proto_depIdxs,
		MessageInfos:      file_kk_etcd_api_hub_kv_kVGet_api_proto_msgTypes,
	}.Build()
	File_kk_etcd_api_hub_kv_kVGet_api_proto = out.File
	file_kk_etcd_api_hub_kv_kVGet_api_proto_goTypes = nil
	file_kk_etcd_api_hub_kv_kVGet_api_proto_depIdxs = nil
}
