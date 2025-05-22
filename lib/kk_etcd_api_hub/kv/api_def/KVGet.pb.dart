//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/kv/api_def/KVGet.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../kk_etcd_models/pb_kv_kk_etcd.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class KVGet_Input extends $pb.GeneratedMessage {
  factory KVGet_Input({
    $core.String? key,
  }) {
    final $result = create();
    if (key != null) {
      $result.key = key;
    }
    return $result;
  }
  KVGet_Input._() : super();
  factory KVGet_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVGet_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVGet.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVGet_Input clone() => KVGet_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVGet_Input copyWith(void Function(KVGet_Input) updates) => super.copyWith((message) => updates(message as KVGet_Input)) as KVGet_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVGet_Input create() => KVGet_Input._();
  KVGet_Input createEmptyInstance() => create();
  static $pb.PbList<KVGet_Input> createRepeated() => $pb.PbList<KVGet_Input>();
  @$core.pragma('dart2js:noInline')
  static KVGet_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVGet_Input>(create);
  static KVGet_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => $_clearField(1);
}

class KVGet_Output extends $pb.GeneratedMessage {
  factory KVGet_Output({
    $0.PBKV? kV,
  }) {
    final $result = create();
    if (kV != null) {
      $result.kV = kV;
    }
    return $result;
  }
  KVGet_Output._() : super();
  factory KVGet_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVGet_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVGet.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOM<$0.PBKV>(1, _omitFieldNames ? '' : 'KV', protoName: 'KV', subBuilder: $0.PBKV.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVGet_Output clone() => KVGet_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVGet_Output copyWith(void Function(KVGet_Output) updates) => super.copyWith((message) => updates(message as KVGet_Output)) as KVGet_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVGet_Output create() => KVGet_Output._();
  KVGet_Output createEmptyInstance() => create();
  static $pb.PbList<KVGet_Output> createRepeated() => $pb.PbList<KVGet_Output>();
  @$core.pragma('dart2js:noInline')
  static KVGet_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVGet_Output>(create);
  static KVGet_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBKV get kV => $_getN(0);
  @$pb.TagNumber(1)
  set kV($0.PBKV v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasKV() => $_has(0);
  @$pb.TagNumber(1)
  void clearKV() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBKV ensureKV() => $_ensure(0);
}

/// get kv
class KVGet extends $pb.GeneratedMessage {
  factory KVGet() => create();
  KVGet._() : super();
  factory KVGet.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVGet.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVGet', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVGet clone() => KVGet()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVGet copyWith(void Function(KVGet) updates) => super.copyWith((message) => updates(message as KVGet)) as KVGet;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVGet create() => KVGet._();
  KVGet createEmptyInstance() => create();
  static $pb.PbList<KVGet> createRepeated() => $pb.PbList<KVGet>();
  @$core.pragma('dart2js:noInline')
  static KVGet getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVGet>(create);
  static KVGet? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
