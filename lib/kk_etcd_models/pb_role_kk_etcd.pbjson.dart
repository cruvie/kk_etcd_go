//
//  Generated code. Do not modify.
//  source: pb_role_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBRoleDescriptor instead')
const PBRole$json = {
  '1': 'PBRole',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    {'1': 'Perms', '3': 2, '4': 3, '5': 11, '6': '.kk_etcd_models.Permission', '10': 'Perms'},
  ],
};

/// Descriptor for `PBRole`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBRoleDescriptor = $convert.base64Decode(
    'CgZQQlJvbGUSEgoETmFtZRgBIAEoCVIETmFtZRIwCgVQZXJtcxgCIAMoCzIaLmtrX2V0Y2RfbW'
    '9kZWxzLlBlcm1pc3Npb25SBVBlcm1z');

@$core.Deprecated('Use permissionDescriptor instead')
const Permission$json = {
  '1': 'Permission',
  '2': [
    {'1': 'Key', '3': 2, '4': 1, '5': 9, '10': 'Key'},
    {'1': 'RangeEnd', '3': 3, '4': 1, '5': 9, '10': 'RangeEnd'},
    {'1': 'PermissionType', '3': 4, '4': 1, '5': 5, '10': 'PermissionType'},
  ],
};

/// Descriptor for `Permission`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List permissionDescriptor = $convert.base64Decode(
    'CgpQZXJtaXNzaW9uEhAKA0tleRgCIAEoCVIDS2V5EhoKCFJhbmdlRW5kGAMgASgJUghSYW5nZU'
    'VuZBImCg5QZXJtaXNzaW9uVHlwZRgEIAEoBVIOUGVybWlzc2lvblR5cGU=');

@$core.Deprecated('Use pBListRoleDescriptor instead')
const PBListRole$json = {
  '1': 'PBListRole',
  '2': [
    {'1': 'List', '3': 1, '4': 3, '5': 11, '6': '.kk_etcd_models.PBRole', '10': 'List'},
  ],
};

/// Descriptor for `PBListRole`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListRoleDescriptor = $convert.base64Decode(
    'CgpQQkxpc3RSb2xlEioKBExpc3QYASADKAsyFi5ra19ldGNkX21vZGVscy5QQlJvbGVSBExpc3'
    'Q=');

