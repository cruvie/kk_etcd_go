//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_user/myInfo/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use myInfoDescriptor instead')
const MyInfo$json = {
  '1': 'MyInfo',
  '3': [MyInfo_Input$json, MyInfo_Output$json],
};

@$core.Deprecated('Use myInfoDescriptor instead')
const MyInfo_Input$json = {
  '1': 'Input',
};

@$core.Deprecated('Use myInfoDescriptor instead')
const MyInfo_Output$json = {
  '1': 'Output',
  '2': [
    {'1': 'UserName', '3': 2, '4': 1, '5': 9, '10': 'UserName'},
    {'1': 'Roles', '3': 4, '4': 3, '5': 9, '10': 'Roles'},
  ],
};

/// Descriptor for `MyInfo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List myInfoDescriptor = $convert.base64Decode(
    'CgZNeUluZm8aBwoFSW5wdXQaOgoGT3V0cHV0EhoKCFVzZXJOYW1lGAIgASgJUghVc2VyTmFtZR'
    'IUCgVSb2xlcxgEIAMoCVIFUm9sZXM=');
