part of 'input.dart';

mixin _$Account {
  String get email => throw UnimplementedError();
  AccountType get accountType => throw UnimplementedError();

  $AccountCopyWith get copyWith => throw UnimplementedError();
}

class _Account implements Account {
  @override
  final String email;

  @override
  final AccountType accountType;

  _Account({
    required this.email,
    required this.accountType,
  });

  @override
  $AccountCopyWith get copyWith => _$AccountCopyWithImpl(this);

  @override
  String toString() {
    return "Account(email: $email, accountType: $accountType)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Account) return false;
    if (!identical(other.email, email) && other.email != email) return false;
    if (!identical(other.accountType, accountType) && other.accountType != accountType) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      email,
      accountType,
    );
  }
}

abstract class $AccountCopyWith {
  Account call({
    String email,
    AccountType accountType,
  });
}

class _$AccountCopyWithImpl implements $AccountCopyWith {
  final _$Account value;

  _$AccountCopyWithImpl(this.value);

  @override
  Account call({
    Object? email = genq,
    Object? accountType = genq,
  }) {
    return Account(
      email: email == genq ? value.email : email as String,
      accountType: accountType == genq ? value.accountType : accountType as AccountType,
    );
  }
}

Account $AccountFromJson(Map<String, dynamic> json) {
  return Account(
    email: json['email'] as String,
    accountType: $AccountTypeFromJson(json['test'], AccountType.unknown),
  );
}

Map<String, dynamic> $AccountToJson(Account obj) {
  return {
    'email': obj.email,
    'test': $AccountTypeToJson(obj.accountType),
  };
}


AccountType $AccountTypeFromJson(Object json, [AccountType? unknownEnumValue]) {
  switch (json) {
    case "free":
      return AccountType.free;
    case "premium":
      return AccountType.premium;
    case "vip":
      return AccountType.vip;
    case "unknown":
      return AccountType.unknown;
    default:
      if (unknownEnumValue != null) {
        return unknownEnumValue;
      } else {
        throw UnsupportedError('The value $json is not a valid value for enum AccountType');
      }
  }
}

Object $AccountTypeToJson(AccountType value) {
  switch (value) {
    case AccountType.free:
      return "free";
    case AccountType.premium:
      return "premium";
    case AccountType.vip:
      return "vip";
    case AccountType.unknown:
      return "unknown";
    default:
      throw UnsupportedError('Could not map $value to a JSON value');
  }
}