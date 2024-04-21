part of 'input.dart';

mixin _$User {
  List<Address> get addressesList => throw UnimplementedError();
  List<Address?> get addressesListNullable => throw UnimplementedError();
  Set<Address> get addressesSet => throw UnimplementedError();
  Set<Address?> get addressesSetNullable => throw UnimplementedError();
  Map<String, Address> get addressesMap => throw UnimplementedError();
  Map<Uri, Address> get addressesUriMap => throw UnimplementedError();
  Map<String, Address?> get addressesMapNullableValue => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  @override
  final List<Address> addressesList;

  @override
  final List<Address?> addressesListNullable;

  @override
  final Set<Address> addressesSet;

  @override
  final Set<Address?> addressesSetNullable;

  @override
  final Map<String, Address> addressesMap;

  @override
  final Map<Uri, Address> addressesUriMap;

  @override
  final Map<String, Address?> addressesMapNullableValue;

  _User({
    required this.addressesList,
    required this.addressesListNullable,
    required this.addressesSet,
    required this.addressesSetNullable,
    required this.addressesMap,
    required this.addressesUriMap,
    required this.addressesMapNullableValue,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(addressesList: $addressesList, addressesListNullable: $addressesListNullable, addressesSet: $addressesSet, addressesSetNullable: $addressesSetNullable, addressesMap: $addressesMap, addressesUriMap: $addressesUriMap, addressesMapNullableValue: $addressesMapNullableValue)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!const DeepCollectionEquality().equals(other.addressesList, addressesList)) return false;
    if (!const DeepCollectionEquality().equals(other.addressesListNullable, addressesListNullable)) return false;
    if (!const DeepCollectionEquality().equals(other.addressesSet, addressesSet)) return false;
    if (!const DeepCollectionEquality().equals(other.addressesSetNullable, addressesSetNullable)) return false;
    if (!const DeepCollectionEquality().equals(other.addressesMap, addressesMap)) return false;
    if (!const DeepCollectionEquality().equals(other.addressesUriMap, addressesUriMap)) return false;
    if (!const DeepCollectionEquality().equals(other.addressesMapNullableValue, addressesMapNullableValue)) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      addressesList,
      addressesListNullable,
      addressesSet,
      addressesSetNullable,
      addressesMap,
      addressesUriMap,
      addressesMapNullableValue,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    List<Address> addressesList,
    List<Address?> addressesListNullable,
    Set<Address> addressesSet,
    Set<Address?> addressesSetNullable,
    Map<String, Address> addressesMap,
    Map<Uri, Address> addressesUriMap,
    Map<String, Address?> addressesMapNullableValue,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User value;

  _$UserCopyWithImpl(this.value);

  @override
  User call({
    Object? addressesList = genq,
    Object? addressesListNullable = genq,
    Object? addressesSet = genq,
    Object? addressesSetNullable = genq,
    Object? addressesMap = genq,
    Object? addressesUriMap = genq,
    Object? addressesMapNullableValue = genq,
  }) {
    return User(
      addressesList: addressesList == genq ? value.addressesList : addressesList as List<Address>,
      addressesListNullable: addressesListNullable == genq ? value.addressesListNullable : addressesListNullable as List<Address?>,
      addressesSet: addressesSet == genq ? value.addressesSet : addressesSet as Set<Address>,
      addressesSetNullable: addressesSetNullable == genq ? value.addressesSetNullable : addressesSetNullable as Set<Address?>,
      addressesMap: addressesMap == genq ? value.addressesMap : addressesMap as Map<String, Address>,
      addressesUriMap: addressesUriMap == genq ? value.addressesUriMap : addressesUriMap as Map<Uri, Address>,
      addressesMapNullableValue: addressesMapNullableValue == genq ? value.addressesMapNullableValue : addressesMapNullableValue as Map<String, Address?>,
    );
  }
}

User $UserFromJson(Map<String, dynamic> json) {
  return User(
    addressesList: List.of(json['addressesList']).map((e) => $AddressFromJson(e)).toList(),
    addressesListNullable: List.of(json['addressesListNullable']).map((e) => e == null ? null : $AddressFromJson(e)).toList(),
    addressesSet: Set.of(json['addressesSet']).map((e) => $AddressFromJson(e)).toSet(),
    addressesSetNullable: Set.of(json['addressesSetNullable']).map((e) => e == null ? null : $AddressFromJson(e)).toSet(),
    addressesMap: Map.of(json['addressesMap']).map((key, value) => MapEntry(key as String, $AddressFromJson(value))),
    addressesUriMap: Map.of(json['addressesUriMap']).map((key, value) => MapEntry(Uri.parse(key), $AddressFromJson(value))),
    addressesMapNullableValue: Map.of(json['addressesMapNullableValue']).map((key, value) => MapEntry(key as String, value == null ? null : $AddressFromJson(value))),
  );
}

Map<String, dynamic> $UserToJson(User obj) {
  return {
    'addressesList': obj.addressesList.map((e) => $AddressToJson(e)).toList(),
    'addressesListNullable': obj.addressesListNullable.map((e) => e == null ? null : $AddressToJson(e)).toList(),
    'addressesSet': obj.addressesSet.map((e) => $AddressToJson(e)).toList(),
    'addressesSetNullable': obj.addressesSetNullable.map((e) => e == null ? null : $AddressToJson(e)).toList(),
    'addressesMap': Map.of(obj.addressesMap).map((key, value) => MapEntry(key, $AddressToJson(value))),
    'addressesUriMap': Map.of(obj.addressesUriMap).map((key, value) => MapEntry(key.toString(), $AddressToJson(value))),
    'addressesMapNullableValue': Map.of(obj.addressesMapNullableValue).map((key, value) => MapEntry(key, value == null ? null : $AddressToJson(value))),
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