//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_kv/kVGet/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

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
