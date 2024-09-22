part of 'input.dart';

mixin _$User {
  T get data => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  @override
  final T data;

  _User({
    required this.data,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(data: $data)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!identical(other.data, data) && other.data != data) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      data,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    T data,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User __value;

  _$UserCopyWithImpl(this.__value);

  @override
  User call({
    Object? data = genq,
  }) {
    return User(
      data: data == genq ? __value.data : data as T,
    );
  }
}