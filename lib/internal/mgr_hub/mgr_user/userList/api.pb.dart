//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_user/userList/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../../kk_etcd_models/pb_user_kk_etcd.pb.dart' as $0;

class UserList_Input extends $pb.GeneratedMessage {
  factory UserList_Input() => create();

  UserList_Input._() : super();

  factory UserList_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory UserList_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserList.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'userList'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserList_Input clone() => UserList_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserList_Input copyWith(void Function(UserList_Input) updates) =>
      super.copyWith((message) => updates(message as UserList_Input))
          as UserList_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserList_Input create() => UserList_Input._();

  UserList_Input createEmptyInstance() => create();

  static $pb.PbList<UserList_Input> createRepeated() =>
      $pb.PbList<UserList_Input>();

  @$core.pragma('dart2js:noInline')
  static UserList_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserList_Input>(create);
  static UserList_Input? _defaultInstance;
}

class UserList_Output extends $pb.GeneratedMessage {
  factory UserList_Output({
    $0.PBListUser? listUser,
  }) {
    final $result = create();
    if (listUser != null) {
      $result.listUser = listUser;
    }
    return $result;
  }

  UserList_Output._() : super();

  factory UserList_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory UserList_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserList.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'userList'),
      createEmptyInstance: create)
    ..aOM<$0.PBListUser>(1, _omitFieldNames ? '' : 'ListUser',
        protoName: 'ListUser', subBuilder: $0.PBListUser.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserList_Output clone() => UserList_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserList_Output copyWith(void Function(UserList_Output) updates) =>
      super.copyWith((message) => updates(message as UserList_Output))
          as UserList_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserList_Output create() => UserList_Output._();

  UserList_Output createEmptyInstance() => create();

  static $pb.PbList<UserList_Output> createRepeated() =>
      $pb.PbList<UserList_Output>();

  @$core.pragma('dart2js:noInline')
  static UserList_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserList_Output>(create);
  static UserList_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBListUser get listUser => $_getN(0);

  @$pb.TagNumber(1)
  set listUser($0.PBListUser v) {
    setField(1, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasListUser() => $_has(0);

  @$pb.TagNumber(1)
  void clearListUser() => clearField(1);

  @$pb.TagNumber(1)
  $0.PBListUser ensureListUser() => $_ensure(0);
}

/// list user
class UserList extends $pb.GeneratedMessage {
  factory UserList() => create();

  UserList._() : super();

  factory UserList.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory UserList.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserList',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'userList'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  UserList clone() => UserList()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  UserList copyWith(void Function(UserList) updates) =>
      super.copyWith((message) => updates(message as UserList)) as UserList;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserList create() => UserList._();

  UserList createEmptyInstance() => create();

  static $pb.PbList<UserList> createRepeated() => $pb.PbList<UserList>();

  @$core.pragma('dart2js:noInline')
  static UserList getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserList>(create);
  static UserList? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
