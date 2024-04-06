part of 'json_user.dart';

mixin _$JsonUser {
  String get name => throw UnimplementedError();
  int? get age => throw UnimplementedError();
  bool get registered => throw UnimplementedError();
  JsonUserStatus? get status => throw UnimplementedError();
  JsonAddress? get address => throw UnimplementedError();

  $JsonUserCopyWith get copyWith => throw UnimplementedError();
}

class _JsonUser implements JsonUser {
  @override
  final String name;

  @override
  final int? age;

  @override
  final bool registered;

  @override
  final JsonUserStatus? status;

  @override
  final JsonAddress? address;

  _JsonUser({
    required this.name,
    required this.age,
    required this.registered,
    required this.status,
    this.address,
  });

  @override
  $JsonUserCopyWith get copyWith => _$JsonUserCopyWithImpl(this);

  @override
  String toString() {
    return "JsonUser(name: $name, age: $age, registered: $registered, status: $status, address: $address)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! JsonUser) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!identical(other.registered, registered) && other.registered != registered) return false;
    if (!identical(other.status, status) && other.status != status) return false;
    if (!identical(other.address, address) && other.address != address) return false;
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
      address,
    );
  }
}

abstract class $JsonUserCopyWith {
  JsonUser call({
    String name,
    int? age,
    bool registered,
    JsonUserStatus? status,
    JsonAddress? address,
  });
}

class _$JsonUserCopyWithImpl implements $JsonUserCopyWith {
  final _$JsonUser value;

  _$JsonUserCopyWithImpl(this.value);

  @override
  JsonUser call({
    Object? name = genq,
    Object? age = genq,
    Object? registered = genq,
    Object? status = genq,
    Object? address = genq,
  }) {
    return JsonUser(
      name: name == genq ? value.name : name as String,
      age: age == genq ? value.age : age as int?,
      registered: registered == genq ? value.registered : registered as bool,
      status: status == genq ? value.status : status as JsonUserStatus?,
      address: address == genq ? value.address : address as JsonAddress?,
    );
  }
}

JsonUser $JsonUserFromJson(Map<String, dynamic> json) {
  return JsonUser(
    name: json['full_name'] as String,
    age: json['age'] == null ? null : (json['age'] as int?),
    registered: json['registered'] as bool,
    status: json['status'] == null ? null : ($JsonUserStatusFromJson(json['status']) as JsonUserStatus?),
    address: json['address'] == null ? null : ($JsonAddressFromJson(json['address']) as JsonAddress?),
  );
}

Map<String, dynamic> $JsonUserToJson(JsonUser obj) {
  return {
    'full_name': obj.name,
    'age': obj.age == null ? null : obj.age!,
    'registered': obj.registered,
    'status': obj.status == null ? null : $JsonUserStatusToJson(obj.status!),
    'address': obj.address == null ? null : $JsonAddressToJson(obj.address!),
  };
}