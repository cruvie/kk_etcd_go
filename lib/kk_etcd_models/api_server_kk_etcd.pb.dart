//
//  Generated code. Do not modify.
//  source: api_server_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'pb_server_kk_etcd.pb.dart' as $3;

class ServerListParam extends $pb.GeneratedMessage {
  factory ServerListParam({
    $core.String? serverType,
    $core.String? serverName,
  }) {
    final $result = create();
    if (serverType != null) {
      $result.serverType = serverType;
    }
    if (serverName != null) {
      $result.serverName = serverName;
    }
    return $result;
  }
  ServerListParam._() : super();
  factory ServerListParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerListParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ServerListParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'ServerType', protoName: 'ServerType')
    ..aOS(2, _omitFieldNames ? '' : 'ServerName', protoName: 'ServerName')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServerListParam clone() => ServerListParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServerListParam copyWith(void Function(ServerListParam) updates) => super.copyWith((message) => updates(message as ServerListParam)) as ServerListParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ServerListParam create() => ServerListParam._();
  ServerListParam createEmptyInstance() => create();
  static $pb.PbList<ServerListParam> createRepeated() => $pb.PbList<ServerListParam>();
  @$core.pragma('dart2js:noInline')
  static ServerListParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerListParam>(create);
  static ServerListParam? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get serverType => $_getSZ(0);
  @$pb.TagNumber(1)
  set serverType($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasServerType() => $_has(0);
  @$pb.TagNumber(1)
  void clearServerType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get serverName => $_getSZ(1);
  @$pb.TagNumber(2)
  set serverName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasServerName() => $_has(1);
  @$pb.TagNumber(2)
  void clearServerName() => clearField(2);
}

class ServerListResponse extends $pb.GeneratedMessage {
  factory ServerListResponse({
    $3.PBListServer? serverList,
  }) {
    final $result = create();
    if (serverList != null) {
      $result.serverList = serverList;
    }
    return $result;
  }
  ServerListResponse._() : super();
  factory ServerListResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerListResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ServerListResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOM<$3.PBListServer>(1, _omitFieldNames ? '' : 'ServerList', protoName: 'ServerList', subBuilder: $3.PBListServer.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServerListResponse clone() => ServerListResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServerListResponse copyWith(void Function(ServerListResponse) updates) => super.copyWith((message) => updates(message as ServerListResponse)) as ServerListResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ServerListResponse create() => ServerListResponse._();
  ServerListResponse createEmptyInstance() => create();
  static $pb.PbList<ServerListResponse> createRepeated() => $pb.PbList<ServerListResponse>();
  @$core.pragma('dart2js:noInline')
  static ServerListResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerListResponse>(create);
  static ServerListResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $3.PBListServer get serverList => $_getN(0);
  @$pb.TagNumber(1)
  set serverList($3.PBListServer v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServerList() => $_has(0);
  @$pb.TagNumber(1)
  void clearServerList() => clearField(1);
  @$pb.TagNumber(1)
  $3.PBListServer ensureServerList() => $_ensure(0);
}

class DeregisterServerParam extends $pb.GeneratedMessage {
  factory DeregisterServerParam({
    $3.PBServer? server,
  }) {
    final $result = create();
    if (server != null) {
      $result.server = server;
    }
    return $result;
  }
  DeregisterServerParam._() : super();
  factory DeregisterServerParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeregisterServerParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeregisterServerParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOM<$3.PBServer>(1, _omitFieldNames ? '' : 'Server', protoName: 'Server', subBuilder: $3.PBServer.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeregisterServerParam clone() => DeregisterServerParam()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeregisterServerParam copyWith(void Function(DeregisterServerParam) updates) => super.copyWith((message) => updates(message as DeregisterServerParam)) as DeregisterServerParam;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterServerParam create() => DeregisterServerParam._();
  DeregisterServerParam createEmptyInstance() => create();
  static $pb.PbList<DeregisterServerParam> createRepeated() => $pb.PbList<DeregisterServerParam>();
  @$core.pragma('dart2js:noInline')
  static DeregisterServerParam getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeregisterServerParam>(create);
  static DeregisterServerParam? _defaultInstance;

  @$pb.TagNumber(1)
  $3.PBServer get server => $_getN(0);
  @$pb.TagNumber(1)
  set server($3.PBServer v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServer() => $_has(0);
  @$pb.TagNumber(1)
  void clearServer() => clearField(1);
  @$pb.TagNumber(1)
  $3.PBServer ensureServer() => $_ensure(0);
}

class DeregisterServerResponse extends $pb.GeneratedMessage {
  factory DeregisterServerResponse() => create();
  DeregisterServerResponse._() : super();
  factory DeregisterServerResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeregisterServerResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeregisterServerResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeregisterServerResponse clone() => DeregisterServerResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeregisterServerResponse copyWith(void Function(DeregisterServerResponse) updates) => super.copyWith((message) => updates(message as DeregisterServerResponse)) as DeregisterServerResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterServerResponse create() => DeregisterServerResponse._();
  DeregisterServerResponse createEmptyInstance() => create();
  static $pb.PbList<DeregisterServerResponse> createRepeated() => $pb.PbList<DeregisterServerResponse>();
  @$core.pragma('dart2js:noInline')
  static DeregisterServerResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeregisterServerResponse>(create);
  static DeregisterServerResponse? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
