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

@JsonEnum()
enum UserStatus {
  @JsonValue("REGISTERED")
  registered,
  @JsonValue("UNREGISTERED")
  unregistered,
  @JsonValue(0)
  pending,
  @JsonValue(1)
  blocked,
}
