// This is a generated file - do not edit.
//
// Generated from kk_etcd_api_hub/user/api_def/GetUser.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../kk_etcd_models/pb_user_kk_etcd.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class GetUser_Input extends $pb.GeneratedMessage {
  factory GetUser_Input({
    $core.String? userName,
  }) {
    final result = create();
    if (userName != null) result.userName = userName;
    return result;
  }

  GetUser_Input._();

  factory GetUser_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory GetUser_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'GetUser.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetUser_Input clone() => GetUser_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetUser_Input copyWith(void Function(GetUser_Input) updates) =>
      super.copyWith((message) => updates(message as GetUser_Input))
          as GetUser_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetUser_Input create() => GetUser_Input._();
  @$core.override
  GetUser_Input createEmptyInstance() => create();
  static $pb.PbList<GetUser_Input> createRepeated() =>
      $pb.PbList<GetUser_Input>();
  @$core.pragma('dart2js:noInline')
  static GetUser_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<GetUser_Input>(create);
  static GetUser_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userName => $_getSZ(0);
  @$pb.TagNumber(1)
  set userName($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasUserName() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserName() => $_clearField(1);
}

class GetUser_Output extends $pb.GeneratedMessage {
  factory GetUser_Output({
    $0.PBUser? user,
  }) {
    final result = create();
    if (user != null) result.user = user;
    return result;
  }

  GetUser_Output._();

  factory GetUser_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory GetUser_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'GetUser.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOM<$0.PBUser>(1, _omitFieldNames ? '' : 'User',
        protoName: 'User', subBuilder: $0.PBUser.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetUser_Output clone() => GetUser_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetUser_Output copyWith(void Function(GetUser_Output) updates) =>
      super.copyWith((message) => updates(message as GetUser_Output))
          as GetUser_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetUser_Output create() => GetUser_Output._();
  @$core.override
  GetUser_Output createEmptyInstance() => create();
  static $pb.PbList<GetUser_Output> createRepeated() =>
      $pb.PbList<GetUser_Output>();
  @$core.pragma('dart2js:noInline')
  static GetUser_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<GetUser_Output>(create);
  static GetUser_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBUser get user => $_getN(0);
  @$pb.TagNumber(1)
  set user($0.PBUser value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasUser() => $_has(0);
  @$pb.TagNumber(1)
  void clearUser() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBUser ensureUser() => $_ensure(0);
}

/// get user
class GetUser extends $pb.GeneratedMessage {
  factory GetUser() => create();

  GetUser._();

  factory GetUser.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory GetUser.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'GetUser',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetUser clone() => GetUser()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  GetUser copyWith(void Function(GetUser) updates) =>
      super.copyWith((message) => updates(message as GetUser)) as GetUser;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetUser create() => GetUser._();
  @$core.override
  GetUser createEmptyInstance() => create();
  static $pb.PbList<GetUser> createRepeated() => $pb.PbList<GetUser>();
  @$core.pragma('dart2js:noInline')
  static GetUser getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetUser>(create);
  static GetUser? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
