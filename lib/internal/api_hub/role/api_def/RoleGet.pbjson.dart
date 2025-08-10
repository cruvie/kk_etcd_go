// This is a generated file - do not edit.
//
// Generated from internal/api_hub/role/api_def/RoleGet.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roleGetDescriptor instead')
const RoleGet$json = {
  '1': 'RoleGet',
  '3': [RoleGet_Input$json, RoleGet_Output$json],
};

@$core.Deprecated('Use roleGetDescriptor instead')
const RoleGet_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
  ],
};

@$core.Deprecated('Use roleGetDescriptor instead')
const RoleGet_Output$json = {
  '1': 'Output',
  '2': [
    {
      '1': 'Role',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBRole',
      '10': 'Role'
    },
  ],
};

/// Descriptor for `RoleGet`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleGetDescriptor = $convert.base64Decode(
    'CgdSb2xlR2V0GhsKBUlucHV0EhIKBE5hbWUYASABKAlSBE5hbWUaNAoGT3V0cHV0EioKBFJvbG'
    'UYASABKAsyFi5ra19ldGNkX21vZGVscy5QQlJvbGVSBFJvbGU=');
