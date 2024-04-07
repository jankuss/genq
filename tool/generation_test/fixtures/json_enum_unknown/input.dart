import 'package:genq/genq.dart';

part 'input.genq.dart';

@Genq(json: true)
class Account with _$Account {
  factory Account({
    required String email,
    @JsonKey(name: 'test', unknownEnumValue: AccountType.unknown) 
    required AccountType accountType,
  }) = _Account;
}

@GenqJsonEnum()
enum AccountType { free, premium, vip, unknown }
