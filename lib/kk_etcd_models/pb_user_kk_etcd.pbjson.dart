//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_user_kk_etcd.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

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
    {'1': 'ListUser', '3': 1, '4': 3, '5': 11, '6': '.kk_etcd_models.PBUser', '10': 'ListUser'},
  ],
};

/// Descriptor for `PBListUser`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListUserDescriptor = $convert.base64Decode(
    'CgpQQkxpc3RVc2VyEjIKCExpc3RVc2VyGAEgAygLMhYua2tfZXRjZF9tb2RlbHMuUEJVc2VyUg'
    'hMaXN0VXNlcg==');

