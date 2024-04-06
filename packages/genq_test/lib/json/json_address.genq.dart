part of 'json_address.dart';

mixin _$JsonAddress {
  String get street => throw UnimplementedError();
  String get city => throw UnimplementedError();
  String get state => throw UnimplementedError();
  String get zip => throw UnimplementedError();

  $JsonAddressCopyWith get copyWith => throw UnimplementedError();
}

class _JsonAddress implements JsonAddress {
  @override
  final String street;

  @override
  final String city;

  @override
  final String state;

  @override
  final String zip;

  _JsonAddress({
    required this.street,
    required this.city,
    required this.state,
    required this.zip,
  });

  @override
  $JsonAddressCopyWith get copyWith => _$JsonAddressCopyWithImpl(this);

  @override
  String toString() {
    return "JsonAddress(street: $street, city: $city, state: $state, zip: $zip)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! JsonAddress) return false;
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

abstract class $JsonAddressCopyWith {
  JsonAddress call({
    String street,
    String city,
    String state,
    String zip,
  });
}

class _$JsonAddressCopyWithImpl implements $JsonAddressCopyWith {
  final _$JsonAddress value;

  _$JsonAddressCopyWithImpl(this.value);

  @override
  JsonAddress call({
    Object? street = genq,
    Object? city = genq,
    Object? state = genq,
    Object? zip = genq,
  }) {
    return JsonAddress(
      street: street == genq ? value.street : street as String,
      city: city == genq ? value.city : city as String,
      state: state == genq ? value.state : state as String,
      zip: zip == genq ? value.zip : zip as String,
    );
  }
}

JsonAddress $JsonAddressFromJson(Map<String, dynamic> json) {
  return JsonAddress(
    street: json['street'] as String,
    city: json['city'] as String,
    state: json['state'] as String,
    zip: json['zip'] as String,
  );
}

Map<String, dynamic> $JsonAddressToJson(JsonAddress obj) {
  return {
    'street': obj.street,
    'city': obj.city,
    'state': obj.state,
    'zip': obj.zip,
  };
}