part of 'json_user_status.dart';

JsonUserStatus $JsonUserStatusFromJson(Object json, [JsonUserStatus? unknownEnumValue]) {
  switch (json) {
    case 'active':
      return JsonUserStatus.active;
    case 'INACTIVE':
      return JsonUserStatus.inactive;
    default:
      if (unknownEnumValue != null) {
        return unknownEnumValue;
      } else {
        throw UnsupportedError('The value $json is not a valid value for enum JsonUserStatus');
      }
  }
}

Object $JsonUserStatusToJson(JsonUserStatus value) {
  switch (value) {
    case JsonUserStatus.active:
      return 'active';
    case JsonUserStatus.inactive:
      return 'INACTIVE';
    default:
      throw UnsupportedError('Could not map $value to a JSON value');
  }
}