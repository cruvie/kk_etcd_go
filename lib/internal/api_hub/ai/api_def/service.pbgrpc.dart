// This is a generated file - do not edit.
//
// Generated from internal/api_hub/ai/api_def/service.proto.

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

import 'Query.pb.dart' as $0;

export 'service.pb.dart';

@$pb.GrpcServiceName('api_def.AI')
class AIClient extends $grpc.Client {
  /// The hostname for this service.
  static const $core.String defaultHost = '';

  /// OAuth scopes needed for the client.
  static const $core.List<$core.String> oauthScopes = [
    '',
  ];

  AIClient(super.channel, {super.options, super.interceptors});

  $grpc.ResponseFuture<$0.Query_Output> query(
    $0.Query_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$query, request, options: options);
  }

  // method descriptors

  static final _$query = $grpc.ClientMethod<$0.Query_Input, $0.Query_Output>(
      '/api_def.AI/Query',
      ($0.Query_Input value) => value.writeToBuffer(),
      $0.Query_Output.fromBuffer);
}

@$pb.GrpcServiceName('api_def.AI')
abstract class AIServiceBase extends $grpc.Service {
  $core.String get $name => 'api_def.AI';

  AIServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Query_Input, $0.Query_Output>(
        'Query',
        query_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Query_Input.fromBuffer(value),
        ($0.Query_Output value) => value.writeToBuffer()));
  }

  $async.Future<$0.Query_Output> query_Pre(
      $grpc.ServiceCall $call, $async.Future<$0.Query_Input> $request) async {
    return query($call, await $request);
  }

  $async.Future<$0.Query_Output> query(
      $grpc.ServiceCall call, $0.Query_Input request);
}
