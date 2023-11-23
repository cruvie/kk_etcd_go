//
//  Generated code. Do not modify.
//  source: pb_server.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBServerDescriptor instead')
const PBServer$json = {
  '1': 'PBServer',
  '2': [
    {'1': 'ServiceName', '3': 1, '4': 1, '5': 9, '10': 'ServiceName'},
    {'1': 'ServiceAddr', '3': 2, '4': 1, '5': 9, '10': 'ServiceAddr'},
  ],
};

/// Descriptor for `PBServer`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBServerDescriptor = $convert.base64Decode(
    'CghQQlNlcnZlchIgCgtTZXJ2aWNlTmFtZRgBIAEoCVILU2VydmljZU5hbWUSIAoLU2VydmljZU'
    'FkZHIYAiABKAlSC1NlcnZpY2VBZGRy');

@$core.Deprecated('Use pBListServerDescriptor instead')
const PBListServer$json = {
  '1': 'PBListServer',
  '2': [
    {'1': 'ListServer', '3': 1, '4': 3, '5': 11, '6': '.kk_etcd_models.PBServer', '10': 'ListServer'},
  ],
};

/// Descriptor for `PBListServer`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListServerDescriptor = $convert.base64Decode(
    'CgxQQkxpc3RTZXJ2ZXISOAoKTGlzdFNlcnZlchgBIAMoCzIYLmtrX2V0Y2RfbW9kZWxzLlBCU2'
    'VydmVyUgpMaXN0U2VydmVy');

