import 'package:genq/genq.dart';

part "input.genq.dart";

class SomeAnnotation {
  final String value;
  final bool isNice;

  const SomeAnnotation({
    required this.value,
    required this.isNice,
  });
}

@genq
class User with _$User {
  factory User({
    @SomeAnnotation(value: 'Annotation', isNice: true) required String name,
    required int? age,
    required bool registered,
  }) = _User;
}
