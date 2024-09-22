part of 'input.dart';

mixin _$User {
  String get name => throw UnimplementedError();
  int? get age => throw UnimplementedError();
  bool get registered => throw UnimplementedError();
  int Function(String str) get a => throw UnimplementedError();
  bool Function(int value) Function(String str) get b => throw UnimplementedError();
  User Function(String str) Function(int value, {String test}) Function(String str) get c => throw UnimplementedError();
  void Function(void Function() a, void Function() b) Function(void Function(void Function() d) c) get d => throw UnimplementedError();

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
  final int Function(String str) a;

  @override
  final bool Function(int value) Function(String str) b;

  @override
  final User Function(String str) Function(int value, {String test}) Function(String str) c;

  @override
  final void Function(void Function() a, void Function() b) Function(void Function(void Function() d) c) d;

  _User({
    required this.name,
    required this.age,
    required this.registered,
    required this.a,
    required this.b,
    required this.c,
    required this.d,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(name: $name, age: $age, registered: $registered, a: $a, b: $b, c: $c, d: $d)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!identical(other.registered, registered) && other.registered != registered) return false;
    if (!identical(other.a, a) && other.a != a) return false;
    if (!identical(other.b, b) && other.b != b) return false;
    if (!identical(other.c, c) && other.c != c) return false;
    if (!identical(other.d, d) && other.d != d) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
      age,
      registered,
      a,
      b,
      c,
      d,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    String name,
    int? age,
    bool registered,
    int Function(String str) a,
    bool Function(int value) Function(String str) b,
    User Function(String str) Function(int value, {String test}) Function(String str) c,
    void Function(void Function() a, void Function() b) Function(void Function(void Function() d) c) d,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User __value;

  _$UserCopyWithImpl(this.__value);

  @override
  User call({
    Object? name = genq,
    Object? age = genq,
    Object? registered = genq,
    Object? a = genq,
    Object? b = genq,
    Object? c = genq,
    Object? d = genq,
  }) {
    return User(
      name: name == genq ? __value.name : name as String,
      age: age == genq ? __value.age : age as int?,
      registered: registered == genq ? __value.registered : registered as bool,
      a: a == genq ? __value.a : a as int Function(String str),
      b: b == genq ? __value.b : b as bool Function(int value) Function(String str),
      c: c == genq ? __value.c : c as User Function(String str) Function(int value, {String test}) Function(String str),
      d: d == genq ? __value.d : d as void Function(void Function() a, void Function() b) Function(void Function(void Function() d) c),
    );
  }
}