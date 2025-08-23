// This is a generated file - do not edit.
//
// Generated from internal/service_hub/user/api_def/UserAdd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class UserAdd_Input extends $pb.GeneratedMessage {
  factory UserAdd_Input({
    $core.String? userName,
    $core.String? password,
    $core.Iterable<$core.String>? roles,
  }) {
    final result = create();
    if (userName != null) result.userName = userName;
    if (password != null) result.password = password;
    if (roles != null) result.roles.addAll(roles);
    return result;
  }

  UserAdd_Input._();

  factory UserAdd_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UserAdd_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserAdd.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(2, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..aOS(3, _omitFieldNames ? '' : 'Password', protoName: 'Password')
    ..pPS(4, _omitFieldNames ? '' : 'Roles', protoName: 'Roles')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserAdd_Input clone() => UserAdd_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserAdd_Input copyWith(void Function(UserAdd_Input) updates) =>
      super.copyWith((message) => updates(message as UserAdd_Input))
          as UserAdd_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserAdd_Input create() => UserAdd_Input._();
  @$core.override
  UserAdd_Input createEmptyInstance() => create();
  static $pb.PbList<UserAdd_Input> createRepeated() =>
      $pb.PbList<UserAdd_Input>();
  @$core.pragma('dart2js:noInline')
  static UserAdd_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserAdd_Input>(create);
  static UserAdd_Input? _defaultInstance;

  @$pb.TagNumber(2)
  $core.String get userName => $_getSZ(0);
  @$pb.TagNumber(2)
  set userName($core.String value) => $_setString(0, value);
  @$pb.TagNumber(2)
  $core.bool hasUserName() => $_has(0);
  @$pb.TagNumber(2)
  void clearUserName() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get password => $_getSZ(1);
  @$pb.TagNumber(3)
  set password($core.String value) => $_setString(1, value);
  @$pb.TagNumber(3)
  $core.bool hasPassword() => $_has(1);
  @$pb.TagNumber(3)
  void clearPassword() => $_clearField(3);

  @$pb.TagNumber(4)
  $pb.PbList<$core.String> get roles => $_getList(2);
}

class UserAdd_Output extends $pb.GeneratedMessage {
  factory UserAdd_Output() => create();

  UserAdd_Output._();

  factory UserAdd_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UserAdd_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserAdd.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserAdd_Output clone() => UserAdd_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserAdd_Output copyWith(void Function(UserAdd_Output) updates) =>
      super.copyWith((message) => updates(message as UserAdd_Output))
          as UserAdd_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserAdd_Output create() => UserAdd_Output._();
  @$core.override
  UserAdd_Output createEmptyInstance() => create();
  static $pb.PbList<UserAdd_Output> createRepeated() =>
      $pb.PbList<UserAdd_Output>();
  @$core.pragma('dart2js:noInline')
  static UserAdd_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserAdd_Output>(create);
  static UserAdd_Output? _defaultInstance;
}

/// add user
class UserAdd extends $pb.GeneratedMessage {
  factory UserAdd() => create();

  UserAdd._();

  factory UserAdd.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UserAdd.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserAdd',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserAdd clone() => UserAdd()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserAdd copyWith(void Function(UserAdd) updates) =>
      super.copyWith((message) => updates(message as UserAdd)) as UserAdd;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserAdd create() => UserAdd._();
  @$core.override
  UserAdd createEmptyInstance() => create();
  static $pb.PbList<UserAdd> createRepeated() => $pb.PbList<UserAdd>();
  @$core.pragma('dart2js:noInline')
  static UserAdd getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserAdd>(create);
  static UserAdd? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
