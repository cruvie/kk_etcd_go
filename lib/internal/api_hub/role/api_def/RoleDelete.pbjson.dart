// This is a generated file - do not edit.
//
// Generated from internal/api_hub/role/api_def/RoleDelete.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roleDeleteDescriptor instead')
const RoleDelete$json = {
  '1': 'RoleDelete',
  '3': [RoleDelete_Input$json, RoleDelete_Output$json],
};

@$core.Deprecated('Use roleDeleteDescriptor instead')
const RoleDelete_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
  ],
};

@$core.Deprecated('Use roleDeleteDescriptor instead')
const RoleDelete_Output$json = {
  '1': 'Output',
};

/// Descriptor for `RoleDelete`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleDeleteDescriptor = $convert.base64Decode(
    'CgpSb2xlRGVsZXRlGhsKBUlucHV0EhIKBE5hbWUYASABKAlSBE5hbWUaCAoGT3V0cHV0');
