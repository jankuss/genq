import 'package:genq/genq.dart';

part "input.genq.dart";

@genq
class User with _$User {
  factory User({
    required String name,
    required int? age,
    required bool registered,
    required String enumReservedName,
    required String classReservedName,
  }) = _User;
}
