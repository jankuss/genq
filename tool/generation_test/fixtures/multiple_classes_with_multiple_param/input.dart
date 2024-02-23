import 'package:genq/genq.dart';

part "input.genq.dart";

@genq
class User with _$User {
  factory User({
    required String name,
    required int? age,
    required bool registered,
    required Address address,
  }) = _User;
}

@genq
class Address with _$Address {
  factory Address({
    required String street,
    required int zipCode,
    required String city,
  }) = _Address;
}


