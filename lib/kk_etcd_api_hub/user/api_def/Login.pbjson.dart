//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/user/api_def/Login.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use loginDescriptor instead')
const Login$json = {
  '1': 'Login',
  '3': [Login_Input$json, Login_Output$json],
};

@$core.Deprecated('Use loginDescriptor instead')
const Login_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'UserName', '3': 1, '4': 1, '5': 9, '10': 'UserName'},
    {'1': 'Password', '3': 2, '4': 1, '5': 9, '10': 'Password'},
  ],
};

@$core.Deprecated('Use loginDescriptor instead')
const Login_Output$json = {
  '1': 'Output',
  '2': [
    {'1': 'Token', '3': 1, '4': 1, '5': 9, '10': 'Token'},
  ],
};

/// Descriptor for `Login`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginDescriptor = $convert.base64Decode(
    'CgVMb2dpbho/CgVJbnB1dBIaCghVc2VyTmFtZRgBIAEoCVIIVXNlck5hbWUSGgoIUGFzc3dvcm'
    'QYAiABKAlSCFBhc3N3b3JkGh4KBk91dHB1dBIUCgVUb2tlbhgBIAEoCVIFVG9rZW4=');

