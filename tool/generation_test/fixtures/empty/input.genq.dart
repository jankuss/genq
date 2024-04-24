part of 'input.dart';

mixin _$User {

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  _User();

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User()";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    return true;
  }

  @override
  int get hashCode {
    return runtimeType.hashCode;
  }
}

abstract class $UserCopyWith {
  User call();
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User value;

  _$UserCopyWithImpl(this.value);

  @override
  User call() {
    return User();
  }
}
