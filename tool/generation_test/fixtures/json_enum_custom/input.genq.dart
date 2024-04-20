part of 'input.dart';

mixin _$Account {
  String get email => throw UnimplementedError();
  Deez get accountType => throw UnimplementedError();
  int get age => throw UnimplementedError();

  $AccountCopyWith get copyWith => throw UnimplementedError();
}

class _Account implements Account {
  @override
  final String email;

  @override
  final Deez accountType;

  @override
  final int age;

  _Account({
    required this.email,
    required this.accountType,
    required this.age,
  });

  @override
  $AccountCopyWith get copyWith => _$AccountCopyWithImpl(this);

  @override
  String toString() {
    return "Account(email: $email, accountType: $accountType, age: $age)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Account) return false;
    if (!identical(other.email, email) && other.email != email) return false;
    if (!identical(other.accountType, accountType) && other.accountType != accountType) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      email,
      accountType,
      age,
    );
  }
}

abstract class $AccountCopyWith {
  Account call({
    String email,
    Deez accountType,
    int age,
  });
}

class _$AccountCopyWithImpl implements $AccountCopyWith {
  final _$Account value;

  _$AccountCopyWithImpl(this.value);

  @override
  Account call({
    Object? email = genq,
    Object? accountType = genq,
    Object? age = genq,
  }) {
    return Account(
      email: email == genq ? value.email : email as String,
      accountType: accountType == genq ? value.accountType : accountType as Deez,
      age: age == genq ? value.age : age as int,
    );
  }
}

Account $AccountFromJson(Map<String, dynamic> json) {
  return Account(
    email: json['email'] as String,
    accountType: Deez._fromJson(json['accountType']),
    age: alwaysEleven(json['age']),
  );
}

Map<String, dynamic> $AccountToJson(Account obj) {
  return {
    'email': obj.email,
    'accountType': Deez._toJson(obj.accountType),
    'age': alwaysEleven(obj.age),
  };
}
