part of 'json_default_values.dart';

mixin _$TestDefaultValues {
  String get name => throw UnimplementedError();
  int? get age => throw UnimplementedError();
  bool get registered => throw UnimplementedError();

  $TestDefaultValuesCopyWith get copyWith => throw UnimplementedError();
}

class _TestDefaultValues implements TestDefaultValues {
  @override
  final String name;

  @override
  final int? age;

  @override
  final bool registered;

  _TestDefaultValues({
    required this.name,
    required this.age,
    required this.registered,
  });

  @override
  $TestDefaultValuesCopyWith get copyWith => _$TestDefaultValuesCopyWithImpl(this);

  @override
  String toString() {
    return "TestDefaultValues(name: $name, age: $age, registered: $registered)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! TestDefaultValues) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!identical(other.registered, registered) && other.registered != registered) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
      age,
      registered,
    );
  }
}

abstract class $TestDefaultValuesCopyWith {
  TestDefaultValues call({
    String name,
    int? age,
    bool registered,
  });
}

class _$TestDefaultValuesCopyWithImpl implements $TestDefaultValuesCopyWith {
  final _$TestDefaultValues value;

  _$TestDefaultValuesCopyWithImpl(this.value);

  @override
  TestDefaultValues call({
    Object? name = genq,
    Object? age = genq,
    Object? registered = genq,
  }) {
    return TestDefaultValues(
      name: name == genq ? value.name : name as String,
      age: age == genq ? value.age : age as int?,
      registered: registered == genq ? value.registered : registered as bool,
    );
  }
}

TestDefaultValues $TestDefaultValuesFromJson(Map<String, dynamic> json) {
  return TestDefaultValues(
    name: json['name'] == null ? 'Supername' : (json['name'] as String),
    age: json['age'] == null ? 999 : (json['age'] as int?),
    registered: json['registered'] as bool,
  );
}

Map<String, dynamic> $TestDefaultValuesToJson(TestDefaultValues obj) {
  return {
    'name': obj.name,
    'age': obj.age == null ? null : obj.age!,
    'registered': obj.registered,
  };
}