import 'package:genq/genq.dart';

part 'address.genq.dart';

@genq
class Address with _$Address {
  factory Address({
    required String street,
    required int zipCode,
    required String city,
  }) = _Address;
}
