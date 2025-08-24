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
      '6': '.kk_etcd.PBServiceRegistration',
      '10': 'ServiceRegistration'
    },
    {
      '1': 'Status',
      '3': 5,
      '4': 1,
      '5': 14,
      '6': '.kk_etcd.PBService.ServiceStatus',
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
    'CglQQlNlcnZpY2USUAoTU2VydmljZVJlZ2lzdHJhdGlvbhgBIAEoCzIeLmtrX2V0Y2QuUEJTZX'
    'J2aWNlUmVnaXN0cmF0aW9uUhNTZXJ2aWNlUmVnaXN0cmF0aW9uEjgKBlN0YXR1cxgFIAEoDjIg'
    'LmtrX2V0Y2QuUEJTZXJ2aWNlLlNlcnZpY2VTdGF0dXNSBlN0YXR1cyIzCg1TZXJ2aWNlU3RhdH'
    'VzEgsKB1VuS25vd24QABILCgdSdW5uaW5nEAESCAoEU3RvcBAC');

@$core.Deprecated('Use pBListServiceDescriptor instead')
const PBListService$json = {
  '1': 'PBListService',
  '2': [
    {
      '1': 'ListService',
      '3': 1,
      '4': 3,
      '5': 11,
      '6': '.kk_etcd.PBService',
      '10': 'ListService'
    },
  ],
};

/// Descriptor for `PBListService`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListServiceDescriptor = $convert.base64Decode(
    'Cg1QQkxpc3RTZXJ2aWNlEjQKC0xpc3RTZXJ2aWNlGAEgAygLMhIua2tfZXRjZC5QQlNlcnZpY2'
    'VSC0xpc3RTZXJ2aWNl');
