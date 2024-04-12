import 'package:genq/genq.dart';
import 'package:genq_test/address.dart';

part 'user.genq.dart';

@genq
class User with _$User {
  factory User({
    required String username,
    required String email,
    required int? age,
    Address? address,
  }) = _User;
}
