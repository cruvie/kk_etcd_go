//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_server_kk_etcd.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'pb_server_kk_etcd.pbenum.dart';
import 'pb_server_registration.pb.dart' as $1;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

export 'pb_server_kk_etcd.pbenum.dart';

class PBServer extends $pb.GeneratedMessage {
  factory PBServer({
    $1.PBServerRegistration? serverRegistration,
    PBServer_ServerStatus? status,
  }) {
    final $result = create();
    if (serverRegistration != null) {
      $result.serverRegistration = serverRegistration;
    }
    if (status != null) {
      $result.status = status;
    }
    return $result;
  }
  PBServer._() : super();
  factory PBServer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBServer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBServer', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOM<$1.PBServerRegistration>(1, _omitFieldNames ? '' : 'ServerRegistration', protoName: 'ServerRegistration', subBuilder: $1.PBServerRegistration.create)
    ..e<PBServer_ServerStatus>(5, _omitFieldNames ? '' : 'Status', $pb.PbFieldType.OE, protoName: 'Status', defaultOrMaker: PBServer_ServerStatus.UnKnown, valueOf: PBServer_ServerStatus.valueOf, enumValues: PBServer_ServerStatus.values)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PBServer clone() => PBServer()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PBServer copyWith(void Function(PBServer) updates) => super.copyWith((message) => updates(message as PBServer)) as PBServer;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBServer create() => PBServer._();
  PBServer createEmptyInstance() => create();
  static $pb.PbList<PBServer> createRepeated() => $pb.PbList<PBServer>();
  @$core.pragma('dart2js:noInline')
  static PBServer getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBServer>(create);
  static PBServer? _defaultInstance;

  @$pb.TagNumber(1)
  $1.PBServerRegistration get serverRegistration => $_getN(0);
  @$pb.TagNumber(1)
  set serverRegistration($1.PBServerRegistration v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServerRegistration() => $_has(0);
  @$pb.TagNumber(1)
  void clearServerRegistration() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.PBServerRegistration ensureServerRegistration() => $_ensure(0);

  @$pb.TagNumber(5)
  PBServer_ServerStatus get status => $_getN(1);
  @$pb.TagNumber(5)
  set status(PBServer_ServerStatus v) { $_setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasStatus() => $_has(1);
  @$pb.TagNumber(5)
  void clearStatus() => $_clearField(5);
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
  factory PBListServer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBListServer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBListServer', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..pc<PBServer>(1, _omitFieldNames ? '' : 'ListServer', $pb.PbFieldType.PM, protoName: 'ListServer', subBuilder: PBServer.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PBListServer clone() => PBListServer()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PBListServer copyWith(void Function(PBListServer) updates) => super.copyWith((message) => updates(message as PBListServer)) as PBListServer;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListServer create() => PBListServer._();
  PBListServer createEmptyInstance() => create();
  static $pb.PbList<PBListServer> createRepeated() => $pb.PbList<PBListServer>();
  @$core.pragma('dart2js:noInline')
  static PBListServer getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBListServer>(create);
  static PBListServer? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<PBServer> get listServer => $_getList(0);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
