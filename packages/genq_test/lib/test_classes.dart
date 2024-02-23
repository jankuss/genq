import 'package:genq/genq.dart';

part 'test_classes.genq.dart';

@genq
class User with _$User {
  factory User({
    required String username,
    required String email,
    required int? age,
    Address? address,
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

@genq 
class CollectionTests with _$CollectionTests {
  factory CollectionTests({
    required List<User> users,
    required Map<String, User> userMap,
    required Set<User> userSet,
  }) = _CollectionTests;
}

