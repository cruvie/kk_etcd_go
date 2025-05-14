//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/role/api_def/RoleGet.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../kk_etcd_models/pb_role_kk_etcd.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class RoleGet_Input extends $pb.GeneratedMessage {
  factory RoleGet_Input({
    $core.String? name,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  RoleGet_Input._() : super();
  factory RoleGet_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGet_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGet.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGet_Input clone() => RoleGet_Input()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGet_Input copyWith(void Function(RoleGet_Input) updates) => super.copyWith((message) => updates(message as RoleGet_Input)) as RoleGet_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGet_Input create() => RoleGet_Input._();
  RoleGet_Input createEmptyInstance() => create();
  static $pb.PbList<RoleGet_Input> createRepeated() => $pb.PbList<RoleGet_Input>();
  @$core.pragma('dart2js:noInline')
  static RoleGet_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGet_Input>(create);
  static RoleGet_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);
}

class RoleGet_Output extends $pb.GeneratedMessage {
  factory RoleGet_Output({
    $0.PBRole? role,
  }) {
    final $result = create();
    if (role != null) {
      $result.role = role;
    }
    return $result;
  }
  RoleGet_Output._() : super();
  factory RoleGet_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGet_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGet.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOM<$0.PBRole>(1, _omitFieldNames ? '' : 'Role', protoName: 'Role', subBuilder: $0.PBRole.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGet_Output clone() => RoleGet_Output()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGet_Output copyWith(void Function(RoleGet_Output) updates) => super.copyWith((message) => updates(message as RoleGet_Output)) as RoleGet_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGet_Output create() => RoleGet_Output._();
  RoleGet_Output createEmptyInstance() => create();
  static $pb.PbList<RoleGet_Output> createRepeated() => $pb.PbList<RoleGet_Output>();
  @$core.pragma('dart2js:noInline')
  static RoleGet_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGet_Output>(create);
  static RoleGet_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBRole get role => $_getN(0);
  @$pb.TagNumber(1)
  set role($0.PBRole v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasRole() => $_has(0);
  @$pb.TagNumber(1)
  void clearRole() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBRole ensureRole() => $_ensure(0);
}

/// get role
class RoleGet extends $pb.GeneratedMessage {
  factory RoleGet() => create();
  RoleGet._() : super();
  factory RoleGet.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGet.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGet', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGet clone() => RoleGet()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGet copyWith(void Function(RoleGet) updates) => super.copyWith((message) => updates(message as RoleGet)) as RoleGet;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGet create() => RoleGet._();
  RoleGet createEmptyInstance() => create();
  static $pb.PbList<RoleGet> createRepeated() => $pb.PbList<RoleGet>();
  @$core.pragma('dart2js:noInline')
  static RoleGet getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGet>(create);
  static RoleGet? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
