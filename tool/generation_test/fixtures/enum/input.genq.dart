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
  final _$Account __value;

  _$AccountCopyWithImpl(this.__value);

  @override
  Account call({
    Object? email = genq,
    Object? accountType = genq,
  }) {
    return Account(
      email: email == genq ? __value.email : email as String,
      accountType: accountType == genq ? __value.accountType : accountType as AccountType,
    );
  }
}