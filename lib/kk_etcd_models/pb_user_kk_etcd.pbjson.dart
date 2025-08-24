// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_user_kk_etcd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBUserDescriptor instead')
const PBUser$json = {
  '1': 'PBUser',
  '2': [
    {'1': 'UserName', '3': 2, '4': 1, '5': 9, '10': 'UserName'},
    {'1': 'Password', '3': 3, '4': 1, '5': 9, '10': 'Password'},
    {'1': 'Roles', '3': 4, '4': 3, '5': 9, '10': 'Roles'},
  ],
};

/// Descriptor for `PBUser`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBUserDescriptor = $convert.base64Decode(
    'CgZQQlVzZXISGgoIVXNlck5hbWUYAiABKAlSCFVzZXJOYW1lEhoKCFBhc3N3b3JkGAMgASgJUg'
    'hQYXNzd29yZBIUCgVSb2xlcxgEIAMoCVIFUm9sZXM=');

@$core.Deprecated('Use pBListUserDescriptor instead')
const PBListUser$json = {
  '1': 'PBListUser',
  '2': [
    {
      '1': 'ListUser',
      '3': 1,
      '4': 3,
      '5': 11,
      '6': '.kk_etcd.PBUser',
      '10': 'ListUser'
    },
  ],
};

/// Descriptor for `PBListUser`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListUserDescriptor = $convert.base64Decode(
    'CgpQQkxpc3RVc2VyEisKCExpc3RVc2VyGAEgAygLMg8ua2tfZXRjZC5QQlVzZXJSCExpc3RVc2'
    'Vy');
