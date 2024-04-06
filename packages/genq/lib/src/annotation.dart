export 'package:collection/collection.dart' show DeepCollectionEquality;

class Genq {
  final bool json;

  const Genq({
    this.json = false,
  });
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
