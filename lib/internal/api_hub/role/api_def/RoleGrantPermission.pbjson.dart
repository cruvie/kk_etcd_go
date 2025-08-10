// This is a generated file - do not edit.
//
// Generated from internal/api_hub/role/api_def/RoleGrantPermission.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roleGrantPermissionDescriptor instead')
const RoleGrantPermission$json = {
  '1': 'RoleGrantPermission',
  '3': [RoleGrantPermission_Input$json, RoleGrantPermission_Output$json],
};

@$core.Deprecated('Use roleGrantPermissionDescriptor instead')
const RoleGrantPermission_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    {
      '1': 'Perm',
      '3': 2,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBPermission',
      '10': 'Perm'
    },
  ],
};

@$core.Deprecated('Use roleGrantPermissionDescriptor instead')
const RoleGrantPermission_Output$json = {
  '1': 'Output',
};

/// Descriptor for `RoleGrantPermission`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleGrantPermissionDescriptor = $convert.base64Decode(
    'ChNSb2xlR3JhbnRQZXJtaXNzaW9uGk0KBUlucHV0EhIKBE5hbWUYASABKAlSBE5hbWUSMAoEUG'
    'VybRgCIAEoCzIcLmtrX2V0Y2RfbW9kZWxzLlBCUGVybWlzc2lvblIEUGVybRoICgZPdXRwdXQ=');
