//
//  Generated code. Do not modify.
//  source: api_kv.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'pb_kv.pb.dart' as $0;

class KVPutParam extends $pb.GeneratedMessage {
  factory KVPutParam({
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
  KVPutParam._() : super();
  factory KVPutParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVPutParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVPutParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..aOS(2, _omitFieldNames ? '' : 'Value', protoName: 'Value')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVPutParam clone() => KVPutParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVPutParam copyWith(void Function(KVPutParam) updates) => super.copyWith((message) => updates(message as KVPutParam)) as KVPutParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVPutParam create() => KVPutParam._();
  KVPutParam createEmptyInstance() => create();
  static $pb.PbList<KVPutParam> createRepeated() => $pb.PbList<KVPutParam>();
  @$core.pragma('dart2js:noInline')
  static KVPutParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVPutParam>(create);
  static KVPutParam? _defaultInstance;

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

class KVPutResponse extends $pb.GeneratedMessage {
  factory KVPutResponse() => create();
  KVPutResponse._() : super();
  factory KVPutResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVPutResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVPutResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVPutResponse clone() => KVPutResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVPutResponse copyWith(void Function(KVPutResponse) updates) => super.copyWith((message) => updates(message as KVPutResponse)) as KVPutResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVPutResponse create() => KVPutResponse._();
  KVPutResponse createEmptyInstance() => create();
  static $pb.PbList<KVPutResponse> createRepeated() => $pb.PbList<KVPutResponse>();
  @$core.pragma('dart2js:noInline')
  static KVPutResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVPutResponse>(create);
  static KVPutResponse? _defaultInstance;
}

class KVGetParam extends $pb.GeneratedMessage {
  factory KVGetParam({
    $core.String? key,
  }) {
    final $result = create();
    if (key != null) {
      $result.key = key;
    }
    return $result;
  }
  KVGetParam._() : super();
  factory KVGetParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVGetParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVGetParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVGetParam clone() => KVGetParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVGetParam copyWith(void Function(KVGetParam) updates) => super.copyWith((message) => updates(message as KVGetParam)) as KVGetParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVGetParam create() => KVGetParam._();
  KVGetParam createEmptyInstance() => create();
  static $pb.PbList<KVGetParam> createRepeated() => $pb.PbList<KVGetParam>();
  @$core.pragma('dart2js:noInline')
  static KVGetParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVGetParam>(create);
  static KVGetParam? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => clearField(1);
}

class KVGetResponse extends $pb.GeneratedMessage {
  factory KVGetResponse({
    $0.PBKV? kV,
  }) {
    final $result = create();
    if (kV != null) {
      $result.kV = kV;
    }
    return $result;
  }
  KVGetResponse._() : super();
  factory KVGetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVGetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVGetResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOM<$0.PBKV>(1, _omitFieldNames ? '' : 'KV', protoName: 'KV', subBuilder: $0.PBKV.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVGetResponse clone() => KVGetResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVGetResponse copyWith(void Function(KVGetResponse) updates) => super.copyWith((message) => updates(message as KVGetResponse)) as KVGetResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVGetResponse create() => KVGetResponse._();
  KVGetResponse createEmptyInstance() => create();
  static $pb.PbList<KVGetResponse> createRepeated() => $pb.PbList<KVGetResponse>();
  @$core.pragma('dart2js:noInline')
  static KVGetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVGetResponse>(create);
  static KVGetResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBKV get kV => $_getN(0);
  @$pb.TagNumber(1)
  set kV($0.PBKV v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasKV() => $_has(0);
  @$pb.TagNumber(1)
  void clearKV() => clearField(1);
  @$pb.TagNumber(1)
  $0.PBKV ensureKV() => $_ensure(0);
}

class KVListParam extends $pb.GeneratedMessage {
  factory KVListParam({
    $core.String? prefix,
  }) {
    final $result = create();
    if (prefix != null) {
      $result.prefix = prefix;
    }
    return $result;
  }
  KVListParam._() : super();
  factory KVListParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVListParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVListParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Prefix', protoName: 'Prefix')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVListParam clone() => KVListParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVListParam copyWith(void Function(KVListParam) updates) => super.copyWith((message) => updates(message as KVListParam)) as KVListParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVListParam create() => KVListParam._();
  KVListParam createEmptyInstance() => create();
  static $pb.PbList<KVListParam> createRepeated() => $pb.PbList<KVListParam>();
  @$core.pragma('dart2js:noInline')
  static KVListParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVListParam>(create);
  static KVListParam? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get prefix => $_getSZ(0);
  @$pb.TagNumber(1)
  set prefix($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasPrefix() => $_has(0);
  @$pb.TagNumber(1)
  void clearPrefix() => clearField(1);
}

class KVListResponse extends $pb.GeneratedMessage {
  factory KVListResponse({
    $0.PBListKV? kVList,
  }) {
    final $result = create();
    if (kVList != null) {
      $result.kVList = kVList;
    }
    return $result;
  }
  KVListResponse._() : super();
  factory KVListResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVListResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVListResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOM<$0.PBListKV>(1, _omitFieldNames ? '' : 'KVList', protoName: 'KVList', subBuilder: $0.PBListKV.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVListResponse clone() => KVListResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVListResponse copyWith(void Function(KVListResponse) updates) => super.copyWith((message) => updates(message as KVListResponse)) as KVListResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVListResponse create() => KVListResponse._();
  KVListResponse createEmptyInstance() => create();
  static $pb.PbList<KVListResponse> createRepeated() => $pb.PbList<KVListResponse>();
  @$core.pragma('dart2js:noInline')
  static KVListResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVListResponse>(create);
  static KVListResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBListKV get kVList => $_getN(0);
  @$pb.TagNumber(1)
  set kVList($0.PBListKV v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasKVList() => $_has(0);
  @$pb.TagNumber(1)
  void clearKVList() => clearField(1);
  @$pb.TagNumber(1)
  $0.PBListKV ensureKVList() => $_ensure(0);
}

class KVDelParam extends $pb.GeneratedMessage {
  factory KVDelParam({
    $core.String? key,
  }) {
    final $result = create();
    if (key != null) {
      $result.key = key;
    }
    return $result;
  }
  KVDelParam._() : super();
  factory KVDelParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVDelParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVDelParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVDelParam clone() => KVDelParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVDelParam copyWith(void Function(KVDelParam) updates) => super.copyWith((message) => updates(message as KVDelParam)) as KVDelParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVDelParam create() => KVDelParam._();
  KVDelParam createEmptyInstance() => create();
  static $pb.PbList<KVDelParam> createRepeated() => $pb.PbList<KVDelParam>();
  @$core.pragma('dart2js:noInline')
  static KVDelParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVDelParam>(create);
  static KVDelParam? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => clearField(1);
}

class KVDelResponse extends $pb.GeneratedMessage {
  factory KVDelResponse() => create();
  KVDelResponse._() : super();
  factory KVDelResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVDelResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVDelResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVDelResponse clone() => KVDelResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVDelResponse copyWith(void Function(KVDelResponse) updates) => super.copyWith((message) => updates(message as KVDelResponse)) as KVDelResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVDelResponse create() => KVDelResponse._();
  KVDelResponse createEmptyInstance() => create();
  static $pb.PbList<KVDelResponse> createRepeated() => $pb.PbList<KVDelResponse>();
  @$core.pragma('dart2js:noInline')
  static KVDelResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVDelResponse>(create);
  static KVDelResponse? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
