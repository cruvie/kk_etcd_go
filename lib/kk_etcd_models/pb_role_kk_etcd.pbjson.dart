// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_role_kk_etcd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBRoleDescriptor instead')
const PBRole$json = {
  '1': 'PBRole',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    {
      '1': 'Perms',
      '3': 2,
      '4': 3,
      '5': 11,
      '6': '.kk_etcd.PBPermission',
      '10': 'Perms'
    },
  ],
};

/// Descriptor for `PBRole`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBRoleDescriptor = $convert.base64Decode(
    'CgZQQlJvbGUSEgoETmFtZRgBIAEoCVIETmFtZRIrCgVQZXJtcxgCIAMoCzIVLmtrX2V0Y2QuUE'
    'JQZXJtaXNzaW9uUgVQZXJtcw==');

@$core.Deprecated('Use pBPermissionDescriptor instead')
const PBPermission$json = {
  '1': 'PBPermission',
  '2': [
    {'1': 'Key', '3': 2, '4': 1, '5': 9, '10': 'Key'},
    {'1': 'RangeEnd', '3': 3, '4': 1, '5': 9, '10': 'RangeEnd'},
    {'1': 'PermissionType', '3': 4, '4': 1, '5': 5, '10': 'PermissionType'},
  ],
};

/// Descriptor for `PBPermission`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBPermissionDescriptor = $convert.base64Decode(
    'CgxQQlBlcm1pc3Npb24SEAoDS2V5GAIgASgJUgNLZXkSGgoIUmFuZ2VFbmQYAyABKAlSCFJhbm'
    'dlRW5kEiYKDlBlcm1pc3Npb25UeXBlGAQgASgFUg5QZXJtaXNzaW9uVHlwZQ==');

@$core.Deprecated('Use pBListRoleDescriptor instead')
const PBListRole$json = {
  '1': 'PBListRole',
  '2': [
    {
      '1': 'List',
      '3': 1,
      '4': 3,
      '5': 11,
      '6': '.kk_etcd.PBRole',
      '10': 'List'
    },
  ],
};

/// Descriptor for `PBListRole`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListRoleDescriptor = $convert.base64Decode(
    'CgpQQkxpc3RSb2xlEiMKBExpc3QYASADKAsyDy5ra19ldGNkLlBCUm9sZVIETGlzdA==');
