part of 'input.dart';

mixin _$User {
  String get name => throw UnimplementedError();
  int? get age => throw UnimplementedError();
  bool get registered => throw UnimplementedError();
  Address? get address => throw UnimplementedError();
  DateTime? get birthday => throw UnimplementedError();
  BigInt? get balance => throw UnimplementedError();
  Object? get someObject => throw UnimplementedError();
  dynamic get someDynamic => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  @override
  final String name;

  @override
  final int? age;

  @override
  final bool registered;

  @override
  final Address? address;

  @override
  final DateTime? birthday;

  @override
  final BigInt? balance;

  @override
  final Object? someObject;

  @override
  final dynamic someDynamic;

  _User({
    required this.name,
    required this.age,
    required this.registered,
    required this.address,
    required this.birthday,
    required this.balance,
    required this.someObject,
    required this.someDynamic,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(name: $name, age: $age, registered: $registered, address: $address, birthday: $birthday, balance: $balance, someObject: $someObject, someDynamic: $someDynamic)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!identical(other.registered, registered) && other.registered != registered) return false;
    if (!identical(other.address, address) && other.address != address) return false;
    if (!identical(other.birthday, birthday) && other.birthday != birthday) return false;
    if (!identical(other.balance, balance) && other.balance != balance) return false;
    if (!identical(other.someObject, someObject) && other.someObject != someObject) return false;
    if (!identical(other.someDynamic, someDynamic) && other.someDynamic != someDynamic) return false;
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
      birthday,
      balance,
      someObject,
      someDynamic,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    String name,
    int? age,
    bool registered,
    Address? address,
    DateTime? birthday,
    BigInt? balance,
    Object? someObject,
    dynamic someDynamic,
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
    Object? birthday = genq,
    Object? balance = genq,
    Object? someObject = genq,
    Object? someDynamic = genq,
  }) {
    return User(
      name: name == genq ? value.name : name as String,
      age: age == genq ? value.age : age as int?,
      registered: registered == genq ? value.registered : registered as bool,
      address: address == genq ? value.address : address as Address?,
      birthday: birthday == genq ? value.birthday : birthday as DateTime?,
      balance: balance == genq ? value.balance : balance as BigInt?,
      someObject: someObject == genq ? value.someObject : someObject as Object?,
      someDynamic: someDynamic == genq ? value.someDynamic : someDynamic as dynamic,
    );
  }
}

User $UserFromJson(Map<String, dynamic> json) {
  return User(
    name: json['full_name'] as String,
    age: json['age'] == null ? null : json['age'] as int?,
    registered: json['registered'] as bool,
    address: json['address'] == null ? null : $AddressFromJson(json['address']),
    birthday: json['birthday'] == null ? null : DateTime.parse(json['birthday']),
    balance: json['balance'] == null ? null : BigInt.parse(json['balance']),
    someObject: json['someObject'] == null ? null : json['someObject'] as Object?,
    someDynamic: json['someDynamic'] as dynamic,
  );
}

Map<String, dynamic> $UserToJson(User obj) {
  return {
    'full_name': obj.name,
    'age': obj.age == null ? null : obj.age!,
    'registered': obj.registered,
    'address': obj.address == null ? null : $AddressToJson(obj.address!),
    'birthday': obj.birthday == null ? null : obj.birthday!.toIso8601String(),
    'balance': obj.balance == null ? null : obj.balance!.toString(),
    'someObject': obj.someObject == null ? null : obj.someObject!,
    'someDynamic': obj.someDynamic,
  };
}

mixin _$Address {
  String get street => throw UnimplementedError();
  String get city => throw UnimplementedError();
  String get state => throw UnimplementedError();
  String get zip => throw UnimplementedError();

  $AddressCopyWith get copyWith => throw UnimplementedError();
}

class _Address implements Address {
  @override
  final String street;

  @override
  final String city;

  @override
  final String state;

  @override
  final String zip;

  _Address({
    required this.street,
    required this.city,
    required this.state,
    required this.zip,
  });

  @override
  $AddressCopyWith get copyWith => _$AddressCopyWithImpl(this);

  @override
  String toString() {
    return "Address(street: $street, city: $city, state: $state, zip: $zip)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Address) return false;
    if (!identical(other.street, street) && other.street != street) return false;
    if (!identical(other.city, city) && other.city != city) return false;
    if (!identical(other.state, state) && other.state != state) return false;
    if (!identical(other.zip, zip) && other.zip != zip) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      street,
      city,
      state,
      zip,
    );
  }
}

abstract class $AddressCopyWith {
  Address call({
    String street,
    String city,
    String state,
    String zip,
  });
}

class _$AddressCopyWithImpl implements $AddressCopyWith {
  final _$Address value;

  _$AddressCopyWithImpl(this.value);

  @override
  Address call({
    Object? street = genq,
    Object? city = genq,
    Object? state = genq,
    Object? zip = genq,
  }) {
    return Address(
      street: street == genq ? value.street : street as String,
      city: city == genq ? value.city : city as String,
      state: state == genq ? value.state : state as String,
      zip: zip == genq ? value.zip : zip as String,
    );
  }
}

Address $AddressFromJson(Map<String, dynamic> json) {
  return Address(
    street: json['street'] as String,
    city: json['city'] as String,
    state: json['state'] as String,
    zip: json['zip'] as String,
  );
}

Map<String, dynamic> $AddressToJson(Address obj) {
  return {
    'street': obj.street,
    'city': obj.city,
    'state': obj.state,
    'zip': obj.zip,
  };
}