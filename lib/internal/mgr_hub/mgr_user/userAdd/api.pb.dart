//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_user/userAdd/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class UserAdd_Input extends $pb.GeneratedMessage {
  factory UserAdd_Input({
    $core.String? userName,
    $core.String? password,
    $core.Iterable<$core.String>? roles,
  }) {
    final $result = create();
    if (userName != null) {
      $result.userName = userName;
    }
    if (password != null) {
      $result.password = password;
    }
    if (roles != null) {
      $result.roles.addAll(roles);
    }
    return $result;
  }

  UserAdd_Input._() : super();

  factory UserAdd_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory UserAdd_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserAdd.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'userAdd'),
      createEmptyInstance: create)
    ..aOS(2, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..aOS(3, _omitFieldNames ? '' : 'Password', protoName: 'Password')
    ..pPS(4, _omitFieldNames ? '' : 'Roles', protoName: 'Roles')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserAdd_Input clone() => UserAdd_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserAdd_Input copyWith(void Function(UserAdd_Input) updates) =>
      super.copyWith((message) => updates(message as UserAdd_Input))
          as UserAdd_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserAdd_Input create() => UserAdd_Input._();

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
  set userName($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasUserName() => $_has(0);

  @$pb.TagNumber(2)
  void clearUserName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get password => $_getSZ(1);

  @$pb.TagNumber(3)
  set password($core.String v) {
    $_setString(1, v);
  }

  @$pb.TagNumber(3)
  $core.bool hasPassword() => $_has(1);

  @$pb.TagNumber(3)
  void clearPassword() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<$core.String> get roles => $_getList(2);
}

class UserAdd_Output extends $pb.GeneratedMessage {
  factory UserAdd_Output() => create();

  UserAdd_Output._() : super();

  factory UserAdd_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory UserAdd_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserAdd.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'userAdd'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserAdd_Output clone() => UserAdd_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserAdd_Output copyWith(void Function(UserAdd_Output) updates) =>
      super.copyWith((message) => updates(message as UserAdd_Output))
          as UserAdd_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserAdd_Output create() => UserAdd_Output._();

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

  UserAdd._() : super();

  factory UserAdd.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory UserAdd.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserAdd',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'userAdd'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserAdd clone() => UserAdd()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserAdd copyWith(void Function(UserAdd) updates) =>
      super.copyWith((message) => updates(message as UserAdd)) as UserAdd;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserAdd create() => UserAdd._();

  UserAdd createEmptyInstance() => create();

  static $pb.PbList<UserAdd> createRepeated() => $pb.PbList<UserAdd>();

  @$core.pragma('dart2js:noInline')
  static UserAdd getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserAdd>(create);
  static UserAdd? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
