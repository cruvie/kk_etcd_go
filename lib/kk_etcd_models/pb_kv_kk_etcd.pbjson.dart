// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_kv_kk_etcd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBKVDescriptor instead')
const PBKV$json = {
  '1': 'PBKV',
  '2': [
    {'1': 'Key', '3': 1, '4': 1, '5': 9, '10': 'Key'},
    {'1': 'Value', '3': 2, '4': 1, '5': 9, '10': 'Value'},
  ],
};

/// Descriptor for `PBKV`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBKVDescriptor = $convert.base64Decode(
    'CgRQQktWEhAKA0tleRgBIAEoCVIDS2V5EhQKBVZhbHVlGAIgASgJUgVWYWx1ZQ==');

@$core.Deprecated('Use pBListKVDescriptor instead')
const PBListKV$json = {
  '1': 'PBListKV',
  '2': [
    {
      '1': 'ListKV',
      '3': 1,
      '4': 3,
      '5': 11,
      '6': '.kk_etcd.PBKV',
      '10': 'ListKV'
    },
  ],
};

/// Descriptor for `PBListKV`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListKVDescriptor = $convert.base64Decode(
    'CghQQkxpc3RLVhIlCgZMaXN0S1YYASADKAsyDS5ra19ldGNkLlBCS1ZSBkxpc3RLVg==');
