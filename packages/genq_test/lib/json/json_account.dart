import 'package:genq/genq.dart';

part 'json_account.genq.dart';

@Genq(json: true)
class JsonAccount with _$JsonAccount {
  factory JsonAccount({
    required String email,
    @JsonKey(unknownEnumValue: JsonStatus.unknown)
    required JsonStatus status,
  }) = _JsonAccount;
}

@GenqJsonEnum()
enum JsonStatus {
  active,
  inactive,
  unknown,
}
