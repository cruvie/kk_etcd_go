//
//  Generated code. Do not modify.
//  source: api_user.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use loginParamDescriptor instead')
const LoginParam$json = {
  '1': 'LoginParam',
  '2': [
    {'1': 'UserName', '3': 1, '4': 1, '5': 9, '10': 'UserName'},
    {'1': 'Password', '3': 2, '4': 1, '5': 9, '10': 'Password'},
  ],
};

/// Descriptor for `LoginParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginParamDescriptor = $convert.base64Decode(
    'CgpMb2dpblBhcmFtEhoKCFVzZXJOYW1lGAEgASgJUghVc2VyTmFtZRIaCghQYXNzd29yZBgCIA'
    'EoCVIIUGFzc3dvcmQ=');

@$core.Deprecated('Use loginResponseDescriptor instead')
const LoginResponse$json = {
  '1': 'LoginResponse',
  '2': [
    {'1': 'Token', '3': 1, '4': 1, '5': 9, '10': 'Token'},
  ],
};

/// Descriptor for `LoginResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginResponseDescriptor = $convert.base64Decode(
    'Cg1Mb2dpblJlc3BvbnNlEhQKBVRva2VuGAEgASgJUgVUb2tlbg==');

@$core.Deprecated('Use logoutParamDescriptor instead')
const LogoutParam$json = {
  '1': 'LogoutParam',
};

/// Descriptor for `LogoutParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List logoutParamDescriptor = $convert.base64Decode(
    'CgtMb2dvdXRQYXJhbQ==');

@$core.Deprecated('Use logoutResponseDescriptor instead')
const LogoutResponse$json = {
  '1': 'LogoutResponse',
};

/// Descriptor for `LogoutResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List logoutResponseDescriptor = $convert.base64Decode(
    'Cg5Mb2dvdXRSZXNwb25zZQ==');

@$core.Deprecated('Use userAddParamDescriptor instead')
const UserAddParam$json = {
  '1': 'UserAddParam',
  '2': [
    {'1': 'UserName', '3': 2, '4': 1, '5': 9, '10': 'UserName'},
    {'1': 'Password', '3': 3, '4': 1, '5': 9, '10': 'Password'},
    {'1': 'Roles', '3': 4, '4': 3, '5': 9, '10': 'Roles'},
  ],
};

/// Descriptor for `UserAddParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userAddParamDescriptor = $convert.base64Decode(
    'CgxVc2VyQWRkUGFyYW0SGgoIVXNlck5hbWUYAiABKAlSCFVzZXJOYW1lEhoKCFBhc3N3b3JkGA'
    'MgASgJUghQYXNzd29yZBIUCgVSb2xlcxgEIAMoCVIFUm9sZXM=');

@$core.Deprecated('Use userAddResponseDescriptor instead')
const UserAddResponse$json = {
  '1': 'UserAddResponse',
};

/// Descriptor for `UserAddResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userAddResponseDescriptor = $convert.base64Decode(
    'Cg9Vc2VyQWRkUmVzcG9uc2U=');

@$core.Deprecated('Use userDeleteParamDescriptor instead')
const UserDeleteParam$json = {
  '1': 'UserDeleteParam',
  '2': [
    {'1': 'UserName', '3': 1, '4': 1, '5': 9, '10': 'UserName'},
  ],
};

/// Descriptor for `UserDeleteParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userDeleteParamDescriptor = $convert.base64Decode(
    'Cg9Vc2VyRGVsZXRlUGFyYW0SGgoIVXNlck5hbWUYASABKAlSCFVzZXJOYW1l');

@$core.Deprecated('Use userDeleteResponseDescriptor instead')
const UserDeleteResponse$json = {
  '1': 'UserDeleteResponse',
};

/// Descriptor for `UserDeleteResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userDeleteResponseDescriptor = $convert.base64Decode(
    'ChJVc2VyRGVsZXRlUmVzcG9uc2U=');

@$core.Deprecated('Use getUserParamDescriptor instead')
const GetUserParam$json = {
  '1': 'GetUserParam',
  '2': [
    {'1': 'UserName', '3': 1, '4': 1, '5': 9, '10': 'UserName'},
  ],
};

/// Descriptor for `GetUserParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getUserParamDescriptor = $convert.base64Decode(
    'CgxHZXRVc2VyUGFyYW0SGgoIVXNlck5hbWUYASABKAlSCFVzZXJOYW1l');

@$core.Deprecated('Use getUserResponseDescriptor instead')
const GetUserResponse$json = {
  '1': 'GetUserResponse',
};

/// Descriptor for `GetUserResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getUserResponseDescriptor = $convert.base64Decode(
    'Cg9HZXRVc2VyUmVzcG9uc2U=');

@$core.Deprecated('Use myInfoParamDescriptor instead')
const MyInfoParam$json = {
  '1': 'MyInfoParam',
};

/// Descriptor for `MyInfoParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List myInfoParamDescriptor = $convert.base64Decode(
    'CgtNeUluZm9QYXJhbQ==');

@$core.Deprecated('Use myInfoResponseDescriptor instead')
const MyInfoResponse$json = {
  '1': 'MyInfoResponse',
  '2': [
    {'1': 'UserName', '3': 2, '4': 1, '5': 9, '10': 'UserName'},
    {'1': 'Roles', '3': 4, '4': 3, '5': 9, '10': 'Roles'},
  ],
};

/// Descriptor for `MyInfoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List myInfoResponseDescriptor = $convert.base64Decode(
    'Cg5NeUluZm9SZXNwb25zZRIaCghVc2VyTmFtZRgCIAEoCVIIVXNlck5hbWUSFAoFUm9sZXMYBC'
    'ADKAlSBVJvbGVz');

@$core.Deprecated('Use userListParamDescriptor instead')
const UserListParam$json = {
  '1': 'UserListParam',
};

/// Descriptor for `UserListParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userListParamDescriptor = $convert.base64Decode(
    'Cg1Vc2VyTGlzdFBhcmFt');

@$core.Deprecated('Use userListResponseDescriptor instead')
const UserListResponse$json = {
  '1': 'UserListResponse',
  '2': [
    {'1': 'ListUser', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBListUser', '10': 'ListUser'},
  ],
};

/// Descriptor for `UserListResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userListResponseDescriptor = $convert.base64Decode(
    'ChBVc2VyTGlzdFJlc3BvbnNlEjYKCExpc3RVc2VyGAEgASgLMhoua2tfZXRjZF9tb2RlbHMuUE'
    'JMaXN0VXNlclIITGlzdFVzZXI=');

@$core.Deprecated('Use userGrantRoleParamDescriptor instead')
const UserGrantRoleParam$json = {
  '1': 'UserGrantRoleParam',
  '2': [
    {'1': 'UserName', '3': 2, '4': 1, '5': 9, '10': 'UserName'},
    {'1': 'Roles', '3': 4, '4': 3, '5': 9, '10': 'Roles'},
  ],
};

/// Descriptor for `UserGrantRoleParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userGrantRoleParamDescriptor = $convert.base64Decode(
    'ChJVc2VyR3JhbnRSb2xlUGFyYW0SGgoIVXNlck5hbWUYAiABKAlSCFVzZXJOYW1lEhQKBVJvbG'
    'VzGAQgAygJUgVSb2xlcw==');

@$core.Deprecated('Use userGrantRoleResponseDescriptor instead')
const UserGrantRoleResponse$json = {
  '1': 'UserGrantRoleResponse',
};

/// Descriptor for `UserGrantRoleResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userGrantRoleResponseDescriptor = $convert.base64Decode(
    'ChVVc2VyR3JhbnRSb2xlUmVzcG9uc2U=');

