import 'package:genq/genq.dart';

part 'input.genq.dart';

@genq
class User with _$User {
  factory User({
    @JsonKey(name: 'full_name') required String name,
    required int? age,
    required bool registered,
    required UserStatus? status,
  }) = _User;

  factory User.fromJson(Map<String, dynamic> json) => $UserFromJson(json);
}

@genqJsonEnum
enum UserStatus {
  registered,
  unregistered,
  pending,
  blocked,
}
