part of 'input.dart';

mixin _$User {
  List<Address> get addressesList => throw UnimplementedError();
  Set<Address> get addressesSet => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  @override
  final List<Address> addressesList;

  @override
  final Set<Address> addressesSet;

  _User({
    required this.addressesList,
    required this.addressesSet,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(addressesList: $addressesList, addressesSet: $addressesSet)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!const DeepCollectionEquality().equals(other.addressesList, addressesList)) return false;
    if (!const DeepCollectionEquality().equals(other.addressesSet, addressesSet)) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      addressesList,
      addressesSet,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    List<Address> addressesList,
    Set<Address> addressesSet,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User value;

  _$UserCopyWithImpl(this.value);

  @override
  User call({
    Object? addressesList = genq,
    Object? addressesSet = genq,
  }) {
    return User(
      addressesList: addressesList == genq ? value.addressesList : addressesList as List<Address>,
      addressesSet: addressesSet == genq ? value.addressesSet : addressesSet as Set<Address>,
    );
  }
}

User $UserFromJson(Map<String, dynamic> json) {
  return User(
    addressesList: List.of(json['addressesList']).map((e) => $AddressFromJson(e)).toList(),
    addressesSet: Set.of(json['addressesSet']).map((e) => $AddressFromJson(e)).toSet(),
  );
}

Map<String, dynamic> $UserToJson(User obj) {
  return {
    'addressesList': obj.addressesList.map((e) => $AddressToJson(e)).toList(),
    'addressesSet': obj.addressesSet.map((e) => $AddressToJson(e)).toSet(),
  };
}

mixin _$Address {
  String get street => throw UnimplementedError();

  $AddressCopyWith get copyWith => throw UnimplementedError();
}

class _Address implements Address {
  @override
  final String street;

  _Address({
    required this.street,
  });

  @override
  $AddressCopyWith get copyWith => _$AddressCopyWithImpl(this);

  @override
  String toString() {
    return "Address(street: $street)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Address) return false;
    if (!identical(other.street, street) && other.street != street) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      street,
    );
  }
}

abstract class $AddressCopyWith {
  Address call({
    String street,
  });
}

class _$AddressCopyWithImpl implements $AddressCopyWith {
  final _$Address value;

  _$AddressCopyWithImpl(this.value);

  @override
  Address call({
    Object? street = genq,
  }) {
    return Address(
      street: street == genq ? value.street : street as String,
    );
  }
}

Address $AddressFromJson(Map<String, dynamic> json) {
  return Address(
    street: json['street'] as String,
  );
}

Map<String, dynamic> $AddressToJson(Address obj) {
  return {
    'street': obj.street,
  };
}