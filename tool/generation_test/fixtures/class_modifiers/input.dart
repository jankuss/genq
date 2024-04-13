import 'package:genq/genq.dart';

part "input.genq.dart";

@genq
abstract interface class A with _$A {
  factory A({
    required String name,
  }) = _A;
}

@genq
interface class B with _$B {
  factory B({
    required String name,
  }) = _B;
}

