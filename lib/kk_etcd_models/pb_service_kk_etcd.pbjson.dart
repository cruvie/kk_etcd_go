// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_service_kk_etcd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBServiceDescriptor instead')
const PBService$json = {
  '1': 'PBService',
  '2': [
    {
      '1': 'ServiceRegistration',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBServiceRegistration',
      '10': 'ServiceRegistration'
    },
    {
      '1': 'Status',
      '3': 5,
      '4': 1,
      '5': 14,
      '6': '.kk_etcd_models.PBService.ServiceStatus',
      '10': 'Status'
    },
  ],
  '4': [PBService_ServiceStatus$json],
};

@$core.Deprecated('Use pBServiceDescriptor instead')
const PBService_ServiceStatus$json = {
  '1': 'ServiceStatus',
  '2': [
    {'1': 'UnKnown', '2': 0},
    {'1': 'Running', '2': 1},
    {'1': 'Stop', '2': 2},
  ],
};

/// Descriptor for `PBService`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBServiceDescriptor = $convert.base64Decode(
    'CglQQlNlcnZpY2USVwoTU2VydmljZVJlZ2lzdHJhdGlvbhgBIAEoCzIlLmtrX2V0Y2RfbW9kZW'
    'xzLlBCU2VydmljZVJlZ2lzdHJhdGlvblITU2VydmljZVJlZ2lzdHJhdGlvbhI/CgZTdGF0dXMY'
    'BSABKA4yJy5ra19ldGNkX21vZGVscy5QQlNlcnZpY2UuU2VydmljZVN0YXR1c1IGU3RhdHVzIj'
    'MKDVNlcnZpY2VTdGF0dXMSCwoHVW5Lbm93bhAAEgsKB1J1bm5pbmcQARIICgRTdG9wEAI=');

@$core.Deprecated('Use pBListServiceDescriptor instead')
const PBListService$json = {
  '1': 'PBListService',
  '2': [
    {
      '1': 'ListService',
      '3': 1,
      '4': 3,
      '5': 11,
      '6': '.kk_etcd_models.PBService',
      '10': 'ListService'
    },
  ],
};

/// Descriptor for `PBListService`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListServiceDescriptor = $convert.base64Decode(
    'Cg1QQkxpc3RTZXJ2aWNlEjsKC0xpc3RTZXJ2aWNlGAEgAygLMhkua2tfZXRjZF9tb2RlbHMuUE'
    'JTZXJ2aWNlUgtMaXN0U2VydmljZQ==');
