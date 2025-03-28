//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/role/roleGrantPermission/api.proto
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

class RoleGrantPermission_Input extends $pb.GeneratedMessage {
  factory RoleGrantPermission_Input({
    $core.String? name,
    $0.PBPermission? perm,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    if (perm != null) {
      $result.perm = perm;
    }
    return $result;
  }
  RoleGrantPermission_Input._() : super();
  factory RoleGrantPermission_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGrantPermission_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGrantPermission.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'roleGrantPermission'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..aOM<$0.PBPermission>(2, _omitFieldNames ? '' : 'Perm', protoName: 'Perm', subBuilder: $0.PBPermission.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGrantPermission_Input clone() => RoleGrantPermission_Input()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGrantPermission_Input copyWith(void Function(RoleGrantPermission_Input) updates) => super.copyWith((message) => updates(message as RoleGrantPermission_Input)) as RoleGrantPermission_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission_Input create() => RoleGrantPermission_Input._();
  RoleGrantPermission_Input createEmptyInstance() => create();
  static $pb.PbList<RoleGrantPermission_Input> createRepeated() => $pb.PbList<RoleGrantPermission_Input>();
  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGrantPermission_Input>(create);
  static RoleGrantPermission_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);

  @$pb.TagNumber(2)
  $0.PBPermission get perm => $_getN(1);
  @$pb.TagNumber(2)
  set perm($0.PBPermission v) { $_setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasPerm() => $_has(1);
  @$pb.TagNumber(2)
  void clearPerm() => $_clearField(2);
  @$pb.TagNumber(2)
  $0.PBPermission ensurePerm() => $_ensure(1);
}

class RoleGrantPermission_Output extends $pb.GeneratedMessage {
  factory RoleGrantPermission_Output() => create();
  RoleGrantPermission_Output._() : super();
  factory RoleGrantPermission_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGrantPermission_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGrantPermission.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'roleGrantPermission'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGrantPermission_Output clone() => RoleGrantPermission_Output()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGrantPermission_Output copyWith(void Function(RoleGrantPermission_Output) updates) => super.copyWith((message) => updates(message as RoleGrantPermission_Output)) as RoleGrantPermission_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission_Output create() => RoleGrantPermission_Output._();
  RoleGrantPermission_Output createEmptyInstance() => create();
  static $pb.PbList<RoleGrantPermission_Output> createRepeated() => $pb.PbList<RoleGrantPermission_Output>();
  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGrantPermission_Output>(create);
  static RoleGrantPermission_Output? _defaultInstance;
}

/// grant permission
class RoleGrantPermission extends $pb.GeneratedMessage {
  factory RoleGrantPermission() => create();
  RoleGrantPermission._() : super();
  factory RoleGrantPermission.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGrantPermission.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGrantPermission', package: const $pb.PackageName(_omitMessageNames ? '' : 'roleGrantPermission'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGrantPermission clone() => RoleGrantPermission()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGrantPermission copyWith(void Function(RoleGrantPermission) updates) => super.copyWith((message) => updates(message as RoleGrantPermission)) as RoleGrantPermission;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission create() => RoleGrantPermission._();
  RoleGrantPermission createEmptyInstance() => create();
  static $pb.PbList<RoleGrantPermission> createRepeated() => $pb.PbList<RoleGrantPermission>();
  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGrantPermission>(create);
  static RoleGrantPermission? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
