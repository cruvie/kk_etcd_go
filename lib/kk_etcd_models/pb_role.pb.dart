//
//  Generated code. Do not modify.
//  source: pb_role.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PBRole extends $pb.GeneratedMessage {
  factory PBRole({
    $core.String? name,
    $core.String? key,
    $core.String? rangeEnd,
    $core.int? permissionType,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    if (key != null) {
      $result.key = key;
    }
    if (rangeEnd != null) {
      $result.rangeEnd = rangeEnd;
    }
    if (permissionType != null) {
      $result.permissionType = permissionType;
    }
    return $result;
  }
  PBRole._() : super();
  factory PBRole.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBRole.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBRole', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..aOS(2, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..aOS(3, _omitFieldNames ? '' : 'RangeEnd', protoName: 'RangeEnd')
    ..a<$core.int>(4, _omitFieldNames ? '' : 'PermissionType', $pb.PbFieldType.O3, protoName: 'PermissionType')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PBRole clone() => PBRole()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PBRole copyWith(void Function(PBRole) updates) => super.copyWith((message) => updates(message as PBRole)) as PBRole;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBRole create() => PBRole._();
  PBRole createEmptyInstance() => create();
  static $pb.PbList<PBRole> createRepeated() => $pb.PbList<PBRole>();
  @$core.pragma('dart2js:noInline')
  static PBRole getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBRole>(create);
  static PBRole? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get key => $_getSZ(1);
  @$pb.TagNumber(2)
  set key($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearKey() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get rangeEnd => $_getSZ(2);
  @$pb.TagNumber(3)
  set rangeEnd($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRangeEnd() => $_has(2);
  @$pb.TagNumber(3)
  void clearRangeEnd() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get permissionType => $_getIZ(3);
  @$pb.TagNumber(4)
  set permissionType($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPermissionType() => $_has(3);
  @$pb.TagNumber(4)
  void clearPermissionType() => clearField(4);
}

class PBListRole extends $pb.GeneratedMessage {
  factory PBListRole({
    $core.Iterable<PBRole>? list,
  }) {
    final $result = create();
    if (list != null) {
      $result.list.addAll(list);
    }
    return $result;
  }
  PBListRole._() : super();
  factory PBListRole.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBListRole.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBListRole', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..pc<PBRole>(1, _omitFieldNames ? '' : 'List', $pb.PbFieldType.PM, protoName: 'List', subBuilder: PBRole.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PBListRole clone() => PBListRole()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PBListRole copyWith(void Function(PBListRole) updates) => super.copyWith((message) => updates(message as PBListRole)) as PBListRole;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListRole create() => PBListRole._();
  PBListRole createEmptyInstance() => create();
  static $pb.PbList<PBListRole> createRepeated() => $pb.PbList<PBListRole>();
  @$core.pragma('dart2js:noInline')
  static PBListRole getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBListRole>(create);
  static PBListRole? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<PBRole> get list => $_getList(0);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
