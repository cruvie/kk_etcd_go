// This is a generated file - do not edit.
//
// Generated from internal/api_hub/service/api_def/ServiceList.proto.

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
      '6': '.kk_etcd_models.PBServiceType',
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
      '6': '.kk_etcd_models.PBListService',
      '10': 'ServiceList'
    },
  ],
};

/// Descriptor for `ServiceList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List serviceListDescriptor = $convert.base64Decode(
    'CgtTZXJ2aWNlTGlzdBpqCgVJbnB1dBI/CgtTZXJ2aWNlVHlwZRgBIAEoDjIdLmtrX2V0Y2RfbW'
    '9kZWxzLlBCU2VydmljZVR5cGVSC1NlcnZpY2VUeXBlEiAKC1NlcnZpY2VOYW1lGAIgASgJUgtT'
    'ZXJ2aWNlTmFtZRpJCgZPdXRwdXQSPwoLU2VydmljZUxpc3QYASABKAsyHS5ra19ldGNkX21vZG'
    'Vscy5QQkxpc3RTZXJ2aWNlUgtTZXJ2aWNlTGlzdA==');
