import 'package:genq/genq.dart';

part 'input.genq.dart';

class Deez {
  final String type;

  const Deez._(this.type);

  static Deez _fromJson(String type) {
    return Deez._(type);
  }

  static String _toJson(Deez deez) {
    return deez.type;
  }
}

@Genq(json: true)
class Account with _$Account {
  factory Account({
    required String email,
    @JsonKey(fromJson: Deez._fromJson, toJson: Deez._toJson) required Deez accountType,
  }) = _Account;
}

@GenqJsonEnum()
enum AccountType { free, premium, vip, unknown }
