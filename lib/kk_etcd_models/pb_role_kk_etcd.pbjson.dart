//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_role_kk_etcd.proto
//
// @dart = 3.3

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
    {'1': 'Perms', '3': 2, '4': 3, '5': 11, '6': '.kk_etcd_models.PBPermission', '10': 'Perms'},
  ],
};

/// Descriptor for `PBRole`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBRoleDescriptor = $convert.base64Decode(
    'CgZQQlJvbGUSEgoETmFtZRgBIAEoCVIETmFtZRIyCgVQZXJtcxgCIAMoCzIcLmtrX2V0Y2RfbW'
    '9kZWxzLlBCUGVybWlzc2lvblIFUGVybXM=');

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
    {'1': 'List', '3': 1, '4': 3, '5': 11, '6': '.kk_etcd_models.PBRole', '10': 'List'},
  ],
};

/// Descriptor for `PBListRole`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListRoleDescriptor = $convert.base64Decode(
    'CgpQQkxpc3RSb2xlEioKBExpc3QYASADKAsyFi5ra19ldGNkX21vZGVscy5QQlJvbGVSBExpc3'
    'Q=');

