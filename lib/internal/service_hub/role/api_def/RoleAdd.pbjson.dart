// This is a generated file - do not edit.
//
// Generated from internal/service_hub/role/api_def/RoleAdd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roleAddDescriptor instead')
const RoleAdd$json = {
  '1': 'RoleAdd',
  '3': [RoleAdd_Input$json, RoleAdd_Output$json],
};

@$core.Deprecated('Use roleAddDescriptor instead')
const RoleAdd_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
  ],
};

@$core.Deprecated('Use roleAddDescriptor instead')
const RoleAdd_Output$json = {
  '1': 'Output',
};

/// Descriptor for `RoleAdd`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleAddDescriptor = $convert.base64Decode(
    'CgdSb2xlQWRkGhsKBUlucHV0EhIKBE5hbWUYASABKAlSBE5hbWUaCAoGT3V0cHV0');
