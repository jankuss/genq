import 'package:genq/genq.dart';

part 'input.genq.dart';

@Genq(json: true)
class User with _$User {
  factory User({
    @JsonKey(name: 'full_name', defaultValue: 'Supername') required String name,
    @JsonKey(defaultValue: 999) required int? age,
    required bool registered,
  }) = _User;
}
