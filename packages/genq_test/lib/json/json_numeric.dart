import 'package:genq/genq.dart';

part 'json_numeric.genq.dart';

@Genq(json: true)
class JsonNumeric with _$JsonNumeric {
  factory JsonNumeric({
    required int intField,
    required double doubleField,
    required int? intFieldNullable,
    required double? doubleFieldNullable,
    required num numField,
    required num? numFieldNullable,
  }) = _JsonNumeric;
}
