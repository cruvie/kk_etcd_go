// This is a generated file - do not edit.
//
// Generated from internal/service_hub/service/api_def/ServiceList.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use serviceListDescriptor instead')
const ServiceList$json = {
  '1': 'ServiceList',
  '3': [ServiceList_Input$json, ServiceList_Output$json],
};

@$core.Deprecated('Use serviceListDescriptor instead')
const ServiceList_Input$json = {
  '1': 'Input',
  '2': [
    {
      '1': 'ServiceType',
      '3': 1,
      '4': 1,
      '5': 14,
      '6': '.kk_etcd.PBServiceType',
      '10': 'ServiceType'
    },
    {'1': 'ServiceName', '3': 2, '4': 1, '5': 9, '10': 'ServiceName'},
  ],
};

@$core.Deprecated('Use serviceListDescriptor instead')
const ServiceList_Output$json = {
  '1': 'Output',
  '2': [
    {
      '1': 'ServiceList',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd.PBListService',
      '10': 'ServiceList'
    },
  ],
};

/// Descriptor for `ServiceList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List serviceListDescriptor = $convert.base64Decode(
    'CgtTZXJ2aWNlTGlzdBpjCgVJbnB1dBI4CgtTZXJ2aWNlVHlwZRgBIAEoDjIWLmtrX2V0Y2QuUE'
    'JTZXJ2aWNlVHlwZVILU2VydmljZVR5cGUSIAoLU2VydmljZU5hbWUYAiABKAlSC1NlcnZpY2VO'
    'YW1lGkIKBk91dHB1dBI4CgtTZXJ2aWNlTGlzdBgBIAEoCzIWLmtrX2V0Y2QuUEJMaXN0U2Vydm'
    'ljZVILU2VydmljZUxpc3Q=');
