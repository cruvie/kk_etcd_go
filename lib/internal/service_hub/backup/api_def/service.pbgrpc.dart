// This is a generated file - do not edit.
//
// Generated from internal/service_hub/backup/api_def/service.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'AllKVsBackup.pb.dart' as $0;
import 'AllKVsRestore.pb.dart' as $1;
import 'Snapshot.pb.dart' as $2;
import 'SnapshotInfo.pb.dart' as $3;
import 'SnapshotRestore.pb.dart' as $4;

export 'service.pb.dart';

@$pb.GrpcServiceName('api_def.Backup')
class BackupClient extends $grpc.Client {
  /// The hostname for this service.
  static const $core.String defaultHost = '';

  /// OAuth scopes needed for the client.
  static const $core.List<$core.String> oauthScopes = [
    '',
  ];

  BackupClient(super.channel, {super.options, super.interceptors});

  $grpc.ResponseFuture<$0.AllKVsBackup_Output> allKVsBackup(
    $0.AllKVsBackup_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$allKVsBackup, request, options: options);
  }

  $grpc.ResponseFuture<$1.AllKVsRestore_Output> allKVsRestore(
    $1.AllKVsRestore_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$allKVsRestore, request, options: options);
  }

  $grpc.ResponseFuture<$2.Snapshot_Output> snapshot(
    $2.Snapshot_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$snapshot, request, options: options);
  }

  $grpc.ResponseFuture<$3.SnapshotInfo_Output> snapshotInfo(
    $3.SnapshotInfo_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$snapshotInfo, request, options: options);
  }

  $grpc.ResponseFuture<$4.SnapshotRestore_Output> snapshotRestore(
    $4.SnapshotRestore_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$snapshotRestore, request, options: options);
  }

  // method descriptors

  static final _$allKVsBackup =
      $grpc.ClientMethod<$0.AllKVsBackup_Input, $0.AllKVsBackup_Output>(
          '/api_def.Backup/AllKVsBackup',
          ($0.AllKVsBackup_Input value) => value.writeToBuffer(),
          $0.AllKVsBackup_Output.fromBuffer);
  static final _$allKVsRestore =
      $grpc.ClientMethod<$1.AllKVsRestore_Input, $1.AllKVsRestore_Output>(
          '/api_def.Backup/AllKVsRestore',
          ($1.AllKVsRestore_Input value) => value.writeToBuffer(),
          $1.AllKVsRestore_Output.fromBuffer);
  static final _$snapshot =
      $grpc.ClientMethod<$2.Snapshot_Input, $2.Snapshot_Output>(
          '/api_def.Backup/Snapshot',
          ($2.Snapshot_Input value) => value.writeToBuffer(),
          $2.Snapshot_Output.fromBuffer);
  static final _$snapshotInfo =
      $grpc.ClientMethod<$3.SnapshotInfo_Input, $3.SnapshotInfo_Output>(
          '/api_def.Backup/SnapshotInfo',
          ($3.SnapshotInfo_Input value) => value.writeToBuffer(),
          $3.SnapshotInfo_Output.fromBuffer);
  static final _$snapshotRestore =
      $grpc.ClientMethod<$4.SnapshotRestore_Input, $4.SnapshotRestore_Output>(
          '/api_def.Backup/SnapshotRestore',
          ($4.SnapshotRestore_Input value) => value.writeToBuffer(),
          $4.SnapshotRestore_Output.fromBuffer);
}

@$pb.GrpcServiceName('api_def.Backup')
abstract class BackupServiceBase extends $grpc.Service {
  $core.String get $name => 'api_def.Backup';

  BackupServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$0.AllKVsBackup_Input, $0.AllKVsBackup_Output>(
            'AllKVsBackup',
            allKVsBackup_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.AllKVsBackup_Input.fromBuffer(value),
            ($0.AllKVsBackup_Output value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$1.AllKVsRestore_Input, $1.AllKVsRestore_Output>(
            'AllKVsRestore',
            allKVsRestore_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.AllKVsRestore_Input.fromBuffer(value),
            ($1.AllKVsRestore_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Snapshot_Input, $2.Snapshot_Output>(
        'Snapshot',
        snapshot_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Snapshot_Input.fromBuffer(value),
        ($2.Snapshot_Output value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$3.SnapshotInfo_Input, $3.SnapshotInfo_Output>(
            'SnapshotInfo',
            snapshotInfo_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $3.SnapshotInfo_Input.fromBuffer(value),
            ($3.SnapshotInfo_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.SnapshotRestore_Input,
            $4.SnapshotRestore_Output>(
        'SnapshotRestore',
        snapshotRestore_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $4.SnapshotRestore_Input.fromBuffer(value),
        ($4.SnapshotRestore_Output value) => value.writeToBuffer()));
  }

  $async.Future<$0.AllKVsBackup_Output> allKVsBackup_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$0.AllKVsBackup_Input> $request) async {
    return allKVsBackup($call, await $request);
  }

  $async.Future<$0.AllKVsBackup_Output> allKVsBackup(
      $grpc.ServiceCall call, $0.AllKVsBackup_Input request);

  $async.Future<$1.AllKVsRestore_Output> allKVsRestore_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$1.AllKVsRestore_Input> $request) async {
    return allKVsRestore($call, await $request);
  }

  $async.Future<$1.AllKVsRestore_Output> allKVsRestore(
      $grpc.ServiceCall call, $1.AllKVsRestore_Input request);

  $async.Future<$2.Snapshot_Output> snapshot_Pre($grpc.ServiceCall $call,
      $async.Future<$2.Snapshot_Input> $request) async {
    return snapshot($call, await $request);
  }

  $async.Future<$2.Snapshot_Output> snapshot(
      $grpc.ServiceCall call, $2.Snapshot_Input request);

  $async.Future<$3.SnapshotInfo_Output> snapshotInfo_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$3.SnapshotInfo_Input> $request) async {
    return snapshotInfo($call, await $request);
  }

  $async.Future<$3.SnapshotInfo_Output> snapshotInfo(
      $grpc.ServiceCall call, $3.SnapshotInfo_Input request);

  $async.Future<$4.SnapshotRestore_Output> snapshotRestore_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$4.SnapshotRestore_Input> $request) async {
    return snapshotRestore($call, await $request);
  }

  $async.Future<$4.SnapshotRestore_Output> snapshotRestore(
      $grpc.ServiceCall call, $4.SnapshotRestore_Input request);
}
