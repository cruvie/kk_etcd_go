//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/user/userAdd/api.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use userAddDescriptor instead')
const UserAdd$json = {
  '1': 'UserAdd',
  '3': [UserAdd_Input$json, UserAdd_Output$json],
};

@$core.Deprecated('Use userAddDescriptor instead')
const UserAdd_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'UserName', '3': 2, '4': 1, '5': 9, '10': 'UserName'},
    {'1': 'Password', '3': 3, '4': 1, '5': 9, '10': 'Password'},
    {'1': 'Roles', '3': 4, '4': 3, '5': 9, '10': 'Roles'},
  ],
};

@$core.Deprecated('Use userAddDescriptor instead')
const UserAdd_Output$json = {
  '1': 'Output',
};

/// Descriptor for `UserAdd`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userAddDescriptor = $convert.base64Decode(
    'CgdVc2VyQWRkGlUKBUlucHV0EhoKCFVzZXJOYW1lGAIgASgJUghVc2VyTmFtZRIaCghQYXNzd2'
    '9yZBgDIAEoCVIIUGFzc3dvcmQSFAoFUm9sZXMYBCADKAlSBVJvbGVzGggKBk91dHB1dA==');

