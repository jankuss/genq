import 'package:genq/genq.dart';
import 'package:genq_test/json/json_address.dart';
import 'package:genq_test/json/json_user_status.dart';

part 'json_user.genq.dart';

@Genq(json: true)
class JsonUser with _$JsonUser {
  factory JsonUser({
    @JsonKey(name: 'full_name') required String name,
    required int? age,
    required bool registered,
    required JsonUserStatus? status,
    JsonAddress? address,
  }) = _JsonUser;
}
