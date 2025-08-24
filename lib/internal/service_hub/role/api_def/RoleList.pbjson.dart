// This is a generated file - do not edit.
//
// Generated from internal/service_hub/role/api_def/RoleList.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roleListDescriptor instead')
const RoleList$json = {
  '1': 'RoleList',
  '3': [RoleList_Input$json, RoleList_Output$json],
};

@$core.Deprecated('Use roleListDescriptor instead')
const RoleList_Input$json = {
  '1': 'Input',
};

@$core.Deprecated('Use roleListDescriptor instead')
const RoleList_Output$json = {
  '1': 'Output',
  '2': [
    {
      '1': 'ListRole',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd.PBListRole',
      '10': 'ListRole'
    },
  ],
};

/// Descriptor for `RoleList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleListDescriptor = $convert.base64Decode(
    'CghSb2xlTGlzdBoHCgVJbnB1dBo5CgZPdXRwdXQSLwoITGlzdFJvbGUYASABKAsyEy5ra19ldG'
    'NkLlBCTGlzdFJvbGVSCExpc3RSb2xl');
