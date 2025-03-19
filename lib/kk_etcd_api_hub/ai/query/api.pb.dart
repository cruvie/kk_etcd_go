//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/ai/query/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Query_Input extends $pb.GeneratedMessage {
  factory Query_Input({
    $core.String? question,
  }) {
    final $result = create();
    if (question != null) {
      $result.question = question;
    }
    return $result;
  }

  Query_Input._() : super();

  factory Query_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory Query_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Query.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'query'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Question', protoName: 'Question')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  Query_Input clone() => Query_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  Query_Input copyWith(void Function(Query_Input) updates) =>
      super.copyWith((message) => updates(message as Query_Input))
          as Query_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Query_Input create() => Query_Input._();

  Query_Input createEmptyInstance() => create();

  static $pb.PbList<Query_Input> createRepeated() => $pb.PbList<Query_Input>();

  @$core.pragma('dart2js:noInline')
  static Query_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<Query_Input>(create);
  static Query_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get question => $_getSZ(0);

  @$pb.TagNumber(1)
  set question($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasQuestion() => $_has(0);

  @$pb.TagNumber(1)
  void clearQuestion() => clearField(1);
}

class Query_Output extends $pb.GeneratedMessage {
  factory Query_Output({
    $core.String? answer,
  }) {
    final $result = create();
    if (answer != null) {
      $result.answer = answer;
    }
    return $result;
  }

  Query_Output._() : super();

  factory Query_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory Query_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Query.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'query'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Answer', protoName: 'Answer')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  Query_Output clone() => Query_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  Query_Output copyWith(void Function(Query_Output) updates) =>
      super.copyWith((message) => updates(message as Query_Output))
          as Query_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Query_Output create() => Query_Output._();

  Query_Output createEmptyInstance() => create();

  static $pb.PbList<Query_Output> createRepeated() =>
      $pb.PbList<Query_Output>();

  @$core.pragma('dart2js:noInline')
  static Query_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<Query_Output>(create);
  static Query_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get answer => $_getSZ(0);

  @$pb.TagNumber(1)
  set answer($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasAnswer() => $_has(0);

  @$pb.TagNumber(1)
  void clearAnswer() => clearField(1);
}

/// query ai
class Query extends $pb.GeneratedMessage {
  factory Query() => create();

  Query._() : super();

  factory Query.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory Query.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Query',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'query'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  Query clone() => Query()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  Query copyWith(void Function(Query) updates) =>
      super.copyWith((message) => updates(message as Query)) as Query;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Query create() => Query._();

  Query createEmptyInstance() => create();

  static $pb.PbList<Query> createRepeated() => $pb.PbList<Query>();

  @$core.pragma('dart2js:noInline')
  static Query getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Query>(create);
  static Query? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
