import 'package:genq/genq.dart';

part 'input.genq.dart';

@Genq(json: true)
class User with _$User {
  factory User({
    @JsonKey(name: 'full_name') required String name,
    required int? age,
    required bool registered,
    required Address? address,
    required DateTime? birthday,
    required BigInt? balance,
    required Object? someObject,
    required dynamic someDynamic,
    required double someDoubleValue,
    required double? someDoubleValueNullable,
  }) = _User;

  factory User.fromJson(Map<String, dynamic> json) => $UserFromJson(json);
}

@Genq(json: true)
class Address with _$Address {
  factory Address({
    required String street,
    required String city,
    required String state,
    required String zip,
  }) = _Address;

  factory Address.fromJson(Map<String, dynamic> json) => $AddressFromJson(json);
}
