//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/user/api_def/UserList.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use userListDescriptor instead')
const UserList$json = {
  '1': 'UserList',
  '3': [UserList_Input$json, UserList_Output$json],
};

@$core.Deprecated('Use userListDescriptor instead')
const UserList_Input$json = {
  '1': 'Input',
};

@$core.Deprecated('Use userListDescriptor instead')
const UserList_Output$json = {
  '1': 'Output',
  '2': [
    {'1': 'ListUser', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBListUser', '10': 'ListUser'},
  ],
};

/// Descriptor for `UserList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userListDescriptor = $convert.base64Decode(
    'CghVc2VyTGlzdBoHCgVJbnB1dBpACgZPdXRwdXQSNgoITGlzdFVzZXIYASABKAsyGi5ra19ldG'
    'NkX21vZGVscy5QQkxpc3RVc2VyUghMaXN0VXNlcg==');

