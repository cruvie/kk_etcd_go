// This is a generated file - do not edit.
//
// Generated from kk_etcd_api_hub/role/api_def/RoleRevokePermission.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roleRevokePermissionDescriptor instead')
const RoleRevokePermission$json = {
  '1': 'RoleRevokePermission',
  '3': [RoleRevokePermission_Input$json, RoleRevokePermission_Output$json],
};

@$core.Deprecated('Use roleRevokePermissionDescriptor instead')
const RoleRevokePermission_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    {'1': 'Key', '3': 2, '4': 1, '5': 9, '10': 'Key'},
    {'1': 'RangeEnd', '3': 3, '4': 1, '5': 9, '10': 'RangeEnd'},
  ],
};

@$core.Deprecated('Use roleRevokePermissionDescriptor instead')
const RoleRevokePermission_Output$json = {
  '1': 'Output',
};

/// Descriptor for `RoleRevokePermission`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleRevokePermissionDescriptor = $convert.base64Decode(
    'ChRSb2xlUmV2b2tlUGVybWlzc2lvbhpJCgVJbnB1dBISCgROYW1lGAEgASgJUgROYW1lEhAKA0'
    'tleRgCIAEoCVIDS2V5EhoKCFJhbmdlRW5kGAMgASgJUghSYW5nZUVuZBoICgZPdXRwdXQ=');
