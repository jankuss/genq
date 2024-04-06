part of 'collections.dart';

mixin _$CollectionTests {
  List<User> get users => throw UnimplementedError();
  Map<String, User> get userMap => throw UnimplementedError();
  Set<User> get userSet => throw UnimplementedError();

  $CollectionTestsCopyWith get copyWith => throw UnimplementedError();
}

class _CollectionTests implements CollectionTests {
  @override
  final List<User> users;

  @override
  final Map<String, User> userMap;

  @override
  final Set<User> userSet;

  _CollectionTests({
    required this.users,
    required this.userMap,
    required this.userSet,
  });

  @override
  $CollectionTestsCopyWith get copyWith => _$CollectionTestsCopyWithImpl(this);

  @override
  String toString() {
    return "CollectionTests(users: $users, userMap: $userMap, userSet: $userSet)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! CollectionTests) return false;
    if (!const DeepCollectionEquality().equals(other.users, users)) return false;
    if (!const DeepCollectionEquality().equals(other.userMap, userMap)) return false;
    if (!const DeepCollectionEquality().equals(other.userSet, userSet)) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      users,
      userMap,
      userSet,
    );
  }
}

abstract class $CollectionTestsCopyWith {
  CollectionTests call({
    List<User> users,
    Map<String, User> userMap,
    Set<User> userSet,
  });
}

class _$CollectionTestsCopyWithImpl implements $CollectionTestsCopyWith {
  final _$CollectionTests value;

  _$CollectionTestsCopyWithImpl(this.value);

  @override
  CollectionTests call({
    Object? users = genq,
    Object? userMap = genq,
    Object? userSet = genq,
  }) {
    return CollectionTests(
      users: users == genq ? value.users : users as List<User>,
      userMap: userMap == genq ? value.userMap : userMap as Map<String, User>,
      userSet: userSet == genq ? value.userSet : userSet as Set<User>,
    );
  }
}