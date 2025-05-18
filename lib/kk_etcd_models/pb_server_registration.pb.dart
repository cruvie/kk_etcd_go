//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_server_registration.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../google/protobuf/duration.pb.dart' as $2;
import 'pb_server_registration.pbenum.dart';

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

export 'pb_server_registration.pbenum.dart';

class PBServerRegistration_PBCheckConfig extends $pb.GeneratedMessage {
  factory PBServerRegistration_PBCheckConfig({
    PBServerType? type,
    $2.Duration? timeout,
    $2.Duration? interval,
    $core.String? addr,
  }) {
    final $result = create();
    if (type != null) {
      $result.type = type;
    }
    if (timeout != null) {
      $result.timeout = timeout;
    }
    if (interval != null) {
      $result.interval = interval;
    }
    if (addr != null) {
      $result.addr = addr;
    }
    return $result;
  }
  PBServerRegistration_PBCheckConfig._() : super();
  factory PBServerRegistration_PBCheckConfig.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBServerRegistration_PBCheckConfig.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBServerRegistration.PBCheckConfig', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..e<PBServerType>(1, _omitFieldNames ? '' : 'Type', $pb.PbFieldType.OE, protoName: 'Type', defaultOrMaker: PBServerType.Unknown, valueOf: PBServerType.valueOf, enumValues: PBServerType.values)
    ..aOM<$2.Duration>(2, _omitFieldNames ? '' : 'Timeout', protoName: 'Timeout', subBuilder: $2.Duration.create)
    ..aOM<$2.Duration>(3, _omitFieldNames ? '' : 'Interval', protoName: 'Interval', subBuilder: $2.Duration.create)
    ..aOS(4, _omitFieldNames ? '' : 'Addr', protoName: 'Addr')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PBServerRegistration_PBCheckConfig clone() => PBServerRegistration_PBCheckConfig()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PBServerRegistration_PBCheckConfig copyWith(void Function(PBServerRegistration_PBCheckConfig) updates) => super.copyWith((message) => updates(message as PBServerRegistration_PBCheckConfig)) as PBServerRegistration_PBCheckConfig;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBServerRegistration_PBCheckConfig create() => PBServerRegistration_PBCheckConfig._();
  PBServerRegistration_PBCheckConfig createEmptyInstance() => create();
  static $pb.PbList<PBServerRegistration_PBCheckConfig> createRepeated() => $pb.PbList<PBServerRegistration_PBCheckConfig>();
  @$core.pragma('dart2js:noInline')
  static PBServerRegistration_PBCheckConfig getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBServerRegistration_PBCheckConfig>(create);
  static PBServerRegistration_PBCheckConfig? _defaultInstance;

  @$pb.TagNumber(1)
  PBServerType get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(PBServerType v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => $_clearField(1);

  /// check timeout
  /// must grater than 1 second
  @$pb.TagNumber(2)
  $2.Duration get timeout => $_getN(1);
  @$pb.TagNumber(2)
  set timeout($2.Duration v) { $_setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasTimeout() => $_has(1);
  @$pb.TagNumber(2)
  void clearTimeout() => $_clearField(2);
  @$pb.TagNumber(2)
  $2.Duration ensureTimeout() => $_ensure(1);

  /// check interval, kk_etcd will per Interval to send check request
  /// must grater than Timeout
  @$pb.TagNumber(3)
  $2.Duration get interval => $_getN(2);
  @$pb.TagNumber(3)
  set interval($2.Duration v) { $_setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasInterval() => $_has(2);
  @$pb.TagNumber(3)
  void clearInterval() => $_clearField(3);
  @$pb.TagNumber(3)
  $2.Duration ensureInterval() => $_ensure(2);

  /// Http default http://+Addr+/KKHealthCheck
  /// Grpc Addr in ServerRegistration
  @$pb.TagNumber(4)
  $core.String get addr => $_getSZ(3);
  @$pb.TagNumber(4)
  set addr($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAddr() => $_has(3);
  @$pb.TagNumber(4)
  void clearAddr() => $_clearField(4);
}

class PBServerRegistration extends $pb.GeneratedMessage {
  factory PBServerRegistration({
    PBServerType? serverType,
    $core.String? serverName,
    $core.String? serverAddr,
    PBServerRegistration_PBCheckConfig? checkConfig,
  }) {
    final $result = create();
    if (serverType != null) {
      $result.serverType = serverType;
    }
    if (serverName != null) {
      $result.serverName = serverName;
    }
    if (serverAddr != null) {
      $result.serverAddr = serverAddr;
    }
    if (checkConfig != null) {
      $result.checkConfig = checkConfig;
    }
    return $result;
  }
  PBServerRegistration._() : super();
  factory PBServerRegistration.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBServerRegistration.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBServerRegistration', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..e<PBServerType>(1, _omitFieldNames ? '' : 'ServerType', $pb.PbFieldType.OE, protoName: 'ServerType', defaultOrMaker: PBServerType.Unknown, valueOf: PBServerType.valueOf, enumValues: PBServerType.values)
    ..aOS(2, _omitFieldNames ? '' : 'ServerName', protoName: 'ServerName')
    ..aOS(3, _omitFieldNames ? '' : 'ServerAddr', protoName: 'ServerAddr')
    ..aOM<PBServerRegistration_PBCheckConfig>(4, _omitFieldNames ? '' : 'CheckConfig', protoName: 'CheckConfig', subBuilder: PBServerRegistration_PBCheckConfig.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PBServerRegistration clone() => PBServerRegistration()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PBServerRegistration copyWith(void Function(PBServerRegistration) updates) => super.copyWith((message) => updates(message as PBServerRegistration)) as PBServerRegistration;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBServerRegistration create() => PBServerRegistration._();
  PBServerRegistration createEmptyInstance() => create();
  static $pb.PbList<PBServerRegistration> createRepeated() => $pb.PbList<PBServerRegistration>();
  @$core.pragma('dart2js:noInline')
  static PBServerRegistration getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBServerRegistration>(create);
  static PBServerRegistration? _defaultInstance;

  @$pb.TagNumber(1)
  PBServerType get serverType => $_getN(0);
  @$pb.TagNumber(1)
  set serverType(PBServerType v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServerType() => $_has(0);
  @$pb.TagNumber(1)
  void clearServerType() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get serverName => $_getSZ(1);
  @$pb.TagNumber(2)
  set serverName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasServerName() => $_has(1);
  @$pb.TagNumber(2)
  void clearServerName() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get serverAddr => $_getSZ(2);
  @$pb.TagNumber(3)
  set serverAddr($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasServerAddr() => $_has(2);
  @$pb.TagNumber(3)
  void clearServerAddr() => $_clearField(3);

  @$pb.TagNumber(4)
  PBServerRegistration_PBCheckConfig get checkConfig => $_getN(3);
  @$pb.TagNumber(4)
  set checkConfig(PBServerRegistration_PBCheckConfig v) { $_setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasCheckConfig() => $_has(3);
  @$pb.TagNumber(4)
  void clearCheckConfig() => $_clearField(4);
  @$pb.TagNumber(4)
  PBServerRegistration_PBCheckConfig ensureCheckConfig() => $_ensure(3);
}

class PBListServerRegistration extends $pb.GeneratedMessage {
  factory PBListServerRegistration({
    $core.Iterable<PBServerRegistration>? list,
  }) {
    final $result = create();
    if (list != null) {
      $result.list.addAll(list);
    }
    return $result;
  }
  PBListServerRegistration._() : super();
  factory PBListServerRegistration.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBListServerRegistration.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBListServerRegistration', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..pc<PBServerRegistration>(1, _omitFieldNames ? '' : 'List', $pb.PbFieldType.PM, protoName: 'List', subBuilder: PBServerRegistration.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PBListServerRegistration clone() => PBListServerRegistration()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PBListServerRegistration copyWith(void Function(PBListServerRegistration) updates) => super.copyWith((message) => updates(message as PBListServerRegistration)) as PBListServerRegistration;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListServerRegistration create() => PBListServerRegistration._();
  PBListServerRegistration createEmptyInstance() => create();
  static $pb.PbList<PBListServerRegistration> createRepeated() => $pb.PbList<PBListServerRegistration>();
  @$core.pragma('dart2js:noInline')
  static PBListServerRegistration getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBListServerRegistration>(create);
  static PBListServerRegistration? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<PBServerRegistration> get list => $_getList(0);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
