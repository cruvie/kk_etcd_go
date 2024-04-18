//
//  Generated code. Do not modify.
//  source: api_manage_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class SnapshotParam extends $pb.GeneratedMessage {
  factory SnapshotParam() => create();
  SnapshotParam._() : super();
  factory SnapshotParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SnapshotParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotParam clone() => SnapshotParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotParam copyWith(void Function(SnapshotParam) updates) => super.copyWith((message) => updates(message as SnapshotParam)) as SnapshotParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotParam create() => SnapshotParam._();
  SnapshotParam createEmptyInstance() => create();
  static $pb.PbList<SnapshotParam> createRepeated() => $pb.PbList<SnapshotParam>();
  @$core.pragma('dart2js:noInline')
  static SnapshotParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotParam>(create);
  static SnapshotParam? _defaultInstance;
}

class SnapshotResponse extends $pb.GeneratedMessage {
  factory SnapshotResponse({
    $core.String? name,
    $core.List<$core.int>? file,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    if (file != null) {
      $result.file = file;
    }
    return $result;
  }
  SnapshotResponse._() : super();
  factory SnapshotResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SnapshotResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..a<$core.List<$core.int>>(2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY, protoName: 'File')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotResponse clone() => SnapshotResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotResponse copyWith(void Function(SnapshotResponse) updates) => super.copyWith((message) => updates(message as SnapshotResponse)) as SnapshotResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotResponse create() => SnapshotResponse._();
  SnapshotResponse createEmptyInstance() => create();
  static $pb.PbList<SnapshotResponse> createRepeated() => $pb.PbList<SnapshotResponse>();
  @$core.pragma('dart2js:noInline')
  static SnapshotResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotResponse>(create);
  static SnapshotResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get file => $_getN(1);
  @$pb.TagNumber(2)
  set file($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(1);
  @$pb.TagNumber(2)
  void clearFile() => clearField(2);
}

class SnapshotRestoreParam extends $pb.GeneratedMessage {
  factory SnapshotRestoreParam() => create();
  SnapshotRestoreParam._() : super();
  factory SnapshotRestoreParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotRestoreParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SnapshotRestoreParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotRestoreParam clone() => SnapshotRestoreParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotRestoreParam copyWith(void Function(SnapshotRestoreParam) updates) => super.copyWith((message) => updates(message as SnapshotRestoreParam)) as SnapshotRestoreParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotRestoreParam create() => SnapshotRestoreParam._();
  SnapshotRestoreParam createEmptyInstance() => create();
  static $pb.PbList<SnapshotRestoreParam> createRepeated() => $pb.PbList<SnapshotRestoreParam>();
  @$core.pragma('dart2js:noInline')
  static SnapshotRestoreParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotRestoreParam>(create);
  static SnapshotRestoreParam? _defaultInstance;
}

class SnapshotRestoreResponse extends $pb.GeneratedMessage {
  factory SnapshotRestoreResponse({
    $core.String? cmdStr,
  }) {
    final $result = create();
    if (cmdStr != null) {
      $result.cmdStr = cmdStr;
    }
    return $result;
  }
  SnapshotRestoreResponse._() : super();
  factory SnapshotRestoreResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotRestoreResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SnapshotRestoreResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'CmdStr', protoName: 'CmdStr')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotRestoreResponse clone() => SnapshotRestoreResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotRestoreResponse copyWith(void Function(SnapshotRestoreResponse) updates) => super.copyWith((message) => updates(message as SnapshotRestoreResponse)) as SnapshotRestoreResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotRestoreResponse create() => SnapshotRestoreResponse._();
  SnapshotRestoreResponse createEmptyInstance() => create();
  static $pb.PbList<SnapshotRestoreResponse> createRepeated() => $pb.PbList<SnapshotRestoreResponse>();
  @$core.pragma('dart2js:noInline')
  static SnapshotRestoreResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotRestoreResponse>(create);
  static SnapshotRestoreResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get cmdStr => $_getSZ(0);
  @$pb.TagNumber(1)
  set cmdStr($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCmdStr() => $_has(0);
  @$pb.TagNumber(1)
  void clearCmdStr() => clearField(1);
}

class SnapshotInfoParam extends $pb.GeneratedMessage {
  factory SnapshotInfoParam({
    $core.List<$core.int>? file,
  }) {
    final $result = create();
    if (file != null) {
      $result.file = file;
    }
    return $result;
  }
  SnapshotInfoParam._() : super();
  factory SnapshotInfoParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotInfoParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SnapshotInfoParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY, protoName: 'File')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotInfoParam clone() => SnapshotInfoParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotInfoParam copyWith(void Function(SnapshotInfoParam) updates) => super.copyWith((message) => updates(message as SnapshotInfoParam)) as SnapshotInfoParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotInfoParam create() => SnapshotInfoParam._();
  SnapshotInfoParam createEmptyInstance() => create();
  static $pb.PbList<SnapshotInfoParam> createRepeated() => $pb.PbList<SnapshotInfoParam>();
  @$core.pragma('dart2js:noInline')
  static SnapshotInfoParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotInfoParam>(create);
  static SnapshotInfoParam? _defaultInstance;

  @$pb.TagNumber(2)
  $core.List<$core.int> get file => $_getN(0);
  @$pb.TagNumber(2)
  set file($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(0);
  @$pb.TagNumber(2)
  void clearFile() => clearField(2);
}

class SnapshotInfoResponse extends $pb.GeneratedMessage {
  factory SnapshotInfoResponse({
    $core.String? info,
  }) {
    final $result = create();
    if (info != null) {
      $result.info = info;
    }
    return $result;
  }
  SnapshotInfoResponse._() : super();
  factory SnapshotInfoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotInfoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SnapshotInfoResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Info', protoName: 'Info')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotInfoResponse clone() => SnapshotInfoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotInfoResponse copyWith(void Function(SnapshotInfoResponse) updates) => super.copyWith((message) => updates(message as SnapshotInfoResponse)) as SnapshotInfoResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotInfoResponse create() => SnapshotInfoResponse._();
  SnapshotInfoResponse createEmptyInstance() => create();
  static $pb.PbList<SnapshotInfoResponse> createRepeated() => $pb.PbList<SnapshotInfoResponse>();
  @$core.pragma('dart2js:noInline')
  static SnapshotInfoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotInfoResponse>(create);
  static SnapshotInfoResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get info => $_getSZ(0);
  @$pb.TagNumber(1)
  set info($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasInfo() => $_has(0);
  @$pb.TagNumber(1)
  void clearInfo() => clearField(1);
}

class AllKVsBackupParam extends $pb.GeneratedMessage {
  factory AllKVsBackupParam() => create();
  AllKVsBackupParam._() : super();
  factory AllKVsBackupParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AllKVsBackupParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AllKVsBackupParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AllKVsBackupParam clone() => AllKVsBackupParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AllKVsBackupParam copyWith(void Function(AllKVsBackupParam) updates) => super.copyWith((message) => updates(message as AllKVsBackupParam)) as AllKVsBackupParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsBackupParam create() => AllKVsBackupParam._();
  AllKVsBackupParam createEmptyInstance() => create();
  static $pb.PbList<AllKVsBackupParam> createRepeated() => $pb.PbList<AllKVsBackupParam>();
  @$core.pragma('dart2js:noInline')
  static AllKVsBackupParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AllKVsBackupParam>(create);
  static AllKVsBackupParam? _defaultInstance;
}

class AllKVsBackupResponse extends $pb.GeneratedMessage {
  factory AllKVsBackupResponse({
    $core.String? name,
    $core.List<$core.int>? file,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    if (file != null) {
      $result.file = file;
    }
    return $result;
  }
  AllKVsBackupResponse._() : super();
  factory AllKVsBackupResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AllKVsBackupResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AllKVsBackupResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..a<$core.List<$core.int>>(2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY, protoName: 'File')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AllKVsBackupResponse clone() => AllKVsBackupResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AllKVsBackupResponse copyWith(void Function(AllKVsBackupResponse) updates) => super.copyWith((message) => updates(message as AllKVsBackupResponse)) as AllKVsBackupResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsBackupResponse create() => AllKVsBackupResponse._();
  AllKVsBackupResponse createEmptyInstance() => create();
  static $pb.PbList<AllKVsBackupResponse> createRepeated() => $pb.PbList<AllKVsBackupResponse>();
  @$core.pragma('dart2js:noInline')
  static AllKVsBackupResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AllKVsBackupResponse>(create);
  static AllKVsBackupResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get file => $_getN(1);
  @$pb.TagNumber(2)
  set file($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(1);
  @$pb.TagNumber(2)
  void clearFile() => clearField(2);
}

class AllKVsRestoreParam extends $pb.GeneratedMessage {
  factory AllKVsRestoreParam({
    $core.List<$core.int>? file,
  }) {
    final $result = create();
    if (file != null) {
      $result.file = file;
    }
    return $result;
  }
  AllKVsRestoreParam._() : super();
  factory AllKVsRestoreParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AllKVsRestoreParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AllKVsRestoreParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY, protoName: 'File')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AllKVsRestoreParam clone() => AllKVsRestoreParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AllKVsRestoreParam copyWith(void Function(AllKVsRestoreParam) updates) => super.copyWith((message) => updates(message as AllKVsRestoreParam)) as AllKVsRestoreParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsRestoreParam create() => AllKVsRestoreParam._();
  AllKVsRestoreParam createEmptyInstance() => create();
  static $pb.PbList<AllKVsRestoreParam> createRepeated() => $pb.PbList<AllKVsRestoreParam>();
  @$core.pragma('dart2js:noInline')
  static AllKVsRestoreParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AllKVsRestoreParam>(create);
  static AllKVsRestoreParam? _defaultInstance;

  @$pb.TagNumber(2)
  $core.List<$core.int> get file => $_getN(0);
  @$pb.TagNumber(2)
  set file($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(0);
  @$pb.TagNumber(2)
  void clearFile() => clearField(2);
}

class AllKVsRestoreResponse extends $pb.GeneratedMessage {
  factory AllKVsRestoreResponse() => create();
  AllKVsRestoreResponse._() : super();
  factory AllKVsRestoreResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AllKVsRestoreResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AllKVsRestoreResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AllKVsRestoreResponse clone() => AllKVsRestoreResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AllKVsRestoreResponse copyWith(void Function(AllKVsRestoreResponse) updates) => super.copyWith((message) => updates(message as AllKVsRestoreResponse)) as AllKVsRestoreResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsRestoreResponse create() => AllKVsRestoreResponse._();
  AllKVsRestoreResponse createEmptyInstance() => create();
  static $pb.PbList<AllKVsRestoreResponse> createRepeated() => $pb.PbList<AllKVsRestoreResponse>();
  @$core.pragma('dart2js:noInline')
  static AllKVsRestoreResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AllKVsRestoreResponse>(create);
  static AllKVsRestoreResponse? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
