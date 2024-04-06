part of 'json_user_status.dart';

JsonUserStatus $JsonUserStatusFromJson(Object json) {
  switch (json) {
    case 'active':
      return JsonUserStatus.active;
    case 'INACTIVE':
      return JsonUserStatus.inactive;
    default:
      throw InvalidJsonEnumValueError(json, JsonUserStatus);
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