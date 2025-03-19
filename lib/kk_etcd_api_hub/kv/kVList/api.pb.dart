//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/kv/kVList/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../kk_etcd_models/pb_kv_kk_etcd.pb.dart' as $0;

class KVList_Input extends $pb.GeneratedMessage {
  factory KVList_Input({
    $core.String? prefix,
  }) {
    final $result = create();
    if (prefix != null) {
      $result.prefix = prefix;
    }
    return $result;
  }

  KVList_Input._() : super();

  factory KVList_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory KVList_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'KVList.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kVList'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Prefix', protoName: 'Prefix')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  KVList_Input clone() => KVList_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  KVList_Input copyWith(void Function(KVList_Input) updates) =>
      super.copyWith((message) => updates(message as KVList_Input))
          as KVList_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVList_Input create() => KVList_Input._();

  KVList_Input createEmptyInstance() => create();

  static $pb.PbList<KVList_Input> createRepeated() =>
      $pb.PbList<KVList_Input>();

  @$core.pragma('dart2js:noInline')
  static KVList_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<KVList_Input>(create);
  static KVList_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get prefix => $_getSZ(0);

  @$pb.TagNumber(1)
  set prefix($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasPrefix() => $_has(0);

  @$pb.TagNumber(1)
  void clearPrefix() => clearField(1);
}

class KVList_Output extends $pb.GeneratedMessage {
  factory KVList_Output({
    $0.PBListKV? kVList,
  }) {
    final $result = create();
    if (kVList != null) {
      $result.kVList = kVList;
    }
    return $result;
  }

  KVList_Output._() : super();

  factory KVList_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory KVList_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'KVList.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kVList'),
      createEmptyInstance: create)
    ..aOM<$0.PBListKV>(1, _omitFieldNames ? '' : 'KVList',
        protoName: 'KVList', subBuilder: $0.PBListKV.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  KVList_Output clone() => KVList_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  KVList_Output copyWith(void Function(KVList_Output) updates) =>
      super.copyWith((message) => updates(message as KVList_Output))
          as KVList_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVList_Output create() => KVList_Output._();

  KVList_Output createEmptyInstance() => create();

  static $pb.PbList<KVList_Output> createRepeated() =>
      $pb.PbList<KVList_Output>();

  @$core.pragma('dart2js:noInline')
  static KVList_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<KVList_Output>(create);
  static KVList_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBListKV get kVList => $_getN(0);

  @$pb.TagNumber(1)
  set kVList($0.PBListKV v) {
    setField(1, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasKVList() => $_has(0);

  @$pb.TagNumber(1)
  void clearKVList() => clearField(1);

  @$pb.TagNumber(1)
  $0.PBListKV ensureKVList() => $_ensure(0);
}

/// list kv
class KVList extends $pb.GeneratedMessage {
  factory KVList() => create();

  KVList._() : super();

  factory KVList.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory KVList.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'KVList',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kVList'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  KVList clone() => KVList()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  KVList copyWith(void Function(KVList) updates) =>
      super.copyWith((message) => updates(message as KVList)) as KVList;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVList create() => KVList._();

  KVList createEmptyInstance() => create();

  static $pb.PbList<KVList> createRepeated() => $pb.PbList<KVList>();

  @$core.pragma('dart2js:noInline')
  static KVList getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVList>(create);
  static KVList? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
