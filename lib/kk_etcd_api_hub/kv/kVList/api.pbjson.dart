//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/kv/kVList/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use kVListDescriptor instead')
const KVList$json = {
  '1': 'KVList',
  '3': [KVList_Input$json, KVList_Output$json],
};

@$core.Deprecated('Use kVListDescriptor instead')
const KVList_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Prefix', '3': 1, '4': 1, '5': 9, '10': 'Prefix'},
  ],
};

@$core.Deprecated('Use kVListDescriptor instead')
const KVList_Output$json = {
  '1': 'Output',
  '2': [
    {
      '1': 'KVList',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBListKV',
      '10': 'KVList'
    },
  ],
};

/// Descriptor for `KVList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVListDescriptor = $convert.base64Decode(
    'CgZLVkxpc3QaHwoFSW5wdXQSFgoGUHJlZml4GAEgASgJUgZQcmVmaXgaOgoGT3V0cHV0EjAKBk'
    'tWTGlzdBgBIAEoCzIYLmtrX2V0Y2RfbW9kZWxzLlBCTGlzdEtWUgZLVkxpc3Q=');

