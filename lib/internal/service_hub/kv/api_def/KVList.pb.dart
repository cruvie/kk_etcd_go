// This is a generated file - do not edit.
//
// Generated from internal/service_hub/kv/api_def/KVList.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../../kk_etcd_models/pb_kv_kk_etcd.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class KVList_Input extends $pb.GeneratedMessage {
  factory KVList_Input({
    $core.String? prefix,
  }) {
    final result = create();
    if (prefix != null) result.prefix = prefix;
    return result;
  }

  KVList_Input._();

  factory KVList_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory KVList_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'KVList.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Prefix', protoName: 'Prefix')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVList_Input clone() => KVList_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVList_Input copyWith(void Function(KVList_Input) updates) =>
      super.copyWith((message) => updates(message as KVList_Input))
          as KVList_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVList_Input create() => KVList_Input._();
  @$core.override
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
  set prefix($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasPrefix() => $_has(0);
  @$pb.TagNumber(1)
  void clearPrefix() => $_clearField(1);
}

class KVList_Output extends $pb.GeneratedMessage {
  factory KVList_Output({
    $0.PBListKV? kVList,
  }) {
    final result = create();
    if (kVList != null) result.kVList = kVList;
    return result;
  }

  KVList_Output._();

  factory KVList_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory KVList_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'KVList.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..aOM<$0.PBListKV>(1, _omitFieldNames ? '' : 'KVList',
        protoName: 'KVList', subBuilder: $0.PBListKV.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVList_Output clone() => KVList_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVList_Output copyWith(void Function(KVList_Output) updates) =>
      super.copyWith((message) => updates(message as KVList_Output))
          as KVList_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVList_Output create() => KVList_Output._();
  @$core.override
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
  set kVList($0.PBListKV value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasKVList() => $_has(0);
  @$pb.TagNumber(1)
  void clearKVList() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBListKV ensureKVList() => $_ensure(0);
}

/// list kv
class KVList extends $pb.GeneratedMessage {
  factory KVList() => create();

  KVList._();

  factory KVList.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory KVList.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'KVList',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVList clone() => KVList()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVList copyWith(void Function(KVList) updates) =>
      super.copyWith((message) => updates(message as KVList)) as KVList;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVList create() => KVList._();
  @$core.override
  KVList createEmptyInstance() => create();
  static $pb.PbList<KVList> createRepeated() => $pb.PbList<KVList>();
  @$core.pragma('dart2js:noInline')
  static KVList getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVList>(create);
  static KVList? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
