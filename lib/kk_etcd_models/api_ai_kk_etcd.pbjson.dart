//
//  Generated code. Do not modify.
//  source: api_ai_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use queryParamDescriptor instead')
const QueryParam$json = {
  '1': 'QueryParam',
  '2': [
    {'1': 'Question', '3': 1, '4': 1, '5': 9, '10': 'Question'},
  ],
};

/// Descriptor for `QueryParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List queryParamDescriptor = $convert
    .base64Decode('CgpRdWVyeVBhcmFtEhoKCFF1ZXN0aW9uGAEgASgJUghRdWVzdGlvbg==');

@$core.Deprecated('Use queryResponseDescriptor instead')
const QueryResponse$json = {
  '1': 'QueryResponse',
  '2': [
    {'1': 'Answer', '3': 1, '4': 1, '5': 9, '10': 'Answer'},
  ],
};

/// Descriptor for `QueryResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List queryResponseDescriptor = $convert
    .base64Decode('Cg1RdWVyeVJlc3BvbnNlEhYKBkFuc3dlchgBIAEoCVIGQW5zd2Vy');
