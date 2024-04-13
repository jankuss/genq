part of 'input.dart';

mixin _$A {
  String get name => throw UnimplementedError();

  $ACopyWith get copyWith => throw UnimplementedError();
}

class _A implements A {
  @override
  final String name;

  _A({
    required this.name,
  });

  @override
  $ACopyWith get copyWith => _$ACopyWithImpl(this);

  @override
  String toString() {
    return "A(name: $name)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! A) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
    );
  }
}

abstract class $ACopyWith {
  A call({
    String name,
  });
}

class _$ACopyWithImpl implements $ACopyWith {
  final _$A value;

  _$ACopyWithImpl(this.value);

  @override
  A call({
    Object? name = genq,
  }) {
    return A(
      name: name == genq ? value.name : name as String,
    );
  }
}

mixin _$B {
  String get name => throw UnimplementedError();

  $BCopyWith get copyWith => throw UnimplementedError();
}

class _B implements B {
  @override
  final String name;

  _B({
    required this.name,
  });

  @override
  $BCopyWith get copyWith => _$BCopyWithImpl(this);

  @override
  String toString() {
    return "B(name: $name)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! B) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
    );
  }
}

abstract class $BCopyWith {
  B call({
    String name,
  });
}

class _$BCopyWithImpl implements $BCopyWith {
  final _$B value;

  _$BCopyWithImpl(this.value);

  @override
  B call({
    Object? name = genq,
  }) {
    return B(
      name: name == genq ? value.name : name as String,
    );
  }
}