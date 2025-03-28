//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/ai/query/api.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use queryDescriptor instead')
const Query$json = {
  '1': 'Query',
  '3': [Query_Input$json, Query_Output$json],
};

@$core.Deprecated('Use queryDescriptor instead')
const Query_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Question', '3': 1, '4': 1, '5': 9, '10': 'Question'},
  ],
};

@$core.Deprecated('Use queryDescriptor instead')
const Query_Output$json = {
  '1': 'Output',
  '2': [
    {'1': 'Answer', '3': 1, '4': 1, '5': 9, '10': 'Answer'},
  ],
};

/// Descriptor for `Query`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List queryDescriptor = $convert.base64Decode(
    'CgVRdWVyeRojCgVJbnB1dBIaCghRdWVzdGlvbhgBIAEoCVIIUXVlc3Rpb24aIAoGT3V0cHV0Eh'
    'YKBkFuc3dlchgBIAEoCVIGQW5zd2Vy');

