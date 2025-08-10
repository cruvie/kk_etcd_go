// This is a generated file - do not edit.
//
// Generated from internal/api_hub/ai/api_def/Query.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class Query_Input extends $pb.GeneratedMessage {
  factory Query_Input({
    $core.String? question,
  }) {
    final result = create();
    if (question != null) result.question = question;
    return result;
  }

  Query_Input._();

  factory Query_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory Query_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Query.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Question', protoName: 'Question')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Query_Input clone() => Query_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Query_Input copyWith(void Function(Query_Input) updates) =>
      super.copyWith((message) => updates(message as Query_Input))
          as Query_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Query_Input create() => Query_Input._();
  @$core.override
  Query_Input createEmptyInstance() => create();
  static $pb.PbList<Query_Input> createRepeated() => $pb.PbList<Query_Input>();
  @$core.pragma('dart2js:noInline')
  static Query_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<Query_Input>(create);
  static Query_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get question => $_getSZ(0);
  @$pb.TagNumber(1)
  set question($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasQuestion() => $_has(0);
  @$pb.TagNumber(1)
  void clearQuestion() => $_clearField(1);
}

class Query_Output extends $pb.GeneratedMessage {
  factory Query_Output({
    $core.String? answer,
  }) {
    final result = create();
    if (answer != null) result.answer = answer;
    return result;
  }

  Query_Output._();

  factory Query_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory Query_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Query.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Answer', protoName: 'Answer')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Query_Output clone() => Query_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Query_Output copyWith(void Function(Query_Output) updates) =>
      super.copyWith((message) => updates(message as Query_Output))
          as Query_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Query_Output create() => Query_Output._();
  @$core.override
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
  set answer($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasAnswer() => $_has(0);
  @$pb.TagNumber(1)
  void clearAnswer() => $_clearField(1);
}

/// query ai
class Query extends $pb.GeneratedMessage {
  factory Query() => create();

  Query._();

  factory Query.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory Query.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Query',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Query clone() => Query()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Query copyWith(void Function(Query) updates) =>
      super.copyWith((message) => updates(message as Query)) as Query;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Query create() => Query._();
  @$core.override
  Query createEmptyInstance() => create();
  static $pb.PbList<Query> createRepeated() => $pb.PbList<Query>();
  @$core.pragma('dart2js:noInline')
  static Query getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Query>(create);
  static Query? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
