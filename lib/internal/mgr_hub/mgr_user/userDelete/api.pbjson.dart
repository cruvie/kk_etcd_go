//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_user/userDelete/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use userDeleteDescriptor instead')
const UserDelete$json = {
  '1': 'UserDelete',
  '3': [UserDelete_Input$json, UserDelete_Output$json],
};

@$core.Deprecated('Use userDeleteDescriptor instead')
const UserDelete_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'UserName', '3': 1, '4': 1, '5': 9, '10': 'UserName'},
  ],
};

@$core.Deprecated('Use userDeleteDescriptor instead')
const UserDelete_Output$json = {
  '1': 'Output',
};

/// Descriptor for `UserDelete`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userDeleteDescriptor = $convert.base64Decode(
    'CgpVc2VyRGVsZXRlGiMKBUlucHV0EhoKCFVzZXJOYW1lGAEgASgJUghVc2VyTmFtZRoICgZPdX'
    'RwdXQ=');
