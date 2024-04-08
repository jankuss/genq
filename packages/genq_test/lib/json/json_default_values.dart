import 'package:genq/genq.dart';

part 'json_default_values.genq.dart';

@Genq(json: true)
class TestDefaultValues with _$TestDefaultValues {
  factory TestDefaultValues({
    @JsonKey(defaultValue: 'Supername') required String name,
    @JsonKey(defaultValue: 999) required int? age,
    required bool registered,
  }) = _TestDefaultValues;
}
