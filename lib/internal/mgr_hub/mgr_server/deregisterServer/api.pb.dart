//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_server/deregisterServer/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../../kk_etcd_models/pb_server_kk_etcd.pb.dart' as $1;

class DeregisterServer_Input extends $pb.GeneratedMessage {
  factory DeregisterServer_Input({
    $1.PBServer? server,
  }) {
    final $result = create();
    if (server != null) {
      $result.server = server;
    }
    return $result;
  }

  DeregisterServer_Input._() : super();

  factory DeregisterServer_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory DeregisterServer_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'DeregisterServer.Input',
      package:
          const $pb.PackageName(_omitMessageNames ? '' : 'deregisterServer'),
      createEmptyInstance: create)
    ..aOM<$1.PBServer>(1, _omitFieldNames ? '' : 'Server',
        protoName: 'Server', subBuilder: $1.PBServer.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  DeregisterServer_Input clone() =>
      DeregisterServer_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  DeregisterServer_Input copyWith(
          void Function(DeregisterServer_Input) updates) =>
      super.copyWith((message) => updates(message as DeregisterServer_Input))
          as DeregisterServer_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterServer_Input create() => DeregisterServer_Input._();

  DeregisterServer_Input createEmptyInstance() => create();

  static $pb.PbList<DeregisterServer_Input> createRepeated() =>
      $pb.PbList<DeregisterServer_Input>();

  @$core.pragma('dart2js:noInline')
  static DeregisterServer_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<DeregisterServer_Input>(create);
  static DeregisterServer_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $1.PBServer get server => $_getN(0);

  @$pb.TagNumber(1)
  set server($1.PBServer v) {
    setField(1, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasServer() => $_has(0);

  @$pb.TagNumber(1)
  void clearServer() => clearField(1);

  @$pb.TagNumber(1)
  $1.PBServer ensureServer() => $_ensure(0);
}

class DeregisterServer_Output extends $pb.GeneratedMessage {
  factory DeregisterServer_Output() => create();

  DeregisterServer_Output._() : super();

  factory DeregisterServer_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory DeregisterServer_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'DeregisterServer.Output',
      package:
          const $pb.PackageName(_omitMessageNames ? '' : 'deregisterServer'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  DeregisterServer_Output clone() =>
      DeregisterServer_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  DeregisterServer_Output copyWith(
          void Function(DeregisterServer_Output) updates) =>
      super.copyWith((message) => updates(message as DeregisterServer_Output))
          as DeregisterServer_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterServer_Output create() => DeregisterServer_Output._();

  DeregisterServer_Output createEmptyInstance() => create();

  static $pb.PbList<DeregisterServer_Output> createRepeated() =>
      $pb.PbList<DeregisterServer_Output>();

  @$core.pragma('dart2js:noInline')
  static DeregisterServer_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<DeregisterServer_Output>(create);
  static DeregisterServer_Output? _defaultInstance;
}

/// deregister server
class DeregisterServer extends $pb.GeneratedMessage {
  factory DeregisterServer() => create();

  DeregisterServer._() : super();

  factory DeregisterServer.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory DeregisterServer.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'DeregisterServer',
      package:
          const $pb.PackageName(_omitMessageNames ? '' : 'deregisterServer'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  DeregisterServer clone() => DeregisterServer()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  DeregisterServer copyWith(void Function(DeregisterServer) updates) =>
      super.copyWith((message) => updates(message as DeregisterServer))
          as DeregisterServer;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeregisterServer create() => DeregisterServer._();

  DeregisterServer createEmptyInstance() => create();

  static $pb.PbList<DeregisterServer> createRepeated() =>
      $pb.PbList<DeregisterServer>();

  @$core.pragma('dart2js:noInline')
  static DeregisterServer getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<DeregisterServer>(create);
  static DeregisterServer? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
