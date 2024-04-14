//
//  Generated code. Do not modify.
//  source: api_kv.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use kVPutParamDescriptor instead')
const KVPutParam$json = {
  '1': 'KVPutParam',
  '2': [
    {'1': 'Key', '3': 1, '4': 1, '5': 9, '10': 'Key'},
    {'1': 'Value', '3': 2, '4': 1, '5': 9, '10': 'Value'},
  ],
};

/// Descriptor for `KVPutParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVPutParamDescriptor = $convert.base64Decode(
    'CgpLVlB1dFBhcmFtEhAKA0tleRgBIAEoCVIDS2V5EhQKBVZhbHVlGAIgASgJUgVWYWx1ZQ==');

@$core.Deprecated('Use kVPutResponseDescriptor instead')
const KVPutResponse$json = {
  '1': 'KVPutResponse',
};

/// Descriptor for `KVPutResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVPutResponseDescriptor = $convert.base64Decode(
    'Cg1LVlB1dFJlc3BvbnNl');

@$core.Deprecated('Use kVGetParamDescriptor instead')
const KVGetParam$json = {
  '1': 'KVGetParam',
  '2': [
    {'1': 'Key', '3': 1, '4': 1, '5': 9, '10': 'Key'},
  ],
};

/// Descriptor for `KVGetParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVGetParamDescriptor = $convert.base64Decode(
    'CgpLVkdldFBhcmFtEhAKA0tleRgBIAEoCVIDS2V5');

@$core.Deprecated('Use kVGetResponseDescriptor instead')
const KVGetResponse$json = {
  '1': 'KVGetResponse',
  '2': [
    {'1': 'KV', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBKV', '10': 'KV'},
  ],
};

/// Descriptor for `KVGetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVGetResponseDescriptor = $convert.base64Decode(
    'Cg1LVkdldFJlc3BvbnNlEiQKAktWGAEgASgLMhQua2tfZXRjZF9tb2RlbHMuUEJLVlICS1Y=');

@$core.Deprecated('Use kVListParamDescriptor instead')
const KVListParam$json = {
  '1': 'KVListParam',
  '2': [
    {'1': 'Prefix', '3': 1, '4': 1, '5': 9, '10': 'Prefix'},
  ],
};

/// Descriptor for `KVListParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVListParamDescriptor = $convert.base64Decode(
    'CgtLVkxpc3RQYXJhbRIWCgZQcmVmaXgYASABKAlSBlByZWZpeA==');

@$core.Deprecated('Use kVListResponseDescriptor instead')
const KVListResponse$json = {
  '1': 'KVListResponse',
  '2': [
    {'1': 'KVList', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBListKV', '10': 'KVList'},
  ],
};

/// Descriptor for `KVListResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVListResponseDescriptor = $convert.base64Decode(
    'Cg5LVkxpc3RSZXNwb25zZRIwCgZLVkxpc3QYASABKAsyGC5ra19ldGNkX21vZGVscy5QQkxpc3'
    'RLVlIGS1ZMaXN0');

@$core.Deprecated('Use kVDelParamDescriptor instead')
const KVDelParam$json = {
  '1': 'KVDelParam',
  '2': [
    {'1': 'Key', '3': 1, '4': 1, '5': 9, '10': 'Key'},
  ],
};

/// Descriptor for `KVDelParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVDelParamDescriptor = $convert.base64Decode(
    'CgpLVkRlbFBhcmFtEhAKA0tleRgBIAEoCVIDS2V5');

@$core.Deprecated('Use kVDelResponseDescriptor instead')
const KVDelResponse$json = {
  '1': 'KVDelResponse',
};

/// Descriptor for `KVDelResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVDelResponseDescriptor = $convert.base64Decode(
    'Cg1LVkRlbFJlc3BvbnNl');

