part of 'user.dart';

mixin _$User {
  String get username => throw UnimplementedError();
  String get email => throw UnimplementedError();
  int? get age => throw UnimplementedError();
  Address? get address => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  @override
  final String username;

  @override
  final String email;

  @override
  final int? age;

  @override
  final Address? address;

  _User({
    required this.username,
    required this.email,
    required this.age,
    this.address,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(username: $username, email: $email, age: $age, address: $address)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!identical(other.username, username) && other.username != username) return false;
    if (!identical(other.email, email) && other.email != email) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!identical(other.address, address) && other.address != address) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      username,
      email,
      age,
      address,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    String username,
    String email,
    int? age,
    Address? address,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User value;

  _$UserCopyWithImpl(this.value);

  @override
  User call({
    Object? username = genq,
    Object? email = genq,
    Object? age = genq,
    Object? address = genq,
  }) {
    return User(
      username: username == genq ? value.username : username as String,
      email: email == genq ? value.email : email as String,
      age: age == genq ? value.age : age as int?,
      address: address == genq ? value.address : address as Address?,
    );
  }
}