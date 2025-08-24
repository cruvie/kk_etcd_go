// This is a generated file - do not edit.
//
// Generated from internal/service_hub/service/api_def/service.proto.

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

import 'DeregisterService.pb.dart' as $0;
import 'ServiceList.pb.dart' as $1;

export 'service.pb.dart';

@$pb.GrpcServiceName('kk_etcd.Service')
class ServiceClient extends $grpc.Client {
  /// The hostname for this service.
  static const $core.String defaultHost = '';

  /// OAuth scopes needed for the client.
  static const $core.List<$core.String> oauthScopes = [
    '',
  ];

  ServiceClient(super.channel, {super.options, super.interceptors});

  $grpc.ResponseFuture<$0.DeregisterService_Output> deregisterService(
    $0.DeregisterService_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$deregisterService, request, options: options);
  }

  $grpc.ResponseFuture<$1.ServiceList_Output> serviceList(
    $1.ServiceList_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$serviceList, request, options: options);
  }

  // method descriptors

  static final _$deregisterService = $grpc.ClientMethod<
          $0.DeregisterService_Input, $0.DeregisterService_Output>(
      '/kk_etcd.Service/DeregisterService',
      ($0.DeregisterService_Input value) => value.writeToBuffer(),
      $0.DeregisterService_Output.fromBuffer);
  static final _$serviceList =
      $grpc.ClientMethod<$1.ServiceList_Input, $1.ServiceList_Output>(
          '/kk_etcd.Service/ServiceList',
          ($1.ServiceList_Input value) => value.writeToBuffer(),
          $1.ServiceList_Output.fromBuffer);
}

@$pb.GrpcServiceName('kk_etcd.Service')
abstract class ServiceBase extends $grpc.Service {
  $core.String get $name => 'kk_etcd.Service';

  ServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.DeregisterService_Input,
            $0.DeregisterService_Output>(
        'DeregisterService',
        deregisterService_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.DeregisterService_Input.fromBuffer(value),
        ($0.DeregisterService_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.ServiceList_Input, $1.ServiceList_Output>(
        'ServiceList',
        serviceList_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.ServiceList_Input.fromBuffer(value),
        ($1.ServiceList_Output value) => value.writeToBuffer()));
  }

  $async.Future<$0.DeregisterService_Output> deregisterService_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$0.DeregisterService_Input> $request) async {
    return deregisterService($call, await $request);
  }

  $async.Future<$0.DeregisterService_Output> deregisterService(
      $grpc.ServiceCall call, $0.DeregisterService_Input request);

  $async.Future<$1.ServiceList_Output> serviceList_Pre($grpc.ServiceCall $call,
      $async.Future<$1.ServiceList_Input> $request) async {
    return serviceList($call, await $request);
  }

  $async.Future<$1.ServiceList_Output> serviceList(
      $grpc.ServiceCall call, $1.ServiceList_Input request);
}
