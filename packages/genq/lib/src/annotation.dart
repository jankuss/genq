import 'package:meta/meta_meta.dart';

export 'package:collection/collection.dart' show DeepCollectionEquality;

/// Annotation to mark a class for code generation with genq. If you need to
/// supply options, you can use the [Genq] class instead.
const genq = Genq();

/// Annotation to mark a class for code generation with genq. If you don't have
/// to supply any options, you can use the top-level [genq] annotation instead.
@Target({TargetKind.classType})
class Genq {
  /// Whether or not there should be ToJson and FromJson generated for this class.
  final bool json;

  const Genq({
    this.json = false,
  });
}

/// Annotation to mark an enum for JSON serialization code generation with genq.
@Target({TargetKind.enumType})
class GenqJsonEnum {
  const GenqJsonEnum();
}

@Target({TargetKind.parameter})
class JsonKey {
  /// The key in the JSON map that corresponds to this field.
  final String? name;

  /// A custom function to use when decoding this field.
  final Function? fromJson;

  /// A custom function to use when encoding this field.
  final Function? toJson;

  /// When decoding an enum, the value to use when the value does not match any enums.
  final Enum? unknownEnumValue;

  /// The default value to use when the JSON value is `null`.
  final Object? defaultValue;

  const JsonKey({
    this.name,
    this.unknownEnumValue,
    this.fromJson,
    this.toJson,
    this.defaultValue,
  });
}

/// Annotation to specify how an enum should be serialized.
class JsonValue {
  /// The value to use when serializing this enum.
  final dynamic value;

  const JsonValue(this.value);
}

/// Allows you to specify a default value for a field.
@Target({TargetKind.parameter})
class Default {
  final Object? value;

  const Default(this.value);
}
