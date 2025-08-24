// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_service_registration.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../google/protobuf/duration.pb.dart' as $0;
import 'pb_service_registration.pbenum.dart';

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

export 'pb_service_registration.pbenum.dart';

class PBServiceRegistration_PBCheckConfig extends $pb.GeneratedMessage {
  factory PBServiceRegistration_PBCheckConfig({
    PBServiceType? type,
    $0.Duration? timeout,
    $0.Duration? interval,
    $core.String? addr,
  }) {
    final result = create();
    if (type != null) result.type = type;
    if (timeout != null) result.timeout = timeout;
    if (interval != null) result.interval = interval;
    if (addr != null) result.addr = addr;
    return result;
  }

  PBServiceRegistration_PBCheckConfig._();

  factory PBServiceRegistration_PBCheckConfig.fromBuffer(
          $core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBServiceRegistration_PBCheckConfig.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBServiceRegistration.PBCheckConfig',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..e<PBServiceType>(1, _omitFieldNames ? '' : 'Type', $pb.PbFieldType.OE,
        protoName: 'Type',
        defaultOrMaker: PBServiceType.Unknown,
        valueOf: PBServiceType.valueOf,
        enumValues: PBServiceType.values)
    ..aOM<$0.Duration>(2, _omitFieldNames ? '' : 'Timeout',
        protoName: 'Timeout', subBuilder: $0.Duration.create)
    ..aOM<$0.Duration>(3, _omitFieldNames ? '' : 'Interval',
        protoName: 'Interval', subBuilder: $0.Duration.create)
    ..aOS(4, _omitFieldNames ? '' : 'Addr', protoName: 'Addr')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBServiceRegistration_PBCheckConfig clone() =>
      PBServiceRegistration_PBCheckConfig()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBServiceRegistration_PBCheckConfig copyWith(
          void Function(PBServiceRegistration_PBCheckConfig) updates) =>
      super.copyWith((message) =>
              updates(message as PBServiceRegistration_PBCheckConfig))
          as PBServiceRegistration_PBCheckConfig;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBServiceRegistration_PBCheckConfig create() =>
      PBServiceRegistration_PBCheckConfig._();
  @$core.override
  PBServiceRegistration_PBCheckConfig createEmptyInstance() => create();
  static $pb.PbList<PBServiceRegistration_PBCheckConfig> createRepeated() =>
      $pb.PbList<PBServiceRegistration_PBCheckConfig>();
  @$core.pragma('dart2js:noInline')
  static PBServiceRegistration_PBCheckConfig getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<
          PBServiceRegistration_PBCheckConfig>(create);
  static PBServiceRegistration_PBCheckConfig? _defaultInstance;

  @$pb.TagNumber(1)
  PBServiceType get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(PBServiceType value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => $_clearField(1);

  /// check timeout
  /// must grater than 1 second
  @$pb.TagNumber(2)
  $0.Duration get timeout => $_getN(1);
  @$pb.TagNumber(2)
  set timeout($0.Duration value) => $_setField(2, value);
  @$pb.TagNumber(2)
  $core.bool hasTimeout() => $_has(1);
  @$pb.TagNumber(2)
  void clearTimeout() => $_clearField(2);
  @$pb.TagNumber(2)
  $0.Duration ensureTimeout() => $_ensure(1);

  /// check interval, kk_etcd will per Interval to send check request
  /// must grater than Timeout
  @$pb.TagNumber(3)
  $0.Duration get interval => $_getN(2);
  @$pb.TagNumber(3)
  set interval($0.Duration value) => $_setField(3, value);
  @$pb.TagNumber(3)
  $core.bool hasInterval() => $_has(2);
  @$pb.TagNumber(3)
  void clearInterval() => $_clearField(3);
  @$pb.TagNumber(3)
  $0.Duration ensureInterval() => $_ensure(2);

  /// Http default http://+Addr+/KKHealthCheck
  /// Grpc Addr in ServiceRegistration
  @$pb.TagNumber(4)
  $core.String get addr => $_getSZ(3);
  @$pb.TagNumber(4)
  set addr($core.String value) => $_setString(3, value);
  @$pb.TagNumber(4)
  $core.bool hasAddr() => $_has(3);
  @$pb.TagNumber(4)
  void clearAddr() => $_clearField(4);
}

class PBServiceRegistration extends $pb.GeneratedMessage {
  factory PBServiceRegistration({
    PBServiceType? serviceType,
    $core.String? serviceName,
    $core.String? serviceAddr,
    PBServiceRegistration_PBCheckConfig? checkConfig,
  }) {
    final result = create();
    if (serviceType != null) result.serviceType = serviceType;
    if (serviceName != null) result.serviceName = serviceName;
    if (serviceAddr != null) result.serviceAddr = serviceAddr;
    if (checkConfig != null) result.checkConfig = checkConfig;
    return result;
  }

  PBServiceRegistration._();

  factory PBServiceRegistration.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBServiceRegistration.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBServiceRegistration',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..e<PBServiceType>(
        1, _omitFieldNames ? '' : 'ServiceType', $pb.PbFieldType.OE,
        protoName: 'ServiceType',
        defaultOrMaker: PBServiceType.Unknown,
        valueOf: PBServiceType.valueOf,
        enumValues: PBServiceType.values)
    ..aOS(2, _omitFieldNames ? '' : 'ServiceName', protoName: 'ServiceName')
    ..aOS(3, _omitFieldNames ? '' : 'ServiceAddr', protoName: 'ServiceAddr')
    ..aOM<PBServiceRegistration_PBCheckConfig>(
        4, _omitFieldNames ? '' : 'CheckConfig',
        protoName: 'CheckConfig',
        subBuilder: PBServiceRegistration_PBCheckConfig.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBServiceRegistration clone() =>
      PBServiceRegistration()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBServiceRegistration copyWith(
          void Function(PBServiceRegistration) updates) =>
      super.copyWith((message) => updates(message as PBServiceRegistration))
          as PBServiceRegistration;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBServiceRegistration create() => PBServiceRegistration._();
  @$core.override
  PBServiceRegistration createEmptyInstance() => create();
  static $pb.PbList<PBServiceRegistration> createRepeated() =>
      $pb.PbList<PBServiceRegistration>();
  @$core.pragma('dart2js:noInline')
  static PBServiceRegistration getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<PBServiceRegistration>(create);
  static PBServiceRegistration? _defaultInstance;

  @$pb.TagNumber(1)
  PBServiceType get serviceType => $_getN(0);
  @$pb.TagNumber(1)
  set serviceType(PBServiceType value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasServiceType() => $_has(0);
  @$pb.TagNumber(1)
  void clearServiceType() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get serviceName => $_getSZ(1);
  @$pb.TagNumber(2)
  set serviceName($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasServiceName() => $_has(1);
  @$pb.TagNumber(2)
  void clearServiceName() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get serviceAddr => $_getSZ(2);
  @$pb.TagNumber(3)
  set serviceAddr($core.String value) => $_setString(2, value);
  @$pb.TagNumber(3)
  $core.bool hasServiceAddr() => $_has(2);
  @$pb.TagNumber(3)
  void clearServiceAddr() => $_clearField(3);

  @$pb.TagNumber(4)
  PBServiceRegistration_PBCheckConfig get checkConfig => $_getN(3);
  @$pb.TagNumber(4)
  set checkConfig(PBServiceRegistration_PBCheckConfig value) =>
      $_setField(4, value);
  @$pb.TagNumber(4)
  $core.bool hasCheckConfig() => $_has(3);
  @$pb.TagNumber(4)
  void clearCheckConfig() => $_clearField(4);
  @$pb.TagNumber(4)
  PBServiceRegistration_PBCheckConfig ensureCheckConfig() => $_ensure(3);
}

class PBListServiceRegistration extends $pb.GeneratedMessage {
  factory PBListServiceRegistration({
    $core.Iterable<PBServiceRegistration>? list,
  }) {
    final result = create();
    if (list != null) result.list.addAll(list);
    return result;
  }

  PBListServiceRegistration._();

  factory PBListServiceRegistration.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBListServiceRegistration.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBListServiceRegistration',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..pc<PBServiceRegistration>(
        1, _omitFieldNames ? '' : 'List', $pb.PbFieldType.PM,
        protoName: 'List', subBuilder: PBServiceRegistration.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListServiceRegistration clone() =>
      PBListServiceRegistration()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListServiceRegistration copyWith(
          void Function(PBListServiceRegistration) updates) =>
      super.copyWith((message) => updates(message as PBListServiceRegistration))
          as PBListServiceRegistration;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListServiceRegistration create() => PBListServiceRegistration._();
  @$core.override
  PBListServiceRegistration createEmptyInstance() => create();
  static $pb.PbList<PBListServiceRegistration> createRepeated() =>
      $pb.PbList<PBListServiceRegistration>();
  @$core.pragma('dart2js:noInline')
  static PBListServiceRegistration getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<PBListServiceRegistration>(create);
  static PBListServiceRegistration? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<PBServiceRegistration> get list => $_getList(0);
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
