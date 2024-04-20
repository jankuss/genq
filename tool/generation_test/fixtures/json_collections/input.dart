import 'package:genq/genq.dart';

part 'input.genq.dart';

@Genq(json: true)
class User with _$User {
  factory User({
    required List<Address> addressesList,
    required Set<Address> addressesSet,
  }) = _User;

  factory User.fromJson(Map<String, dynamic> json) => $UserFromJson(json);
}

@Genq(json: true)
class Address with _$Address {
  factory Address({
    required String street,
  }) = _Address;

  factory Address.fromJson(Map<String, dynamic> json) => $AddressFromJson(json);
}
