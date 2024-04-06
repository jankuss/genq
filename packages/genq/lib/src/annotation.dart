export 'package:collection/collection.dart' show DeepCollectionEquality;

class Genq {
  const Genq();
}

const genq = Genq();

class JsonKey {
  final String? name;

  const JsonKey({
    this.name,
  });
}

class JsonEnum {
  const JsonEnum();
}

class JsonValue {
  final Object value;

  const JsonValue(this.value);
}
