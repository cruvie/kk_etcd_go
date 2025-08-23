// This is a generated file - do not edit.
//
// Generated from internal/service_hub/kv/api_def/KVGet.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use kVGetDescriptor instead')
const KVGet$json = {
  '1': 'KVGet',
  '3': [KVGet_Input$json, KVGet_Output$json],
};

@$core.Deprecated('Use kVGetDescriptor instead')
const KVGet_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Key', '3': 1, '4': 1, '5': 9, '10': 'Key'},
  ],
};

@$core.Deprecated('Use kVGetDescriptor instead')
const KVGet_Output$json = {
  '1': 'Output',
  '2': [
    {
      '1': 'KV',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBKV',
      '10': 'KV'
    },
  ],
};

/// Descriptor for `KVGet`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVGetDescriptor = $convert.base64Decode(
    'CgVLVkdldBoZCgVJbnB1dBIQCgNLZXkYASABKAlSA0tleRouCgZPdXRwdXQSJAoCS1YYASABKA'
    'syFC5ra19ldGNkX21vZGVscy5QQktWUgJLVg==');
