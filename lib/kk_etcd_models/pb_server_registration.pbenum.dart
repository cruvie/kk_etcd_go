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

class PBServerType extends $pb.ProtobufEnum {
  static const PBServerType Unknown = PBServerType._(0, _omitEnumNames ? '' : 'Unknown');
  static const PBServerType Http = PBServerType._(1, _omitEnumNames ? '' : 'Http');
  static const PBServerType Grpc = PBServerType._(2, _omitEnumNames ? '' : 'Grpc');

  static const $core.List<PBServerType> values = <PBServerType> [
    Unknown,
    Http,
    Grpc,
  ];

  static final $core.Map<$core.int, PBServerType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static PBServerType? valueOf($core.int value) => _byValue[value];

  const PBServerType._(super.v, super.n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
