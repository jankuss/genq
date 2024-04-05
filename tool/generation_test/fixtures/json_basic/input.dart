import 'package:genq/genq.dart';

part 'input.genq.dart';

@genq
class User with _$User {
  factory User({
    @JsonKey(name: 'full_name') required String name,
    required int? age,
    required bool registered,
    required Address? address,
  }) = _User;

  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}

@genq
class Address with _$Address {
  factory Address({
    required String street,
    required String city,
    required String state,
    required String zip,
  }) = _Address;

  factory Address.fromJson(Map<String, dynamic> json) => _$AddressFromJson(json);
}
