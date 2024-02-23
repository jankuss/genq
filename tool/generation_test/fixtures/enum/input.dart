import 'package:genq/genq.dart';

part 'input.genq.dart';

enum AccountType { free, premium, vip }

@genq
class Account with _$Account {
  factory Account({
    required String email,
    required AccountType accountType,
  }) = _Account;
}
