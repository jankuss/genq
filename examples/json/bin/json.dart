import 'dart:convert';
import 'dart:io';

import 'package:genq/genq.dart';

part 'json.genq.dart';

@Genq(json: true)
class User with _$User {
  factory User({
    required int id,
    required String name,
    required String username,
    required String email,
    required Address address,
    required String phone,
    required String website,
  }) = _User;

  factory User.fromJson(Map<String, dynamic> json) => $UserFromJson(json);
}

@Genq(json: true)
class Address with _$Address {
  factory Address({
    required String street,
    required String suite,
    required String city,
    required String zipcode,
    required Geo geo,
  }) = _Address;

  factory Address.fromJson(Map<String, dynamic> json) => $AddressFromJson(json);
}

@Genq(json: true)
class Geo with _$Geo {
  factory Geo({
    required String lat,
    required String lng,
  }) = _Geo;

  factory Geo.fromJson(Map<String, dynamic> json) => $GeoFromJson(json);
}

void main(List<String> arguments) async {
  final value = File('data.json').readAsStringSync();
  final json = List.of(jsonDecode(value)).map((e) => User.fromJson(e)).toList();

  for (final user in json) {
    print(user);
  }
}
