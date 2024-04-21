import 'package:genq/genq.dart';

part 'input.genq.dart';

@Genq(json: true)
class User with _$User {
  factory User({
    required List<Address> addressesList,
    required List<Address?> addressesListNullable,
    required Set<Address> addressesSet,
    required Set<Address?> addressesSetNullable,
    required Map<String, Address> addressesMap,
    required Map<Uri, Address> addressesUriMap,
    required Map<String, Address?> addressesMapNullableValue,
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
