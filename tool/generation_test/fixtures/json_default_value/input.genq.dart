part of 'input.dart';

mixin _$User {
  String get name => throw UnimplementedError();
  int? get age => throw UnimplementedError();
  double? get someDoubleValueNullable => throw UnimplementedError();
  double get someDoubleValue => throw UnimplementedError();
  bool get registered => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  @override
  final String name;

  @override
  final int? age;

  @override
  final double? someDoubleValueNullable;

  @override
  final double someDoubleValue;

  @override
  final bool registered;

  _User({
    required this.name,
    required this.age,
    required this.someDoubleValueNullable,
    required this.someDoubleValue,
    required this.registered,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(name: $name, age: $age, someDoubleValueNullable: $someDoubleValueNullable, someDoubleValue: $someDoubleValue, registered: $registered)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!identical(other.someDoubleValueNullable, someDoubleValueNullable) && other.someDoubleValueNullable != someDoubleValueNullable) return false;
    if (!identical(other.someDoubleValue, someDoubleValue) && other.someDoubleValue != someDoubleValue) return false;
    if (!identical(other.registered, registered) && other.registered != registered) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
      age,
      someDoubleValueNullable,
      someDoubleValue,
      registered,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    String name,
    int? age,
    double? someDoubleValueNullable,
    double someDoubleValue,
    bool registered,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User __value;

  _$UserCopyWithImpl(this.__value);

  @override
  User call({
    Object? name = genq,
    Object? age = genq,
    Object? someDoubleValueNullable = genq,
    Object? someDoubleValue = genq,
    Object? registered = genq,
  }) {
    return User(
      name: name == genq ? __value.name : name as String,
      age: age == genq ? __value.age : age as int?,
      someDoubleValueNullable: someDoubleValueNullable == genq ? __value.someDoubleValueNullable : someDoubleValueNullable as double?,
      someDoubleValue: someDoubleValue == genq ? __value.someDoubleValue : someDoubleValue as double,
      registered: registered == genq ? __value.registered : registered as bool,
    );
  }
}

User $UserFromJson(Map<String, dynamic> json) {
  return User(
    name: json['full_name'] == null ? 'Supername' : json['full_name'] as String,
    age: json['age'] == null ? 999 : (json['age'] as num).toInt(),
    someDoubleValueNullable: json['someDoubleValueNullable'] == null ? 99.9 : (json['someDoubleValueNullable'] as num).toDouble(),
    someDoubleValue: json['some_double_value'] == null ? 1234567.89123456789 : (json['some_double_value'] as num).toDouble(),
    registered: json['registered'] as bool,
  );
}

Map<String, dynamic> $UserToJson(User obj) {
  return {
    'full_name': obj.name,
    'age': obj.age == null ? null : obj.age!,
    'someDoubleValueNullable': obj.someDoubleValueNullable == null ? null : obj.someDoubleValueNullable!,
    'some_double_value': obj.someDoubleValue,
    'registered': obj.registered,
  };
}