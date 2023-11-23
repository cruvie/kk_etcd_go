//
//  Generated code. Do not modify.
//  source: pb_kv.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PBKV extends $pb.GeneratedMessage {
  factory PBKV({
    $core.String? key,
    $core.String? value,
  }) {
    final $result = create();
    if (key != null) {
      $result.key = key;
    }
    if (value != null) {
      $result.value = value;
    }
    return $result;
  }
  PBKV._() : super();
  factory PBKV.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBKV.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBKV', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..aOS(2, _omitFieldNames ? '' : 'Value', protoName: 'Value')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PBKV clone() => PBKV()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PBKV copyWith(void Function(PBKV) updates) => super.copyWith((message) => updates(message as PBKV)) as PBKV;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBKV create() => PBKV._();
  PBKV createEmptyInstance() => create();
  static $pb.PbList<PBKV> createRepeated() => $pb.PbList<PBKV>();
  @$core.pragma('dart2js:noInline')
  static PBKV getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBKV>(create);
  static PBKV? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get value => $_getSZ(1);
  @$pb.TagNumber(2)
  set value($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => clearField(2);
}

class PBListKV extends $pb.GeneratedMessage {
  factory PBListKV({
    $core.Iterable<PBKV>? listKV,
  }) {
    final $result = create();
    if (listKV != null) {
      $result.listKV.addAll(listKV);
    }
    return $result;
  }
  PBListKV._() : super();
  factory PBListKV.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBListKV.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBListKV', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..pc<PBKV>(1, _omitFieldNames ? '' : 'ListKV', $pb.PbFieldType.PM, protoName: 'ListKV', subBuilder: PBKV.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PBListKV clone() => PBListKV()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PBListKV copyWith(void Function(PBListKV) updates) => super.copyWith((message) => updates(message as PBListKV)) as PBListKV;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListKV create() => PBListKV._();
  PBListKV createEmptyInstance() => create();
  static $pb.PbList<PBListKV> createRepeated() => $pb.PbList<PBListKV>();
  @$core.pragma('dart2js:noInline')
  static PBListKV getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBListKV>(create);
  static PBListKV? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<PBKV> get listKV => $_getList(0);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
