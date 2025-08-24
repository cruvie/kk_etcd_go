// This is a generated file - do not edit.
//
// Generated from internal/service_hub/kv/api_def/KVList.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

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
      '6': '.kk_etcd.PBListKV',
      '10': 'KVList'
    },
  ],
};

/// Descriptor for `KVList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVListDescriptor = $convert.base64Decode(
    'CgZLVkxpc3QaHwoFSW5wdXQSFgoGUHJlZml4GAEgASgJUgZQcmVmaXgaMwoGT3V0cHV0EikKBk'
    'tWTGlzdBgBIAEoCzIRLmtrX2V0Y2QuUEJMaXN0S1ZSBktWTGlzdA==');
