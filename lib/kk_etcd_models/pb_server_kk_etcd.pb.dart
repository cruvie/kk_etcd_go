//
//  Generated code. Do not modify.
//  source: pb_server_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PBServer extends $pb.GeneratedMessage {
  factory PBServer({
    $core.String? serviceName,
    $core.String? serviceAddr,
  }) {
    final $result = create();
    if (serviceName != null) {
      $result.serviceName = serviceName;
    }
    if (serviceAddr != null) {
      $result.serviceAddr = serviceAddr;
    }
    return $result;
  }
  PBServer._() : super();
  factory PBServer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBServer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBServer', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'ServiceName', protoName: 'ServiceName')
    ..aOS(2, _omitFieldNames ? '' : 'ServiceAddr', protoName: 'ServiceAddr')
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
  $core.String get serviceName => $_getSZ(0);
  @$pb.TagNumber(1)
  set serviceName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasServiceName() => $_has(0);
  @$pb.TagNumber(1)
  void clearServiceName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get serviceAddr => $_getSZ(1);
  @$pb.TagNumber(2)
  set serviceAddr($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasServiceAddr() => $_has(1);
  @$pb.TagNumber(2)
  void clearServiceAddr() => clearField(2);
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
  $core.List<PBServer> get listServer => $_getList(0);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
