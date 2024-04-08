import 'package:genq/genq.dart';

part 'json_etf.genq.dart';

class ISIN {
  final String value;

  const ISIN(this.value);

  static ISIN fromJson(String value) => ISIN(value);
  static String toJson(ISIN isin) => isin.value;

  @override
  operator ==(Object other) => other is ISIN && other.value == value;

  @override
  int get hashCode => value.hashCode;

}

@Genq(json: true)
class ETF with _$ETF {
  factory ETF({
    @JsonKey(fromJson: ISIN.fromJson, toJson: ISIN.toJson)
    required ISIN isin,
    required String name,
    required double price,
  }) = _ETF;
}
