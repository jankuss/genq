import 'package:genq/genq.dart';

part 'input.genq.dart';

@Genq(json: true)
class User with _$User {
  factory User({
    @JsonKey(name: 'full_name', defaultValue: 'Supername') required String name,
    @JsonKey(defaultValue: 999) required int? age,
    @JsonKey(defaultValue: 99.9) required double? someDoubleValueNullable,
    @JsonKey(name: 'some_double_value', defaultValue: 1234567.89123456789) required double someDoubleValue,
    required bool registered,
  }) = _User;
}
