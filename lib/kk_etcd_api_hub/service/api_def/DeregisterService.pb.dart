//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/service/api_def/DeregisterService.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../kk_etcd_models/pb_service_kk_etcd.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class DeregisterService_Input extends $pb.GeneratedMessage {
  factory DeregisterService_Input({
    $0.PBService? service,
  }) {
    final $result = create();
    if (service != null) {
      $result.service = service;
    }
    return $result;
  }
  DeregisterService_Input._() : super();
  factory DeregisterService_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeregisterService_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeregisterService.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOM<$0.PBService>(1, _omitFieldNames ? '' : 'Service', protoName: 'Service', subBuilder: $0.PBService.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService_Input clone() => DeregisterService_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService_Input copyWith(void Function(DeregisterService_Input) updates) => super.copyWith((message) => updates(message as DeregisterService_Input)) as DeregisterService_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterService_Input create() => DeregisterService_Input._();
  DeregisterService_Input createEmptyInstance() => create();
  static $pb.PbList<DeregisterService_Input> createRepeated() => $pb.PbList<DeregisterService_Input>();
  @$core.pragma('dart2js:noInline')
  static DeregisterService_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeregisterService_Input>(create);
  static DeregisterService_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBService get service => $_getN(0);
  @$pb.TagNumber(1)
  set service($0.PBService v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasService() => $_has(0);
  @$pb.TagNumber(1)
  void clearService() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBService ensureService() => $_ensure(0);
}

class DeregisterService_Output extends $pb.GeneratedMessage {
  factory DeregisterService_Output() => create();
  DeregisterService_Output._() : super();
  factory DeregisterService_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeregisterService_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeregisterService.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService_Output clone() => DeregisterService_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService_Output copyWith(void Function(DeregisterService_Output) updates) => super.copyWith((message) => updates(message as DeregisterService_Output)) as DeregisterService_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterService_Output create() => DeregisterService_Output._();
  DeregisterService_Output createEmptyInstance() => create();
  static $pb.PbList<DeregisterService_Output> createRepeated() => $pb.PbList<DeregisterService_Output>();
  @$core.pragma('dart2js:noInline')
  static DeregisterService_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeregisterService_Output>(create);
  static DeregisterService_Output? _defaultInstance;
}

/// deregister server
class DeregisterService extends $pb.GeneratedMessage {
  factory DeregisterService() => create();
  DeregisterService._() : super();
  factory DeregisterService.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeregisterService.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeregisterService', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService clone() => DeregisterService()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DeregisterService copyWith(void Function(DeregisterService) updates) => super.copyWith((message) => updates(message as DeregisterService)) as DeregisterService;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterService create() => DeregisterService._();
  DeregisterService createEmptyInstance() => create();
  static $pb.PbList<DeregisterService> createRepeated() => $pb.PbList<DeregisterService>();
  @$core.pragma('dart2js:noInline')
  static DeregisterService getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeregisterService>(create);
  static DeregisterService? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
