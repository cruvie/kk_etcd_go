//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_user_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PBUser extends $pb.GeneratedMessage {
  factory PBUser({
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
  PBUser._() : super();

  factory PBUser.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory PBUser.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBUser',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'),
      createEmptyInstance: create)
    ..aOS(2, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..aOS(3, _omitFieldNames ? '' : 'Password', protoName: 'Password')
    ..pPS(4, _omitFieldNames ? '' : 'Roles', protoName: 'Roles')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  PBUser clone() => PBUser()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  PBUser copyWith(void Function(PBUser) updates) =>
      super.copyWith((message) => updates(message as PBUser)) as PBUser;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBUser create() => PBUser._();
  PBUser createEmptyInstance() => create();
  static $pb.PbList<PBUser> createRepeated() => $pb.PbList<PBUser>();
  @$core.pragma('dart2js:noInline')
  static PBUser getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBUser>(create);
  static PBUser? _defaultInstance;

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

class PBListUser extends $pb.GeneratedMessage {
  factory PBListUser({
    $core.Iterable<PBUser>? listUser,
  }) {
    final $result = create();
    if (listUser != null) {
      $result.listUser.addAll(listUser);
    }
    return $result;
  }
  PBListUser._() : super();

  factory PBListUser.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory PBListUser.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBListUser',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'),
      createEmptyInstance: create)
    ..pc<PBUser>(1, _omitFieldNames ? '' : 'ListUser', $pb.PbFieldType.PM,
        protoName: 'ListUser', subBuilder: PBUser.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  PBListUser clone() => PBListUser()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  PBListUser copyWith(void Function(PBListUser) updates) =>
      super.copyWith((message) => updates(message as PBListUser)) as PBListUser;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListUser create() => PBListUser._();
  PBListUser createEmptyInstance() => create();
  static $pb.PbList<PBListUser> createRepeated() => $pb.PbList<PBListUser>();
  @$core.pragma('dart2js:noInline')
  static PBListUser getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<PBListUser>(create);
  static PBListUser? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<PBUser> get listUser => $_getList(0);
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
