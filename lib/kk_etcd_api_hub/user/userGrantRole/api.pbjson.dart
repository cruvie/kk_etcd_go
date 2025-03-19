//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/user/userGrantRole/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use userGrantRoleDescriptor instead')
const UserGrantRole$json = {
  '1': 'UserGrantRole',
  '3': [UserGrantRole_Input$json, UserGrantRole_Output$json],
};

@$core.Deprecated('Use userGrantRoleDescriptor instead')
const UserGrantRole_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'UserName', '3': 2, '4': 1, '5': 9, '10': 'UserName'},
    {'1': 'Roles', '3': 4, '4': 3, '5': 9, '10': 'Roles'},
  ],
};

@$core.Deprecated('Use userGrantRoleDescriptor instead')
const UserGrantRole_Output$json = {
  '1': 'Output',
};

/// Descriptor for `UserGrantRole`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userGrantRoleDescriptor = $convert.base64Decode(
    'Cg1Vc2VyR3JhbnRSb2xlGjkKBUlucHV0EhoKCFVzZXJOYW1lGAIgASgJUghVc2VyTmFtZRIUCg'
    'VSb2xlcxgEIAMoCVIFUm9sZXMaCAoGT3V0cHV0');

