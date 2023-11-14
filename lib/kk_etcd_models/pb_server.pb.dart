///
//  Generated code. Do not modify.
//  source: pb_server.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PBServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'PBServer', package: const $pb.PackageName(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names')
        ? ''
        : 'ServiceName', protoName: 'ServiceName')..aOS(2,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'ServiceAddr', protoName: 'ServiceAddr')
    ..hasRequiredFields = false
  ;

  PBServer._() : super();
  factory PBServer({
    $core.String? serviceName,
    $core.String? serviceAddr,
  }) {
    final _result = create();
    if (serviceName != null) {
      _result.serviceName = serviceName;
    }
    if (serviceAddr != null) {
      _result.serviceAddr = serviceAddr;
    }
    return _result;
  }

  factory PBServer.fromBuffer($core.List<$core.int> i,
      [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()
        ..mergeFromBuffer(i, r);

  factory PBServer.fromJson($core.String i,
      [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()
        ..mergeFromJson(i, r);
  @$core.Deprecated(
      'Using this can add significant overhead to your binary. '
          'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
          'Will be removed in next major version')
  PBServer clone() =>
      PBServer()
        ..mergeFromMessage(this);
  @$core.Deprecated(
      'Using this can add significant overhead to your binary. '
          'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
          'Will be removed in next major version')
  PBServer copyWith(void Function(PBServer) updates) =>
      super.copyWith((message) => updates(
          message as PBServer)) as PBServer; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PBServer create() => PBServer._();
  PBServer createEmptyInstance() => create();
  static $pb.PbList<PBServer> createRepeated() => $pb.PbList<PBServer>();
  @$core.pragma('dart2js:noInline')
  static PBServer getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBServer>(create);
  static PBServer? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get serviceName => $_getSZ(0);
  @$pb.TagNumber(1)
  set serviceName($core.String v) {
    $_setString(0, v);
  }
  @$pb.TagNumber(1)
  $core.bool hasServiceName() => $_has(0);
  @$pb.TagNumber(1)
  void clearServiceName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get serviceAddr => $_getSZ(1);
  @$pb.TagNumber(2)
  set serviceAddr($core.String v) {
    $_setString(1, v);
  }
  @$pb.TagNumber(2)
  $core.bool hasServiceAddr() => $_has(1);
  @$pb.TagNumber(2)
  void clearServiceAddr() => clearField(2);
}

class PBListServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'PBListServer', package: const $pb.PackageName(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'kk_etcd_models'), createEmptyInstance: create)
    ..pc<PBServer>(1,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'ListServer', $pb.PbFieldType.PM, protoName: 'ListServer',
        subBuilder: PBServer.create)
    ..hasRequiredFields = false
  ;

  PBListServer._() : super();
  factory PBListServer({
    $core.Iterable<PBServer>? listServer,
  }) {
    final _result = create();
    if (listServer != null) {
      _result.listServer.addAll(listServer);
    }
    return _result;
  }

  factory PBListServer.fromBuffer($core.List<$core.int> i,
      [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()
        ..mergeFromBuffer(i, r);

  factory PBListServer.fromJson($core.String i,
      [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()
        ..mergeFromJson(i, r);
  @$core.Deprecated(
      'Using this can add significant overhead to your binary. '
          'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
          'Will be removed in next major version')
  PBListServer clone() =>
      PBListServer()
        ..mergeFromMessage(this);
  @$core.Deprecated(
      'Using this can add significant overhead to your binary. '
          'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
          'Will be removed in next major version')
  PBListServer copyWith(void Function(PBListServer) updates) =>
      super.copyWith((message) => updates(
          message as PBListServer)) as PBListServer; // ignore: deprecated_member_use
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

