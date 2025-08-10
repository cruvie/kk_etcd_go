// This is a generated file - do not edit.
//
// Generated from internal/api_hub/kv/api_def/KVDel.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use kVDelDescriptor instead')
const KVDel$json = {
  '1': 'KVDel',
  '3': [KVDel_Input$json, KVDel_Output$json],
};

@$core.Deprecated('Use kVDelDescriptor instead')
const KVDel_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Key', '3': 1, '4': 1, '5': 9, '10': 'Key'},
  ],
};

@$core.Deprecated('Use kVDelDescriptor instead')
const KVDel_Output$json = {
  '1': 'Output',
};

/// Descriptor for `KVDel`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVDelDescriptor = $convert.base64Decode(
    'CgVLVkRlbBoZCgVJbnB1dBIQCgNLZXkYASABKAlSA0tleRoICgZPdXRwdXQ=');
