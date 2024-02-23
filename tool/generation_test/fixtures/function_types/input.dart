import 'package:genq/genq.dart';

part "input.genq.dart";

@genq
class User with _$User {
  factory User({
    required String name,
    required int? age,
    required bool registered,
    required int Function(String str) a,
    required bool Function(int value) Function(String str) b,
    required User Function(String str) Function(int value, { required String test }) Function(String str) c,
    required void Function(void Function() a, void Function() b) Function(void Function(void Function() d) c) d,
  }) = _User;
}
