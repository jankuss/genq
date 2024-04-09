import 'dart:convert';

import 'package:genq/genq.dart';
import 'package:http/io_client.dart';

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
  final http = IOClient();

  final value = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/users'));
  final json = List.of(jsonDecode(value.body)).map((e) => User.fromJson(e)).toList();

  for (final user in json) {
    print(user);
  }
}
