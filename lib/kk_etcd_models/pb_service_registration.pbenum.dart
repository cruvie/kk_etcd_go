//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_service_registration.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PBServiceType extends $pb.ProtobufEnum {
  static const PBServiceType Unknown = PBServiceType._(0, _omitEnumNames ? '' : 'Unknown');
  static const PBServiceType Http = PBServiceType._(1, _omitEnumNames ? '' : 'Http');
  static const PBServiceType Grpc = PBServiceType._(2, _omitEnumNames ? '' : 'Grpc');

  static const $core.List<PBServiceType> values = <PBServiceType> [
    Unknown,
    Http,
    Grpc,
  ];

  static final $core.Map<$core.int, PBServiceType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static PBServiceType? valueOf($core.int value) => _byValue[value];

  const PBServiceType._(super.v, super.n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
