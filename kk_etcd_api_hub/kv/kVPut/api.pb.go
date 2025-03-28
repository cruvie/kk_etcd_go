// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: kk_etcd_api_hub/kv/kVPut/api.proto

package kVPut

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

// put kv
type KVPut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KVPut) Reset() {
	*x = KVPut{}
	mi := &file_kk_etcd_api_hub_kv_kVPut_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KVPut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVPut) ProtoMessage() {}

func (x *KVPut) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_kv_kVPut_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVPut.ProtoReflect.Descriptor instead.
func (*KVPut) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDescGZIP(), []int{0}
}

type KVPut_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KVPut_Input) Reset() {
	*x = KVPut_Input{}
	mi := &file_kk_etcd_api_hub_kv_kVPut_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KVPut_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVPut_Input) ProtoMessage() {}

func (x *KVPut_Input) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_kv_kVPut_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVPut_Input.ProtoReflect.Descriptor instead.
func (*KVPut_Input) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *KVPut_Input) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVPut_Input) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type KVPut_Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KVPut_Output) Reset() {
	*x = KVPut_Output{}
	mi := &file_kk_etcd_api_hub_kv_kVPut_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KVPut_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVPut_Output) ProtoMessage() {}

func (x *KVPut_Output) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_api_hub_kv_kVPut_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVPut_Output.ProtoReflect.Descriptor instead.
func (*KVPut_Output) Descriptor() ([]byte, []int) {
	return file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDescGZIP(), []int{0, 1}
}

var File_kk_etcd_api_hub_kv_kVPut_api_proto protoreflect.FileDescriptor

const file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDesc = "" +
	"\n" +
	"\"kk_etcd_api_hub/kv/kVPut/api.proto\x12\x05kVPut\"B\n" +
	"\x05KVPut\x1a/\n" +
	"\x05Input\x12\x10\n" +
	"\x03Key\x18\x01 \x01(\tR\x03Key\x12\x14\n" +
	"\x05Value\x18\x02 \x01(\tR\x05Value\x1a\b\n" +
	"\x06OutputB\bZ\x06/kVPutb\x06proto3"

var (
	file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDescOnce sync.Once
	file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDescData []byte
)

func file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDescGZIP() []byte {
	file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDescOnce.Do(func() {
		file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDesc), len(file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDesc)))
	})
	return file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDescData
}

var file_kk_etcd_api_hub_kv_kVPut_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kk_etcd_api_hub_kv_kVPut_api_proto_goTypes = []any{
	(*KVPut)(nil),        // 0: kVPut.KVPut
	(*KVPut_Input)(nil),  // 1: kVPut.KVPut.Input
	(*KVPut_Output)(nil), // 2: kVPut.KVPut.Output
}
var file_kk_etcd_api_hub_kv_kVPut_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kk_etcd_api_hub_kv_kVPut_api_proto_init() }
func file_kk_etcd_api_hub_kv_kVPut_api_proto_init() {
	if File_kk_etcd_api_hub_kv_kVPut_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDesc), len(file_kk_etcd_api_hub_kv_kVPut_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kk_etcd_api_hub_kv_kVPut_api_proto_goTypes,
		DependencyIndexes: file_kk_etcd_api_hub_kv_kVPut_api_proto_depIdxs,
		MessageInfos:      file_kk_etcd_api_hub_kv_kVPut_api_proto_msgTypes,
	}.Build()
	File_kk_etcd_api_hub_kv_kVPut_api_proto = out.File
	file_kk_etcd_api_hub_kv_kVPut_api_proto_goTypes = nil
	file_kk_etcd_api_hub_kv_kVPut_api_proto_depIdxs = nil
}
