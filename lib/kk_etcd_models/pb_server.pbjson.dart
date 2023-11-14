///
//  Generated code. Do not modify.
//  source: pb_server.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use pBServerDescriptor instead')
const PBServer$json = const {
  '1': 'PBServer',
  '2': const [
    const {'1': 'ServiceName', '3': 1, '4': 1, '5': 9, '10': 'ServiceName'},
    const {'1': 'ServiceAddr', '3': 2, '4': 1, '5': 9, '10': 'ServiceAddr'},
  ],
};

/// Descriptor for `PBServer`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBServerDescriptor = $convert.base64Decode(
    'CghQQlNlcnZlchIgCgtTZXJ2aWNlTmFtZRgBIAEoCVILU2VydmljZU5hbWUSIAoLU2VydmljZUFkZHIYAiABKAlSC1NlcnZpY2VBZGRy');
@$core.Deprecated('Use pBListServerDescriptor instead')
const PBListServer$json = const {
  '1': 'PBListServer',
  '2': const [
    const {
      '1': 'ListServer',
      '3': 1,
      '4': 3,
      '5': 11,
      '6': '.kk_etcd_models.PBServer',
      '10': 'ListServer'
    },
  ],
};

/// Descriptor for `PBListServer`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListServerDescriptor = $convert.base64Decode(
    'CgxQQkxpc3RTZXJ2ZXISOAoKTGlzdFNlcnZlchgBIAMoCzIYLmtrX2V0Y2RfbW9kZWxzLlBCU2VydmVyUgpMaXN0U2VydmVy');
