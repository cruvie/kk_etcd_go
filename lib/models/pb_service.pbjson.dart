///
//  Generated code. Do not modify.
//  source: pb_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBServiceDescriptor instead')
const PBService$json = const {
  '1': 'PBService',
  '2': const [
    const {'1': 'ServiceName', '3': 1, '4': 1, '5': 9, '10': 'ServiceName'},
    const {'1': 'ServiceAddr', '3': 2, '4': 1, '5': 9, '10': 'ServiceAddr'},
  ],
};

/// Descriptor for `PBService`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBServiceDescriptor = $convert.base64Decode(
    'CglQQlNlcnZpY2USIAoLU2VydmljZU5hbWUYASABKAlSC1NlcnZpY2VOYW1lEiAKC1NlcnZpY2VBZGRyGAIgASgJUgtTZXJ2aWNlQWRkcg==');
@$core.Deprecated('Use pBListServiceDescriptor instead')
const PBListService$json = const {
  '1': 'PBListService',
  '2': const [
    const {
      '1': 'ListService',
      '3': 1,
      '4': 3,
      '5': 11,
      '6': '.models.PBService',
      '10': 'ListService'
    },
  ],
};

/// Descriptor for `PBListService`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListServiceDescriptor = $convert.base64Decode(
    'Cg1QQkxpc3RTZXJ2aWNlEjMKC0xpc3RTZXJ2aWNlGAEgAygLMhEubW9kZWxzLlBCU2VydmljZVILTGlzdFNlcnZpY2U=');
