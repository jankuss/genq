part of 'input.dart';

mixin _$User {
  String get name => throw UnimplementedError();
  int? get age => throw UnimplementedError();
  bool get registered => throw UnimplementedError();
  String get enumReservedName => throw UnimplementedError();
  String get classReservedName => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  @override
  final String name;

  @override
  final int? age;

  @override
  final bool registered;

  @override
  final String enumReservedName;

  @override
  final String classReservedName;

  _User({
    required this.name,
    required this.age,
    required this.registered,
    required this.enumReservedName,
    required this.classReservedName,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(name: $name, age: $age, registered: $registered, enumReservedName: $enumReservedName, classReservedName: $classReservedName)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!identical(other.registered, registered) && other.registered != registered) return false;
    if (!identical(other.enumReservedName, enumReservedName) && other.enumReservedName != enumReservedName) return false;
    if (!identical(other.classReservedName, classReservedName) && other.classReservedName != classReservedName) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
      age,
      registered,
      enumReservedName,
      classReservedName,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    String name,
    int? age,
    bool registered,
    String enumReservedName,
    String classReservedName,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User value;

  _$UserCopyWithImpl(this.value);

  @override
  User call({
    Object? name = genq,
    Object? age = genq,
    Object? registered = genq,
    Object? enumReservedName = genq,
    Object? classReservedName = genq,
  }) {
    return User(
      name: name == genq ? value.name : name as String,
      age: age == genq ? value.age : age as int?,
      registered: registered == genq ? value.registered : registered as bool,
      enumReservedName: enumReservedName == genq ? value.enumReservedName : enumReservedName as String,
      classReservedName: classReservedName == genq ? value.classReservedName : classReservedName as String,
    );
  }
}