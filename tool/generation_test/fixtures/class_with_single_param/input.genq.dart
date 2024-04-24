part of 'input.dart';

mixin _$Test {
  int get param1 => throw UnimplementedError();

  $TestCopyWith get copyWith => throw UnimplementedError();
}

class _Test implements Test {
  @override
  final int param1;

  _Test({
    required this.param1,
  });

  @override
  $TestCopyWith get copyWith => _$TestCopyWithImpl(this);

  @override
  String toString() {
    return "Test(param1: $param1)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Test) return false;
    if (!identical(other.param1, param1) && other.param1 != param1) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      param1,
    );
  }
}

abstract class $TestCopyWith {
  Test call({
    int param1,
  });
}

class _$TestCopyWithImpl implements $TestCopyWith {
  final _$Test value;

  _$TestCopyWithImpl(this.value);

  @override
  Test call({
    Object? param1 = genq,
  }) {
    return Test(
      param1: param1 == genq ? value.param1 : param1 as int,
    );
  }
}
