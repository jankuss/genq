part of 'input.dart';

mixin _$User {
  String get name => throw UnimplementedError();
  int? get age => throw UnimplementedError();
  bool get registered => throw UnimplementedError();
  Address get address => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User extends User {
  @override
  final String name;

  @override
  final int? age;

  @override
  final bool registered;

  @override
  final Address address;

  _User({
    required this.name,
    required this.age,
    required this.registered,
    required this.address,
  }) : super._();

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(name: $name, age: $age, registered: $registered, address: $address)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!identical(other.registered, registered) && other.registered != registered) return false;
    if (!identical(other.address, address) && other.address != address) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
      age,
      registered,
      address,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    String name,
    int? age,
    bool registered,
    Address address,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User value;

  _$UserCopyWithImpl(this.value);

  @override
  User call({
    Object? name = genq,
    Object? age = genq,
    Object? registered = genq,
    Object? address = genq,
  }) {
    return User(
      name: name == genq ? value.name : name as String,
      age: age == genq ? value.age : age as int?,
      registered: registered == genq ? value.registered : registered as bool,
      address: address == genq ? value.address : address as Address,
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
