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

import 'pb_server_kk_etcd.pb.dart' as $2;

class ServerListParam extends $pb.GeneratedMessage {
  factory ServerListParam({
    $core.String? prefix,
  }) {
    final $result = create();
    if (prefix != null) {
      $result.prefix = prefix;
    }
    return $result;
  }
  ServerListParam._() : super();
  factory ServerListParam.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerListParam.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ServerListParam', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Prefix', protoName: 'Prefix')
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
  $core.String get prefix => $_getSZ(0);
  @$pb.TagNumber(1)
  set prefix($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasPrefix() => $_has(0);
  @$pb.TagNumber(1)
  void clearPrefix() => clearField(1);
}

class ServerListResponse extends $pb.GeneratedMessage {
  factory ServerListResponse({
    $2.PBListServer? serverList,
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
    ..aOM<$2.PBListServer>(1, _omitFieldNames ? '' : 'ServerList', protoName: 'ServerList', subBuilder: $2.PBListServer.create)
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
  $2.PBListServer get serverList => $_getN(0);
  @$pb.TagNumber(1)
  set serverList($2.PBListServer v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServerList() => $_has(0);
  @$pb.TagNumber(1)
  void clearServerList() => clearField(1);
  @$pb.TagNumber(1)
  $2.PBListServer ensureServerList() => $_ensure(0);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
