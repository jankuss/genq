part of 'address.dart';

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