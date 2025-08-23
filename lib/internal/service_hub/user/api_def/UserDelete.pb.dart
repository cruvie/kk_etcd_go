// This is a generated file - do not edit.
//
// Generated from internal/service_hub/user/api_def/UserDelete.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class UserDelete_Input extends $pb.GeneratedMessage {
  factory UserDelete_Input({
    $core.String? userName,
  }) {
    final result = create();
    if (userName != null) result.userName = userName;
    return result;
  }

  UserDelete_Input._();

  factory UserDelete_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UserDelete_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserDelete.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete_Input clone() => UserDelete_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete_Input copyWith(void Function(UserDelete_Input) updates) =>
      super.copyWith((message) => updates(message as UserDelete_Input))
          as UserDelete_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserDelete_Input create() => UserDelete_Input._();
  @$core.override
  UserDelete_Input createEmptyInstance() => create();
  static $pb.PbList<UserDelete_Input> createRepeated() =>
      $pb.PbList<UserDelete_Input>();
  @$core.pragma('dart2js:noInline')
  static UserDelete_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserDelete_Input>(create);
  static UserDelete_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userName => $_getSZ(0);
  @$pb.TagNumber(1)
  set userName($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasUserName() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserName() => $_clearField(1);
}

class UserDelete_Output extends $pb.GeneratedMessage {
  factory UserDelete_Output() => create();

  UserDelete_Output._();

  factory UserDelete_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UserDelete_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserDelete.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete_Output clone() => UserDelete_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete_Output copyWith(void Function(UserDelete_Output) updates) =>
      super.copyWith((message) => updates(message as UserDelete_Output))
          as UserDelete_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserDelete_Output create() => UserDelete_Output._();
  @$core.override
  UserDelete_Output createEmptyInstance() => create();
  static $pb.PbList<UserDelete_Output> createRepeated() =>
      $pb.PbList<UserDelete_Output>();
  @$core.pragma('dart2js:noInline')
  static UserDelete_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserDelete_Output>(create);
  static UserDelete_Output? _defaultInstance;
}

/// delete user
class UserDelete extends $pb.GeneratedMessage {
  factory UserDelete() => create();

  UserDelete._();

  factory UserDelete.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UserDelete.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserDelete',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete clone() => UserDelete()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete copyWith(void Function(UserDelete) updates) =>
      super.copyWith((message) => updates(message as UserDelete)) as UserDelete;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserDelete create() => UserDelete._();
  @$core.override
  UserDelete createEmptyInstance() => create();
  static $pb.PbList<UserDelete> createRepeated() => $pb.PbList<UserDelete>();
  @$core.pragma('dart2js:noInline')
  static UserDelete getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserDelete>(create);
  static UserDelete? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
