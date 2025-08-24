// This is a generated file - do not edit.
//
// Generated from internal/service_hub/service/api_def/DeregisterService.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../../kk_etcd_models/pb_service_kk_etcd.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class DeregisterService_Input extends $pb.GeneratedMessage {
  factory DeregisterService_Input({
    $0.PBService? service,
  }) {
    final result = create();
    if (service != null) result.service = service;
    return result;
  }

  DeregisterService_Input._();

  factory DeregisterService_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory DeregisterService_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'DeregisterService.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..aOM<$0.PBService>(1, _omitFieldNames ? '' : 'Service',
        protoName: 'Service', subBuilder: $0.PBService.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService_Input clone() =>
      DeregisterService_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService_Input copyWith(
          void Function(DeregisterService_Input) updates) =>
      super.copyWith((message) => updates(message as DeregisterService_Input))
          as DeregisterService_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterService_Input create() => DeregisterService_Input._();
  @$core.override
  DeregisterService_Input createEmptyInstance() => create();
  static $pb.PbList<DeregisterService_Input> createRepeated() =>
      $pb.PbList<DeregisterService_Input>();
  @$core.pragma('dart2js:noInline')
  static DeregisterService_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<DeregisterService_Input>(create);
  static DeregisterService_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBService get service => $_getN(0);
  @$pb.TagNumber(1)
  set service($0.PBService value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasService() => $_has(0);
  @$pb.TagNumber(1)
  void clearService() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBService ensureService() => $_ensure(0);
}

class DeregisterService_Output extends $pb.GeneratedMessage {
  factory DeregisterService_Output() => create();

  DeregisterService_Output._();

  factory DeregisterService_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory DeregisterService_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'DeregisterService.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService_Output clone() =>
      DeregisterService_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService_Output copyWith(
          void Function(DeregisterService_Output) updates) =>
      super.copyWith((message) => updates(message as DeregisterService_Output))
          as DeregisterService_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterService_Output create() => DeregisterService_Output._();
  @$core.override
  DeregisterService_Output createEmptyInstance() => create();
  static $pb.PbList<DeregisterService_Output> createRepeated() =>
      $pb.PbList<DeregisterService_Output>();
  @$core.pragma('dart2js:noInline')
  static DeregisterService_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<DeregisterService_Output>(create);
  static DeregisterService_Output? _defaultInstance;
}

/// deregister server
class DeregisterService extends $pb.GeneratedMessage {
  factory DeregisterService() => create();

  DeregisterService._();

  factory DeregisterService.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory DeregisterService.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'DeregisterService',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService clone() => DeregisterService()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService copyWith(void Function(DeregisterService) updates) =>
      super.copyWith((message) => updates(message as DeregisterService))
          as DeregisterService;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterService create() => DeregisterService._();
  @$core.override
  DeregisterService createEmptyInstance() => create();
  static $pb.PbList<DeregisterService> createRepeated() =>
      $pb.PbList<DeregisterService>();
  @$core.pragma('dart2js:noInline')
  static DeregisterService getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<DeregisterService>(create);
  static DeregisterService? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
