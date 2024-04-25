import 'package:genq/genq.dart';

part 'input.genq.dart';

@genq
class User with _$User {
  factory User({
    @Default('John Doe')
    String name,
    @Default(18)
    int age,
    @Default(false)
    bool registered,
    @Default(const Address(street: 'Main St', zipCode: 12345, city: 'City'))
    Address address,
    @Default(0.0)
    double balance,
  }) = _User;
}

@genq
class Address with _$Address {
  const factory Address({
    required String? street,
    required int? zipCode,
    required String? city,
  }) = _Address;
}
