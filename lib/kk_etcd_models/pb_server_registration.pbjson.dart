//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_server_registration.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBServerTypeDescriptor instead')
const PBServerType$json = {
  '1': 'PBServerType',
  '2': [
    {'1': 'Unknown', '2': 0},
    {'1': 'Http', '2': 1},
    {'1': 'Grpc', '2': 2},
  ],
};

/// Descriptor for `PBServerType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List pBServerTypeDescriptor = $convert.base64Decode(
    'CgxQQlNlcnZlclR5cGUSCwoHVW5rbm93bhAAEggKBEh0dHAQARIICgRHcnBjEAI=');

@$core.Deprecated('Use pBServerRegistrationDescriptor instead')
const PBServerRegistration$json = {
  '1': 'PBServerRegistration',
  '2': [
    {'1': 'ServerType', '3': 1, '4': 1, '5': 14, '6': '.kk_etcd_models.PBServerType', '10': 'ServerType'},
    {'1': 'ServerName', '3': 2, '4': 1, '5': 9, '10': 'ServerName'},
    {'1': 'ServerAddr', '3': 3, '4': 1, '5': 9, '10': 'ServerAddr'},
    {'1': 'CheckConfig', '3': 4, '4': 1, '5': 11, '6': '.kk_etcd_models.PBServerRegistration.PBCheckConfig', '10': 'CheckConfig'},
  ],
  '3': [PBServerRegistration_PBCheckConfig$json],
};

@$core.Deprecated('Use pBServerRegistrationDescriptor instead')
const PBServerRegistration_PBCheckConfig$json = {
  '1': 'PBCheckConfig',
  '2': [
    {'1': 'Type', '3': 1, '4': 1, '5': 14, '6': '.kk_etcd_models.PBServerType', '10': 'Type'},
    {'1': 'Timeout', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'Timeout'},
    {'1': 'Interval', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'Interval'},
    {'1': 'Addr', '3': 4, '4': 1, '5': 9, '10': 'Addr'},
  ],
};

/// Descriptor for `PBServerRegistration`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBServerRegistrationDescriptor = $convert.base64Decode(
    'ChRQQlNlcnZlclJlZ2lzdHJhdGlvbhI8CgpTZXJ2ZXJUeXBlGAEgASgOMhwua2tfZXRjZF9tb2'
    'RlbHMuUEJTZXJ2ZXJUeXBlUgpTZXJ2ZXJUeXBlEh4KClNlcnZlck5hbWUYAiABKAlSClNlcnZl'
    'ck5hbWUSHgoKU2VydmVyQWRkchgDIAEoCVIKU2VydmVyQWRkchJUCgtDaGVja0NvbmZpZxgEIA'
    'EoCzIyLmtrX2V0Y2RfbW9kZWxzLlBCU2VydmVyUmVnaXN0cmF0aW9uLlBCQ2hlY2tDb25maWdS'
    'C0NoZWNrQ29uZmlnGsEBCg1QQkNoZWNrQ29uZmlnEjAKBFR5cGUYASABKA4yHC5ra19ldGNkX2'
    '1vZGVscy5QQlNlcnZlclR5cGVSBFR5cGUSMwoHVGltZW91dBgCIAEoCzIZLmdvb2dsZS5wcm90'
    'b2J1Zi5EdXJhdGlvblIHVGltZW91dBI1CghJbnRlcnZhbBgDIAEoCzIZLmdvb2dsZS5wcm90b2'
    'J1Zi5EdXJhdGlvblIISW50ZXJ2YWwSEgoEQWRkchgEIAEoCVIEQWRkcg==');

@$core.Deprecated('Use pBListServerRegistrationDescriptor instead')
const PBListServerRegistration$json = {
  '1': 'PBListServerRegistration',
  '2': [
    {'1': 'List', '3': 1, '4': 3, '5': 11, '6': '.kk_etcd_models.PBServerRegistration', '10': 'List'},
  ],
};

/// Descriptor for `PBListServerRegistration`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListServerRegistrationDescriptor = $convert.base64Decode(
    'ChhQQkxpc3RTZXJ2ZXJSZWdpc3RyYXRpb24SOAoETGlzdBgBIAMoCzIkLmtrX2V0Y2RfbW9kZW'
    'xzLlBCU2VydmVyUmVnaXN0cmF0aW9uUgRMaXN0');

