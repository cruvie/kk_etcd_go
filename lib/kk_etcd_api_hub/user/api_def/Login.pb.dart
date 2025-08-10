// This is a generated file - do not edit.
//
// Generated from kk_etcd_api_hub/user/api_def/Login.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class Login_Input extends $pb.GeneratedMessage {
  factory Login_Input({
    $core.String? userName,
    $core.String? password,
  }) {
    final result = create();
    if (userName != null) result.userName = userName;
    if (password != null) result.password = password;
    return result;
  }

  Login_Input._();

  factory Login_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory Login_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Login.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..aOS(2, _omitFieldNames ? '' : 'Password', protoName: 'Password')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Login_Input clone() => Login_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Login_Input copyWith(void Function(Login_Input) updates) =>
      super.copyWith((message) => updates(message as Login_Input))
          as Login_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Login_Input create() => Login_Input._();
  @$core.override
  Login_Input createEmptyInstance() => create();
  static $pb.PbList<Login_Input> createRepeated() => $pb.PbList<Login_Input>();
  @$core.pragma('dart2js:noInline')
  static Login_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<Login_Input>(create);
  static Login_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userName => $_getSZ(0);
  @$pb.TagNumber(1)
  set userName($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasUserName() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserName() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get password => $_getSZ(1);
  @$pb.TagNumber(2)
  set password($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasPassword() => $_has(1);
  @$pb.TagNumber(2)
  void clearPassword() => $_clearField(2);
}

class Login_Output extends $pb.GeneratedMessage {
  factory Login_Output({
    $core.String? token,
  }) {
    final result = create();
    if (token != null) result.token = token;
    return result;
  }

  Login_Output._();

  factory Login_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory Login_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Login.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Token', protoName: 'Token')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Login_Output clone() => Login_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Login_Output copyWith(void Function(Login_Output) updates) =>
      super.copyWith((message) => updates(message as Login_Output))
          as Login_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Login_Output create() => Login_Output._();
  @$core.override
  Login_Output createEmptyInstance() => create();
  static $pb.PbList<Login_Output> createRepeated() =>
      $pb.PbList<Login_Output>();
  @$core.pragma('dart2js:noInline')
  static Login_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<Login_Output>(create);
  static Login_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => $_clearField(1);
}

/// login
class Login extends $pb.GeneratedMessage {
  factory Login() => create();

  Login._();

  factory Login.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory Login.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Login',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Login clone() => Login()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Login copyWith(void Function(Login) updates) =>
      super.copyWith((message) => updates(message as Login)) as Login;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Login create() => Login._();
  @$core.override
  Login createEmptyInstance() => create();
  static $pb.PbList<Login> createRepeated() => $pb.PbList<Login>();
  @$core.pragma('dart2js:noInline')
  static Login getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Login>(create);
  static Login? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
