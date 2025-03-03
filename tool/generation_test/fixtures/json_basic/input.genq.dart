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
  double get someDoubleValue => throw UnimplementedError();
  double? get someDoubleValueNullable => throw UnimplementedError();

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

  @override
  final double someDoubleValue;

  @override
  final double? someDoubleValueNullable;

  _User({
    required this.name,
    required this.age,
    required this.registered,
    required this.address,
    required this.birthday,
    required this.balance,
    required this.someObject,
    required this.someDynamic,
    required this.someDoubleValue,
    required this.someDoubleValueNullable,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(name: $name, age: $age, registered: $registered, address: $address, birthday: $birthday, balance: $balance, someObject: $someObject, someDynamic: $someDynamic, someDoubleValue: $someDoubleValue, someDoubleValueNullable: $someDoubleValueNullable)";
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
    if (!identical(other.someDoubleValue, someDoubleValue) && other.someDoubleValue != someDoubleValue) return false;
    if (!identical(other.someDoubleValueNullable, someDoubleValueNullable) && other.someDoubleValueNullable != someDoubleValueNullable) return false;
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
      someDoubleValue,
      someDoubleValueNullable,
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
    double someDoubleValue,
    double? someDoubleValueNullable,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User __value;

  _$UserCopyWithImpl(this.__value);

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
    Object? someDoubleValue = genq,
    Object? someDoubleValueNullable = genq,
  }) {
    return User(
      name: name == genq ? __value.name : name as String,
      age: age == genq ? __value.age : age as int?,
      registered: registered == genq ? __value.registered : registered as bool,
      address: address == genq ? __value.address : address as Address?,
      birthday: birthday == genq ? __value.birthday : birthday as DateTime?,
      balance: balance == genq ? __value.balance : balance as BigInt?,
      someObject: someObject == genq ? __value.someObject : someObject,
      someDynamic: someDynamic == genq ? __value.someDynamic : someDynamic as dynamic,
      someDoubleValue: someDoubleValue == genq ? __value.someDoubleValue : someDoubleValue as double,
      someDoubleValueNullable: someDoubleValueNullable == genq ? __value.someDoubleValueNullable : someDoubleValueNullable as double?,
    );
  }
}

User $UserFromJson(Map<String, dynamic> json) {
  return User(
    name: json['full_name'] as String,
    age: json['age'] == null ? null : (json['age'] as num).toInt(),
    registered: json['registered'] as bool,
    address: json['address'] == null ? null : $AddressFromJson(json['address']),
    birthday: json['birthday'] == null ? null : DateTime.parse(json['birthday']),
    balance: json['balance'] == null ? null : BigInt.parse(json['balance']),
    someObject: json['someObject'] == null ? null : json['someObject'] as Object?,
    someDynamic: json['someDynamic'] as dynamic,
    someDoubleValue: (json['someDoubleValue'] as num).toDouble(),
    someDoubleValueNullable: json['someDoubleValueNullable'] == null ? null : (json['someDoubleValueNullable'] as num).toDouble(),
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
    'someDoubleValue': obj.someDoubleValue,
    'someDoubleValueNullable': obj.someDoubleValueNullable == null ? null : obj.someDoubleValueNullable!,
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
  final _$Address __value;

  _$AddressCopyWithImpl(this.__value);

  @override
  Address call({
    Object? street = genq,
    Object? city = genq,
    Object? state = genq,
    Object? zip = genq,
  }) {
    return Address(
      street: street == genq ? __value.street : street as String,
      city: city == genq ? __value.city : city as String,
      state: state == genq ? __value.state : state as String,
      zip: zip == genq ? __value.zip : zip as String,
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