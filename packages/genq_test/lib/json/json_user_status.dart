import 'package:genq/genq.dart';

part 'json_user_status.genq.dart';

@GenqJsonEnum()
enum JsonUserStatus {
  @JsonValue('active')
  active,
  @JsonValue('INACTIVE')
  inactive,
}
