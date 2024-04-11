part of 'input.dart';

mixin _$User {
  String get name => throw UnimplementedError();
  int? get age => throw UnimplementedError();
  bool get registered => throw UnimplementedError();
  UserStatus? get status => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  @override
  final String name;

  @override
  final int? age;

  @override
  final bool registered;

  @override
  final UserStatus? status;

  _User({
    required this.name,
    required this.age,
    required this.registered,
    required this.status,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(name: $name, age: $age, registered: $registered, status: $status)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!identical(other.registered, registered) && other.registered != registered) return false;
    if (!identical(other.status, status) && other.status != status) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
      age,
      registered,
      status,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    String name,
    int? age,
    bool registered,
    UserStatus? status,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User value;

  _$UserCopyWithImpl(this.value);

  @override
  User call({
    Object? name = genq,
    Object? age = genq,
    Object? registered = genq,
    Object? status = genq,
  }) {
    return User(
      name: name == genq ? value.name : name as String,
      age: age == genq ? value.age : age as int?,
      registered: registered == genq ? value.registered : registered as bool,
      status: status == genq ? value.status : status as UserStatus?,
    );
  }
}

User $UserFromJson(Map<String, dynamic> json) {
  return User(
    name: json['full_name'] as String,
    age: json['age'] == null ? null : json['age'] as int?,
    registered: json['registered'] as bool,
    status: json['status'] == null ? null : $UserStatusFromJson(json['status']),
  );
}

Map<String, dynamic> $UserToJson(User obj) {
  return {
    'full_name': obj.name,
    'age': obj.age == null ? null : obj.age!,
    'registered': obj.registered,
    'status': obj.status == null ? null : $UserStatusToJson(obj.status!),
  };
}

UserStatus $UserStatusFromJson(Object json, [UserStatus? unknownEnumValue]) {
  switch (json) {
    case "REGISTERED":
      return UserStatus.registered;
    case "UNREGISTERED":
      return UserStatus.unregistered;
    case 0:
      return UserStatus.pending;
    case 1:
      return UserStatus.blocked;
    default:
      if (unknownEnumValue != null) {
        return unknownEnumValue;
      } else {
        throw UnsupportedError('The value $json is not a valid value for enum UserStatus');
      }
  }
}

Object $UserStatusToJson(UserStatus value) {
  switch (value) {
    case UserStatus.registered:
      return "REGISTERED";
    case UserStatus.unregistered:
      return "UNREGISTERED";
    case UserStatus.pending:
      return 0;
    case UserStatus.blocked:
      return 1;
    default:
      throw UnsupportedError('Could not map $value to a JSON value');
  }
}