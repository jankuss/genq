part of 'json_account.dart';

mixin _$JsonAccount {
  String get email => throw UnimplementedError();
  JsonStatus get status => throw UnimplementedError();

  $JsonAccountCopyWith get copyWith => throw UnimplementedError();
}

class _JsonAccount implements JsonAccount {
  @override
  final String email;

  @override
  final JsonStatus status;

  _JsonAccount({
    required this.email,
    required this.status,
  });

  @override
  $JsonAccountCopyWith get copyWith => _$JsonAccountCopyWithImpl(this);

  @override
  String toString() {
    return "JsonAccount(email: $email, status: $status)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! JsonAccount) return false;
    if (!identical(other.email, email) && other.email != email) return false;
    if (!identical(other.status, status) && other.status != status) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      email,
      status,
    );
  }
}

abstract class $JsonAccountCopyWith {
  JsonAccount call({
    String email,
    JsonStatus status,
  });
}

class _$JsonAccountCopyWithImpl implements $JsonAccountCopyWith {
  final _$JsonAccount value;

  _$JsonAccountCopyWithImpl(this.value);

  @override
  JsonAccount call({
    Object? email = genq,
    Object? status = genq,
  }) {
    return JsonAccount(
      email: email == genq ? value.email : email as String,
      status: status == genq ? value.status : status as JsonStatus,
    );
  }
}

JsonAccount $JsonAccountFromJson(Map<String, dynamic> json) {
  return JsonAccount(
    email: json['email'] as String,
    status: $JsonStatusFromJson(json['status'], JsonStatus.unknown) as JsonStatus,
  );
}

Map<String, dynamic> $JsonAccountToJson(JsonAccount obj) {
  return {
    'email': obj.email,
    'status': $JsonStatusToJson(obj.status),
  };
}

JsonStatus $JsonStatusFromJson(Object json, [JsonStatus? unknownEnumValue]) {
  switch (json) {
    case "active":
      return JsonStatus.active;
    case "inactive":
      return JsonStatus.inactive;
    case "unknown":
      return JsonStatus.unknown;
    default:
      if (unknownEnumValue != null) {
        return unknownEnumValue;
      } else {
        throw UnsupportedError('The value $json is not a valid value for enum JsonStatus');
      }
  }
}

Object $JsonStatusToJson(JsonStatus value) {
  switch (value) {
    case JsonStatus.active:
      return "active";
    case JsonStatus.inactive:
      return "inactive";
    case JsonStatus.unknown:
      return "unknown";
    default:
      throw UnsupportedError('Could not map $value to a JSON value');
  }
}