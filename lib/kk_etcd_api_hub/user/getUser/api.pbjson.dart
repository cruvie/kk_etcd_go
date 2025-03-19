//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/user/getUser/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use getUserDescriptor instead')
const GetUser$json = {
  '1': 'GetUser',
  '3': [GetUser_Input$json, GetUser_Output$json],
};

@$core.Deprecated('Use getUserDescriptor instead')
const GetUser_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'UserName', '3': 1, '4': 1, '5': 9, '10': 'UserName'},
  ],
};

@$core.Deprecated('Use getUserDescriptor instead')
const GetUser_Output$json = {
  '1': 'Output',
  '2': [
    {
      '1': 'User',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBUser',
      '10': 'User'
    },
  ],
};

/// Descriptor for `GetUser`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getUserDescriptor = $convert.base64Decode(
    'CgdHZXRVc2VyGiMKBUlucHV0EhoKCFVzZXJOYW1lGAEgASgJUghVc2VyTmFtZRo0CgZPdXRwdX'
    'QSKgoEVXNlchgBIAEoCzIWLmtrX2V0Y2RfbW9kZWxzLlBCVXNlclIEVXNlcg==');

