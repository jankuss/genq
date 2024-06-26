import 'package:genq/genq.dart';

part "input.genq.dart";

@genq
class User with _$User {
  factory User({
    String? name,
    required int? age,
    required bool registered,
  }) = _User;
}
