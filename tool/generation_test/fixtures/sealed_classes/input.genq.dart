part of 'input.dart';

mixin _$LoadingState {

  $LoadingStateCopyWith get copyWith => throw UnimplementedError();
}

class _LoadingState implements LoadingState {
  const _LoadingState();

  @override
  $LoadingStateCopyWith get copyWith => _$LoadingStateCopyWithImpl(this);

  @override
  String toString() {
    return "LoadingState()";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! LoadingState) return false;
    return true;
  }

  @override
  int get hashCode {
    return runtimeType.hashCode;
  }
}

abstract class $LoadingStateCopyWith {
  LoadingState call();
}

class _$LoadingStateCopyWithImpl implements $LoadingStateCopyWith {
  // ignore: unused_field
  final _$LoadingState __value;

  _$LoadingStateCopyWithImpl(this.__value);

  @override
  LoadingState call() {
    return LoadingState();
  }
}

mixin _$SuccessState {
  String get name => throw UnimplementedError();
  int get age => throw UnimplementedError();

  $SuccessStateCopyWith get copyWith => throw UnimplementedError();
}

class _SuccessState implements SuccessState {
  @override
  final String name;

  @override
  final int age;

  const _SuccessState({
    required this.name,
    required this.age,
  });

  @override
  $SuccessStateCopyWith get copyWith => _$SuccessStateCopyWithImpl(this);

  @override
  String toString() {
    return "SuccessState(name: $name, age: $age)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! SuccessState) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
      age,
    );
  }
}

abstract class $SuccessStateCopyWith {
  SuccessState call({
    String name,
    int age,
  });
}

class _$SuccessStateCopyWithImpl implements $SuccessStateCopyWith {
  final _$SuccessState __value;

  _$SuccessStateCopyWithImpl(this.__value);

  @override
  SuccessState call({
    Object? name = genq,
    Object? age = genq,
  }) {
    return SuccessState(
      name: name == genq ? __value.name : name as String,
      age: age == genq ? __value.age : age as int,
    );
  }
}