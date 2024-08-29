//
//  Generated code. Do not modify.
//  source: api_role_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'pb_role_kk_etcd.pb.dart' as $1;

class RoleAddParam extends $pb.GeneratedMessage {
  factory RoleAddParam({
    $core.String? name,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  RoleAddParam._() : super();
  factory RoleAddParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleAddParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleAddParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleAddParam clone() => RoleAddParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleAddParam copyWith(void Function(RoleAddParam) updates) => super.copyWith((message) => updates(message as RoleAddParam)) as RoleAddParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleAddParam create() => RoleAddParam._();
  RoleAddParam createEmptyInstance() => create();
  static $pb.PbList<RoleAddParam> createRepeated() => $pb.PbList<RoleAddParam>();
  @$core.pragma('dart2js:noInline')
  static RoleAddParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleAddParam>(create);
  static RoleAddParam? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);
}

class RoleAddResponse extends $pb.GeneratedMessage {
  factory RoleAddResponse() => create();
  RoleAddResponse._() : super();
  factory RoleAddResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleAddResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleAddResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleAddResponse clone() => RoleAddResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleAddResponse copyWith(void Function(RoleAddResponse) updates) => super.copyWith((message) => updates(message as RoleAddResponse)) as RoleAddResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleAddResponse create() => RoleAddResponse._();
  RoleAddResponse createEmptyInstance() => create();
  static $pb.PbList<RoleAddResponse> createRepeated() => $pb.PbList<RoleAddResponse>();
  @$core.pragma('dart2js:noInline')
  static RoleAddResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleAddResponse>(create);
  static RoleAddResponse? _defaultInstance;
}

class RoleDeleteParam extends $pb.GeneratedMessage {
  factory RoleDeleteParam({
    $core.String? name,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  RoleDeleteParam._() : super();
  factory RoleDeleteParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleDeleteParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleDeleteParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleDeleteParam clone() => RoleDeleteParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleDeleteParam copyWith(void Function(RoleDeleteParam) updates) => super.copyWith((message) => updates(message as RoleDeleteParam)) as RoleDeleteParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleDeleteParam create() => RoleDeleteParam._();
  RoleDeleteParam createEmptyInstance() => create();
  static $pb.PbList<RoleDeleteParam> createRepeated() => $pb.PbList<RoleDeleteParam>();
  @$core.pragma('dart2js:noInline')
  static RoleDeleteParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleDeleteParam>(create);
  static RoleDeleteParam? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);
}

class RoleDeleteResponse extends $pb.GeneratedMessage {
  factory RoleDeleteResponse() => create();
  RoleDeleteResponse._() : super();
  factory RoleDeleteResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleDeleteResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleDeleteResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleDeleteResponse clone() => RoleDeleteResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleDeleteResponse copyWith(void Function(RoleDeleteResponse) updates) => super.copyWith((message) => updates(message as RoleDeleteResponse)) as RoleDeleteResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleDeleteResponse create() => RoleDeleteResponse._();
  RoleDeleteResponse createEmptyInstance() => create();
  static $pb.PbList<RoleDeleteResponse> createRepeated() => $pb.PbList<RoleDeleteResponse>();
  @$core.pragma('dart2js:noInline')
  static RoleDeleteResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleDeleteResponse>(create);
  static RoleDeleteResponse? _defaultInstance;
}

class RoleGetParam extends $pb.GeneratedMessage {
  factory RoleGetParam({
    $core.String? name,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  RoleGetParam._() : super();
  factory RoleGetParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGetParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGetParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGetParam clone() => RoleGetParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGetParam copyWith(void Function(RoleGetParam) updates) => super.copyWith((message) => updates(message as RoleGetParam)) as RoleGetParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGetParam create() => RoleGetParam._();
  RoleGetParam createEmptyInstance() => create();
  static $pb.PbList<RoleGetParam> createRepeated() => $pb.PbList<RoleGetParam>();
  @$core.pragma('dart2js:noInline')
  static RoleGetParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGetParam>(create);
  static RoleGetParam? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);
}

class RoleGetResponse extends $pb.GeneratedMessage {
  factory RoleGetResponse({
    $1.PBRole? role,
  }) {
    final $result = create();
    if (role != null) {
      $result.role = role;
    }
    return $result;
  }
  RoleGetResponse._() : super();
  factory RoleGetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGetResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOM<$1.PBRole>(1, _omitFieldNames ? '' : 'Role', protoName: 'Role', subBuilder: $1.PBRole.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGetResponse clone() => RoleGetResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGetResponse copyWith(void Function(RoleGetResponse) updates) => super.copyWith((message) => updates(message as RoleGetResponse)) as RoleGetResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGetResponse create() => RoleGetResponse._();
  RoleGetResponse createEmptyInstance() => create();
  static $pb.PbList<RoleGetResponse> createRepeated() => $pb.PbList<RoleGetResponse>();
  @$core.pragma('dart2js:noInline')
  static RoleGetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGetResponse>(create);
  static RoleGetResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $1.PBRole get role => $_getN(0);
  @$pb.TagNumber(1)
  set role($1.PBRole v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasRole() => $_has(0);
  @$pb.TagNumber(1)
  void clearRole() => clearField(1);
  @$pb.TagNumber(1)
  $1.PBRole ensureRole() => $_ensure(0);
}

class RoleListParam extends $pb.GeneratedMessage {
  factory RoleListParam() => create();
  RoleListParam._() : super();
  factory RoleListParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleListParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleListParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleListParam clone() => RoleListParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleListParam copyWith(void Function(RoleListParam) updates) => super.copyWith((message) => updates(message as RoleListParam)) as RoleListParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleListParam create() => RoleListParam._();
  RoleListParam createEmptyInstance() => create();
  static $pb.PbList<RoleListParam> createRepeated() => $pb.PbList<RoleListParam>();
  @$core.pragma('dart2js:noInline')
  static RoleListParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleListParam>(create);
  static RoleListParam? _defaultInstance;
}

class RoleListResponse extends $pb.GeneratedMessage {
  factory RoleListResponse({
    $1.PBListRole? listRole,
  }) {
    final $result = create();
    if (listRole != null) {
      $result.listRole = listRole;
    }
    return $result;
  }
  RoleListResponse._() : super();
  factory RoleListResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleListResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleListResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOM<$1.PBListRole>(1, _omitFieldNames ? '' : 'ListRole', protoName: 'ListRole', subBuilder: $1.PBListRole.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleListResponse clone() => RoleListResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleListResponse copyWith(void Function(RoleListResponse) updates) => super.copyWith((message) => updates(message as RoleListResponse)) as RoleListResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleListResponse create() => RoleListResponse._();
  RoleListResponse createEmptyInstance() => create();
  static $pb.PbList<RoleListResponse> createRepeated() => $pb.PbList<RoleListResponse>();
  @$core.pragma('dart2js:noInline')
  static RoleListResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleListResponse>(create);
  static RoleListResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $1.PBListRole get listRole => $_getN(0);
  @$pb.TagNumber(1)
  set listRole($1.PBListRole v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasListRole() => $_has(0);
  @$pb.TagNumber(1)
  void clearListRole() => clearField(1);
  @$pb.TagNumber(1)
  $1.PBListRole ensureListRole() => $_ensure(0);
}

class RoleGrantPermissionParam extends $pb.GeneratedMessage {
  factory RoleGrantPermissionParam({
    $core.String? name,
    $1.PBPermission? perm,
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
  RoleGrantPermissionParam._() : super();
  factory RoleGrantPermissionParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGrantPermissionParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGrantPermissionParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..aOM<$1.PBPermission>(2, _omitFieldNames ? '' : 'Perm', protoName: 'Perm', subBuilder: $1.PBPermission.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGrantPermissionParam clone() => RoleGrantPermissionParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGrantPermissionParam copyWith(void Function(RoleGrantPermissionParam) updates) => super.copyWith((message) => updates(message as RoleGrantPermissionParam)) as RoleGrantPermissionParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGrantPermissionParam create() => RoleGrantPermissionParam._();
  RoleGrantPermissionParam createEmptyInstance() => create();
  static $pb.PbList<RoleGrantPermissionParam> createRepeated() => $pb.PbList<RoleGrantPermissionParam>();
  @$core.pragma('dart2js:noInline')
  static RoleGrantPermissionParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGrantPermissionParam>(create);
  static RoleGrantPermissionParam? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $1.PBPermission get perm => $_getN(1);
  @$pb.TagNumber(2)
  set perm($1.PBPermission v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasPerm() => $_has(1);
  @$pb.TagNumber(2)
  void clearPerm() => clearField(2);
  @$pb.TagNumber(2)
  $1.PBPermission ensurePerm() => $_ensure(1);
}

class RoleGrantPermissionResponse extends $pb.GeneratedMessage {
  factory RoleGrantPermissionResponse() => create();
  RoleGrantPermissionResponse._() : super();
  factory RoleGrantPermissionResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleGrantPermissionResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleGrantPermissionResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleGrantPermissionResponse clone() => RoleGrantPermissionResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleGrantPermissionResponse copyWith(void Function(RoleGrantPermissionResponse) updates) => super.copyWith((message) => updates(message as RoleGrantPermissionResponse)) as RoleGrantPermissionResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGrantPermissionResponse create() => RoleGrantPermissionResponse._();
  RoleGrantPermissionResponse createEmptyInstance() => create();
  static $pb.PbList<RoleGrantPermissionResponse> createRepeated() => $pb.PbList<RoleGrantPermissionResponse>();
  @$core.pragma('dart2js:noInline')
  static RoleGrantPermissionResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGrantPermissionResponse>(create);
  static RoleGrantPermissionResponse? _defaultInstance;
}

class RoleRevokePermissionParam extends $pb.GeneratedMessage {
  factory RoleRevokePermissionParam({
    $core.String? name,
    $core.String? key,
    $core.String? rangeEnd,
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
    return $result;
  }
  RoleRevokePermissionParam._() : super();
  factory RoleRevokePermissionParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleRevokePermissionParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleRevokePermissionParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..aOS(2, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..aOS(3, _omitFieldNames ? '' : 'RangeEnd', protoName: 'RangeEnd')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleRevokePermissionParam clone() => RoleRevokePermissionParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleRevokePermissionParam copyWith(void Function(RoleRevokePermissionParam) updates) => super.copyWith((message) => updates(message as RoleRevokePermissionParam)) as RoleRevokePermissionParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleRevokePermissionParam create() => RoleRevokePermissionParam._();
  RoleRevokePermissionParam createEmptyInstance() => create();
  static $pb.PbList<RoleRevokePermissionParam> createRepeated() => $pb.PbList<RoleRevokePermissionParam>();
  @$core.pragma('dart2js:noInline')
  static RoleRevokePermissionParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleRevokePermissionParam>(create);
  static RoleRevokePermissionParam? _defaultInstance;

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
}

class RoleRevokePermissionResponse extends $pb.GeneratedMessage {
  factory RoleRevokePermissionResponse() => create();
  RoleRevokePermissionResponse._() : super();
  factory RoleRevokePermissionResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleRevokePermissionResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleRevokePermissionResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleRevokePermissionResponse clone() => RoleRevokePermissionResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleRevokePermissionResponse copyWith(void Function(RoleRevokePermissionResponse) updates) => super.copyWith((message) => updates(message as RoleRevokePermissionResponse)) as RoleRevokePermissionResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleRevokePermissionResponse create() => RoleRevokePermissionResponse._();
  RoleRevokePermissionResponse createEmptyInstance() => create();
  static $pb.PbList<RoleRevokePermissionResponse> createRepeated() => $pb.PbList<RoleRevokePermissionResponse>();
  @$core.pragma('dart2js:noInline')
  static RoleRevokePermissionResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleRevokePermissionResponse>(create);
  static RoleRevokePermissionResponse? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
