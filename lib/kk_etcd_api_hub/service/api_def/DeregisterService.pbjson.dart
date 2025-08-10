// This is a generated file - do not edit.
//
// Generated from kk_etcd_api_hub/service/api_def/DeregisterService.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use deregisterServiceDescriptor instead')
const DeregisterService$json = {
  '1': 'DeregisterService',
  '3': [DeregisterService_Input$json, DeregisterService_Output$json],
};

@$core.Deprecated('Use deregisterServiceDescriptor instead')
const DeregisterService_Input$json = {
  '1': 'Input',
  '2': [
    {
      '1': 'Service',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBService',
      '10': 'Service'
    },
  ],
};

@$core.Deprecated('Use deregisterServiceDescriptor instead')
const DeregisterService_Output$json = {
  '1': 'Output',
};

/// Descriptor for `DeregisterService`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deregisterServiceDescriptor = $convert.base64Decode(
    'ChFEZXJlZ2lzdGVyU2VydmljZRo8CgVJbnB1dBIzCgdTZXJ2aWNlGAEgASgLMhkua2tfZXRjZF'
    '9tb2RlbHMuUEJTZXJ2aWNlUgdTZXJ2aWNlGggKBk91dHB1dA==');
