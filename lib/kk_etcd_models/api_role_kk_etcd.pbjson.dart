//
//  Generated code. Do not modify.
//  source: api_role_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roleAddParamDescriptor instead')
const RoleAddParam$json = {
  '1': 'RoleAddParam',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
  ],
};

/// Descriptor for `RoleAddParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleAddParamDescriptor = $convert.base64Decode(
    'CgxSb2xlQWRkUGFyYW0SEgoETmFtZRgBIAEoCVIETmFtZQ==');

@$core.Deprecated('Use roleAddResponseDescriptor instead')
const RoleAddResponse$json = {
  '1': 'RoleAddResponse',
};

/// Descriptor for `RoleAddResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleAddResponseDescriptor = $convert.base64Decode(
    'Cg9Sb2xlQWRkUmVzcG9uc2U=');

@$core.Deprecated('Use roleDeleteParamDescriptor instead')
const RoleDeleteParam$json = {
  '1': 'RoleDeleteParam',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
  ],
};

/// Descriptor for `RoleDeleteParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleDeleteParamDescriptor = $convert.base64Decode(
    'Cg9Sb2xlRGVsZXRlUGFyYW0SEgoETmFtZRgBIAEoCVIETmFtZQ==');

@$core.Deprecated('Use roleDeleteResponseDescriptor instead')
const RoleDeleteResponse$json = {
  '1': 'RoleDeleteResponse',
};

/// Descriptor for `RoleDeleteResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleDeleteResponseDescriptor = $convert.base64Decode(
    'ChJSb2xlRGVsZXRlUmVzcG9uc2U=');

@$core.Deprecated('Use roleGetParamDescriptor instead')
const RoleGetParam$json = {
  '1': 'RoleGetParam',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
  ],
};

/// Descriptor for `RoleGetParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleGetParamDescriptor = $convert.base64Decode(
    'CgxSb2xlR2V0UGFyYW0SEgoETmFtZRgBIAEoCVIETmFtZQ==');

@$core.Deprecated('Use roleGetResponseDescriptor instead')
const RoleGetResponse$json = {
  '1': 'RoleGetResponse',
  '2': [
    {'1': 'Role', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBRole', '10': 'Role'},
  ],
};

/// Descriptor for `RoleGetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleGetResponseDescriptor = $convert.base64Decode(
    'Cg9Sb2xlR2V0UmVzcG9uc2USKgoEUm9sZRgBIAEoCzIWLmtrX2V0Y2RfbW9kZWxzLlBCUm9sZV'
    'IEUm9sZQ==');

@$core.Deprecated('Use roleListParamDescriptor instead')
const RoleListParam$json = {
  '1': 'RoleListParam',
};

/// Descriptor for `RoleListParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleListParamDescriptor = $convert.base64Decode(
    'Cg1Sb2xlTGlzdFBhcmFt');

@$core.Deprecated('Use roleListResponseDescriptor instead')
const RoleListResponse$json = {
  '1': 'RoleListResponse',
  '2': [
    {'1': 'ListRole', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBListRole', '10': 'ListRole'},
  ],
};

/// Descriptor for `RoleListResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleListResponseDescriptor = $convert.base64Decode(
    'ChBSb2xlTGlzdFJlc3BvbnNlEjYKCExpc3RSb2xlGAEgASgLMhoua2tfZXRjZF9tb2RlbHMuUE'
    'JMaXN0Um9sZVIITGlzdFJvbGU=');

@$core.Deprecated('Use roleGrantPermissionParamDescriptor instead')
const RoleGrantPermissionParam$json = {
  '1': 'RoleGrantPermissionParam',
  '2': [
    {'1': 'Role', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBRole', '10': 'Role'},
  ],
};

/// Descriptor for `RoleGrantPermissionParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleGrantPermissionParamDescriptor = $convert.base64Decode(
    'ChhSb2xlR3JhbnRQZXJtaXNzaW9uUGFyYW0SKgoEUm9sZRgBIAEoCzIWLmtrX2V0Y2RfbW9kZW'
    'xzLlBCUm9sZVIEUm9sZQ==');

@$core.Deprecated('Use roleGrantPermissionResponseDescriptor instead')
const RoleGrantPermissionResponse$json = {
  '1': 'RoleGrantPermissionResponse',
};

/// Descriptor for `RoleGrantPermissionResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleGrantPermissionResponseDescriptor = $convert.base64Decode(
    'ChtSb2xlR3JhbnRQZXJtaXNzaW9uUmVzcG9uc2U=');

