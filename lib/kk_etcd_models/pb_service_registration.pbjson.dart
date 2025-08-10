// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_service_registration.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBServiceTypeDescriptor instead')
const PBServiceType$json = {
  '1': 'PBServiceType',
  '2': [
    {'1': 'Unknown', '2': 0},
    {'1': 'Http', '2': 1},
    {'1': 'Grpc', '2': 2},
  ],
};

/// Descriptor for `PBServiceType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List pBServiceTypeDescriptor = $convert.base64Decode(
    'Cg1QQlNlcnZpY2VUeXBlEgsKB1Vua25vd24QABIICgRIdHRwEAESCAoER3JwYxAC');

@$core.Deprecated('Use pBServiceRegistrationDescriptor instead')
const PBServiceRegistration$json = {
  '1': 'PBServiceRegistration',
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
    {'1': 'ServiceAddr', '3': 3, '4': 1, '5': 9, '10': 'ServiceAddr'},
    {
      '1': 'CheckConfig',
      '3': 4,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBServiceRegistration.PBCheckConfig',
      '10': 'CheckConfig'
    },
  ],
  '3': [PBServiceRegistration_PBCheckConfig$json],
};

@$core.Deprecated('Use pBServiceRegistrationDescriptor instead')
const PBServiceRegistration_PBCheckConfig$json = {
  '1': 'PBCheckConfig',
  '2': [
    {
      '1': 'Type',
      '3': 1,
      '4': 1,
      '5': 14,
      '6': '.kk_etcd_models.PBServiceType',
      '10': 'Type'
    },
    {
      '1': 'Timeout',
      '3': 2,
      '4': 1,
      '5': 11,
      '6': '.google.protobuf.Duration',
      '10': 'Timeout'
    },
    {
      '1': 'Interval',
      '3': 3,
      '4': 1,
      '5': 11,
      '6': '.google.protobuf.Duration',
      '10': 'Interval'
    },
    {'1': 'Addr', '3': 4, '4': 1, '5': 9, '10': 'Addr'},
  ],
};

/// Descriptor for `PBServiceRegistration`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBServiceRegistrationDescriptor = $convert.base64Decode(
    'ChVQQlNlcnZpY2VSZWdpc3RyYXRpb24SPwoLU2VydmljZVR5cGUYASABKA4yHS5ra19ldGNkX2'
    '1vZGVscy5QQlNlcnZpY2VUeXBlUgtTZXJ2aWNlVHlwZRIgCgtTZXJ2aWNlTmFtZRgCIAEoCVIL'
    'U2VydmljZU5hbWUSIAoLU2VydmljZUFkZHIYAyABKAlSC1NlcnZpY2VBZGRyElUKC0NoZWNrQ2'
    '9uZmlnGAQgASgLMjMua2tfZXRjZF9tb2RlbHMuUEJTZXJ2aWNlUmVnaXN0cmF0aW9uLlBCQ2hl'
    'Y2tDb25maWdSC0NoZWNrQ29uZmlnGsIBCg1QQkNoZWNrQ29uZmlnEjEKBFR5cGUYASABKA4yHS'
    '5ra19ldGNkX21vZGVscy5QQlNlcnZpY2VUeXBlUgRUeXBlEjMKB1RpbWVvdXQYAiABKAsyGS5n'
    'b29nbGUucHJvdG9idWYuRHVyYXRpb25SB1RpbWVvdXQSNQoISW50ZXJ2YWwYAyABKAsyGS5nb2'
    '9nbGUucHJvdG9idWYuRHVyYXRpb25SCEludGVydmFsEhIKBEFkZHIYBCABKAlSBEFkZHI=');

@$core.Deprecated('Use pBListServiceRegistrationDescriptor instead')
const PBListServiceRegistration$json = {
  '1': 'PBListServiceRegistration',
  '2': [
    {
      '1': 'List',
      '3': 1,
      '4': 3,
      '5': 11,
      '6': '.kk_etcd_models.PBServiceRegistration',
      '10': 'List'
    },
  ],
};

/// Descriptor for `PBListServiceRegistration`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListServiceRegistrationDescriptor =
    $convert.base64Decode(
        'ChlQQkxpc3RTZXJ2aWNlUmVnaXN0cmF0aW9uEjkKBExpc3QYASADKAsyJS5ra19ldGNkX21vZG'
        'Vscy5QQlNlcnZpY2VSZWdpc3RyYXRpb25SBExpc3Q=');
