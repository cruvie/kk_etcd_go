//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_user/userGrantRole/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class UserGrantRole_Input extends $pb.GeneratedMessage {
  factory UserGrantRole_Input({
    $core.String? userName,
    $core.Iterable<$core.String>? roles,
  }) {
    final $result = create();
    if (userName != null) {
      $result.userName = userName;
    }
    if (roles != null) {
      $result.roles.addAll(roles);
    }
    return $result;
  }

  UserGrantRole_Input._() : super();

  factory UserGrantRole_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory UserGrantRole_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserGrantRole.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'userGrantRole'),
      createEmptyInstance: create)
    ..aOS(2, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..pPS(4, _omitFieldNames ? '' : 'Roles', protoName: 'Roles')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserGrantRole_Input clone() => UserGrantRole_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserGrantRole_Input copyWith(void Function(UserGrantRole_Input) updates) =>
      super.copyWith((message) => updates(message as UserGrantRole_Input))
          as UserGrantRole_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserGrantRole_Input create() => UserGrantRole_Input._();

  UserGrantRole_Input createEmptyInstance() => create();

  static $pb.PbList<UserGrantRole_Input> createRepeated() =>
      $pb.PbList<UserGrantRole_Input>();

  @$core.pragma('dart2js:noInline')
  static UserGrantRole_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserGrantRole_Input>(create);
  static UserGrantRole_Input? _defaultInstance;

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

  @$pb.TagNumber(4)
  $core.List<$core.String> get roles => $_getList(1);
}

class UserGrantRole_Output extends $pb.GeneratedMessage {
  factory UserGrantRole_Output() => create();

  UserGrantRole_Output._() : super();

  factory UserGrantRole_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory UserGrantRole_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserGrantRole.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'userGrantRole'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserGrantRole_Output clone() =>
      UserGrantRole_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserGrantRole_Output copyWith(void Function(UserGrantRole_Output) updates) =>
      super.copyWith((message) => updates(message as UserGrantRole_Output))
          as UserGrantRole_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserGrantRole_Output create() => UserGrantRole_Output._();

  UserGrantRole_Output createEmptyInstance() => create();

  static $pb.PbList<UserGrantRole_Output> createRepeated() =>
      $pb.PbList<UserGrantRole_Output>();

  @$core.pragma('dart2js:noInline')
  static UserGrantRole_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserGrantRole_Output>(create);
  static UserGrantRole_Output? _defaultInstance;
}

/// grant role
class UserGrantRole extends $pb.GeneratedMessage {
  factory UserGrantRole() => create();

  UserGrantRole._() : super();

  factory UserGrantRole.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory UserGrantRole.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserGrantRole',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'userGrantRole'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserGrantRole clone() => UserGrantRole()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserGrantRole copyWith(void Function(UserGrantRole) updates) =>
      super.copyWith((message) => updates(message as UserGrantRole))
          as UserGrantRole;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserGrantRole create() => UserGrantRole._();

  UserGrantRole createEmptyInstance() => create();

  static $pb.PbList<UserGrantRole> createRepeated() =>
      $pb.PbList<UserGrantRole>();

  @$core.pragma('dart2js:noInline')
  static UserGrantRole getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserGrantRole>(create);
  static UserGrantRole? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
