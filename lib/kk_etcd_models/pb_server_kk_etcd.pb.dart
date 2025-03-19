//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_server_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../google/protobuf/timestamp.pb.dart' as $1;
import 'pb_server_kk_etcd.pbenum.dart';

export 'pb_server_kk_etcd.pbenum.dart';

class PBServer extends $pb.GeneratedMessage {
  factory PBServer({
    $core.String? serverType,
    $core.String? endpointKey,
    $core.String? endpointAddr,
    PBServer_ServerStatus? status,
    $1.Timestamp? lastCheck,
    $core.String? msg,
  }) {
    final $result = create();
    if (serverType != null) {
      $result.serverType = serverType;
    }
    if (endpointKey != null) {
      $result.endpointKey = endpointKey;
    }
    if (endpointAddr != null) {
      $result.endpointAddr = endpointAddr;
    }
    if (status != null) {
      $result.status = status;
    }
    if (lastCheck != null) {
      $result.lastCheck = lastCheck;
    }
    if (msg != null) {
      $result.msg = msg;
    }
    return $result;
  }
  PBServer._() : super();

  factory PBServer.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory PBServer.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBServer',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'),
      createEmptyInstance: create)
    ..aOS(2, _omitFieldNames ? '' : 'ServerType', protoName: 'ServerType')
    ..aOS(3, _omitFieldNames ? '' : 'EndpointKey', protoName: 'EndpointKey')
    ..aOS(4, _omitFieldNames ? '' : 'EndpointAddr', protoName: 'EndpointAddr')
    ..e<PBServer_ServerStatus>(
        5, _omitFieldNames ? '' : 'Status', $pb.PbFieldType.OE,
        protoName: 'Status',
        defaultOrMaker: PBServer_ServerStatus.UnKnown,
        valueOf: PBServer_ServerStatus.valueOf,
        enumValues: PBServer_ServerStatus.values)
    ..aOM<$1.Timestamp>(6, _omitFieldNames ? '' : 'LastCheck',
        protoName: 'LastCheck', subBuilder: $1.Timestamp.create)
    ..aOS(7, _omitFieldNames ? '' : 'Msg', protoName: 'Msg')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  PBServer clone() => PBServer()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  PBServer copyWith(void Function(PBServer) updates) =>
      super.copyWith((message) => updates(message as PBServer)) as PBServer;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBServer create() => PBServer._();
  PBServer createEmptyInstance() => create();
  static $pb.PbList<PBServer> createRepeated() => $pb.PbList<PBServer>();
  @$core.pragma('dart2js:noInline')
  static PBServer getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBServer>(create);
  static PBServer? _defaultInstance;

  @$pb.TagNumber(2)
  $core.String get serverType => $_getSZ(0);
  @$pb.TagNumber(2)
  set serverType($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasServerType() => $_has(0);
  @$pb.TagNumber(2)
  void clearServerType() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get endpointKey => $_getSZ(1);
  @$pb.TagNumber(3)
  set endpointKey($core.String v) {
    $_setString(1, v);
  }

  @$pb.TagNumber(3)
  $core.bool hasEndpointKey() => $_has(1);
  @$pb.TagNumber(3)
  void clearEndpointKey() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get endpointAddr => $_getSZ(2);
  @$pb.TagNumber(4)
  set endpointAddr($core.String v) {
    $_setString(2, v);
  }

  @$pb.TagNumber(4)
  $core.bool hasEndpointAddr() => $_has(2);
  @$pb.TagNumber(4)
  void clearEndpointAddr() => clearField(4);

  @$pb.TagNumber(5)
  PBServer_ServerStatus get status => $_getN(3);
  @$pb.TagNumber(5)
  set status(PBServer_ServerStatus v) {
    setField(5, v);
  }

  @$pb.TagNumber(5)
  $core.bool hasStatus() => $_has(3);
  @$pb.TagNumber(5)
  void clearStatus() => clearField(5);

  @$pb.TagNumber(6)
  $1.Timestamp get lastCheck => $_getN(4);
  @$pb.TagNumber(6)
  set lastCheck($1.Timestamp v) {
    setField(6, v);
  }

  @$pb.TagNumber(6)
  $core.bool hasLastCheck() => $_has(4);
  @$pb.TagNumber(6)
  void clearLastCheck() => clearField(6);
  @$pb.TagNumber(6)
  $1.Timestamp ensureLastCheck() => $_ensure(4);

  @$pb.TagNumber(7)
  $core.String get msg => $_getSZ(5);
  @$pb.TagNumber(7)
  set msg($core.String v) {
    $_setString(5, v);
  }

  @$pb.TagNumber(7)
  $core.bool hasMsg() => $_has(5);
  @$pb.TagNumber(7)
  void clearMsg() => clearField(7);
}

class PBListServer extends $pb.GeneratedMessage {
  factory PBListServer({
    $core.Iterable<PBServer>? listServer,
  }) {
    final $result = create();
    if (listServer != null) {
      $result.listServer.addAll(listServer);
    }
    return $result;
  }
  PBListServer._() : super();

  factory PBListServer.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory PBListServer.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBListServer',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'),
      createEmptyInstance: create)
    ..pc<PBServer>(1, _omitFieldNames ? '' : 'ListServer', $pb.PbFieldType.PM,
        protoName: 'ListServer', subBuilder: PBServer.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  PBListServer clone() => PBListServer()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  PBListServer copyWith(void Function(PBListServer) updates) =>
      super.copyWith((message) => updates(message as PBListServer))
          as PBListServer;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListServer create() => PBListServer._();
  PBListServer createEmptyInstance() => create();

  static $pb.PbList<PBListServer> createRepeated() =>
      $pb.PbList<PBListServer>();
  @$core.pragma('dart2js:noInline')
  static PBListServer getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<PBListServer>(create);
  static PBListServer? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<PBServer> get listServer => $_getList(0);
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
