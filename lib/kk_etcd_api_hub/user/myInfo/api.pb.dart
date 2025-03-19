//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/user/myInfo/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class MyInfo_Input extends $pb.GeneratedMessage {
  factory MyInfo_Input() => create();

  MyInfo_Input._() : super();

  factory MyInfo_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory MyInfo_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'MyInfo.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'myInfo'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  MyInfo_Input clone() => MyInfo_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  MyInfo_Input copyWith(void Function(MyInfo_Input) updates) =>
      super.copyWith((message) => updates(message as MyInfo_Input))
          as MyInfo_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static MyInfo_Input create() => MyInfo_Input._();

  MyInfo_Input createEmptyInstance() => create();

  static $pb.PbList<MyInfo_Input> createRepeated() =>
      $pb.PbList<MyInfo_Input>();

  @$core.pragma('dart2js:noInline')
  static MyInfo_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<MyInfo_Input>(create);
  static MyInfo_Input? _defaultInstance;
}

class MyInfo_Output extends $pb.GeneratedMessage {
  factory MyInfo_Output({
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

  MyInfo_Output._() : super();

  factory MyInfo_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory MyInfo_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'MyInfo.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'myInfo'),
      createEmptyInstance: create)
    ..aOS(2, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..pPS(4, _omitFieldNames ? '' : 'Roles', protoName: 'Roles')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  MyInfo_Output clone() => MyInfo_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  MyInfo_Output copyWith(void Function(MyInfo_Output) updates) =>
      super.copyWith((message) => updates(message as MyInfo_Output))
          as MyInfo_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static MyInfo_Output create() => MyInfo_Output._();

  MyInfo_Output createEmptyInstance() => create();

  static $pb.PbList<MyInfo_Output> createRepeated() =>
      $pb.PbList<MyInfo_Output>();

  @$core.pragma('dart2js:noInline')
  static MyInfo_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<MyInfo_Output>(create);
  static MyInfo_Output? _defaultInstance;

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

/// get my info
class MyInfo extends $pb.GeneratedMessage {
  factory MyInfo() => create();

  MyInfo._() : super();

  factory MyInfo.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory MyInfo.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'MyInfo',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'myInfo'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  MyInfo clone() => MyInfo()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  MyInfo copyWith(void Function(MyInfo) updates) =>
      super.copyWith((message) => updates(message as MyInfo)) as MyInfo;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static MyInfo create() => MyInfo._();

  MyInfo createEmptyInstance() => create();

  static $pb.PbList<MyInfo> createRepeated() => $pb.PbList<MyInfo>();

  @$core.pragma('dart2js:noInline')
  static MyInfo getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MyInfo>(create);
  static MyInfo? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
