part of 'test_classes.dart';

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

mixin _$Address {
  String get street => throw UnimplementedError();
  int get zipCode => throw UnimplementedError();
  String get city => throw UnimplementedError();

  $AddressCopyWith get copyWith => throw UnimplementedError();
}

class _Address implements Address {
  @override
  final String street;

  @override
  final int zipCode;

  @override
  final String city;

  _Address({
    required this.street,
    required this.zipCode,
    required this.city,
  });

  @override
  $AddressCopyWith get copyWith => _$AddressCopyWithImpl(this);

  @override
  String toString() {
    return "Address(street: $street, zipCode: $zipCode, city: $city)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Address) return false;
    if (!identical(other.street, street) && other.street != street) return false;
    if (!identical(other.zipCode, zipCode) && other.zipCode != zipCode) return false;
    if (!identical(other.city, city) && other.city != city) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      street,
      zipCode,
      city,
    );
  }
}

abstract class $AddressCopyWith {
  Address call({
    String street,
    int zipCode,
    String city,
  });
}

class _$AddressCopyWithImpl implements $AddressCopyWith {
  final _$Address value;

  _$AddressCopyWithImpl(this.value);

  @override
  Address call({
    Object? street = genq,
    Object? zipCode = genq,
    Object? city = genq,
  }) {
    return Address(
      street: street == genq ? value.street : street as String,
      zipCode: zipCode == genq ? value.zipCode : zipCode as int,
      city: city == genq ? value.city : city as String,
    );
  }
}

mixin _$CollectionTests {
  List<User> get users => throw UnimplementedError();
  Map<String, User> get userMap => throw UnimplementedError();
  Set<User> get userSet => throw UnimplementedError();

  $CollectionTestsCopyWith get copyWith => throw UnimplementedError();
}

class _CollectionTests implements CollectionTests {
  @override
  final List<User> users;

  @override
  final Map<String, User> userMap;

  @override
  final Set<User> userSet;

  _CollectionTests({
    required this.users,
    required this.userMap,
    required this.userSet,
  });

  @override
  $CollectionTestsCopyWith get copyWith => _$CollectionTestsCopyWithImpl(this);

  @override
  String toString() {
    return "CollectionTests(users: $users, userMap: $userMap, userSet: $userSet)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! CollectionTests) return false;
    if (!const DeepCollectionEquality().equals(other.users, users)) return false;
    if (!const DeepCollectionEquality().equals(other.userMap, userMap)) return false;
    if (!const DeepCollectionEquality().equals(other.userSet, userSet)) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      users,
      userMap,
      userSet,
    );
  }
}

abstract class $CollectionTestsCopyWith {
  CollectionTests call({
    List<User> users,
    Map<String, User> userMap,
    Set<User> userSet,
  });
}

class _$CollectionTestsCopyWithImpl implements $CollectionTestsCopyWith {
  final _$CollectionTests value;

  _$CollectionTestsCopyWithImpl(this.value);

  @override
  CollectionTests call({
    Object? users = genq,
    Object? userMap = genq,
    Object? userSet = genq,
  }) {
    return CollectionTests(
      users: users == genq ? value.users : users as List<User>,
      userMap: userMap == genq ? value.userMap : userMap as Map<String, User>,
      userSet: userSet == genq ? value.userSet : userSet as Set<User>,
    );
  }
}