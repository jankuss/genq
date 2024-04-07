export 'package:collection/collection.dart' show DeepCollectionEquality;

/// Annotation to mark a class for code generation with genq. If you need to
/// supply options, you can use the [Genq] class instead.
const genq = Genq();

/// Annotation to mark a class for code generation with genq. If you don't have
/// to supply any options, you can use the top-level [genq] annotation instead.
///
/// Options:
/// [json] - Wether to generate a toJson/fromJson method for the class.
class Genq {
  final bool json;

  const Genq({
    this.json = false,
  });
}

/// Annotation to mark an enum for JSON serialization code generation with genq.
class GenqJsonEnum {
  const GenqJsonEnum();
}

class JsonKey {
  final String? name;
  final Function? fromJson;
  final Function? toJson;
  final Enum? unknownEnumValue;

  const JsonKey({
    this.name,
    this.unknownEnumValue,
    this.fromJson,
    this.toJson,
  });
}

class JsonValue {
  final Object value;

  const JsonValue(this.value);
}
