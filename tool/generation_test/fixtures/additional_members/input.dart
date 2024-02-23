import 'package:genq/genq.dart';

part "input.genq.dart";

@genq
class User with _$User {
  User._();
  factory User({
    required String name,
    required int? age,
    required bool registered,
    required Address address,
  }) = _User;

  String _getIntroduction() {
    return "My name is $name and I'm $age years old.";
  }

  void printAddress() {
    print("${_getIntroduction()} I live at ${address.street} in ${address.zipCode} ${address.city}");
  }
}

@genq
class Address with _$Address {
  factory Address({
    required String street,
    required int zipCode,
    required String city,
  }) = _Address;
}


