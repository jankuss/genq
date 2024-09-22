part of 'input.dart';

mixin _$User {
  List<String> get names => throw UnimplementedError();
  Map<String, dynamic> get properties => throw UnimplementedError();
  Set<int> get ages => throw UnimplementedError();

  $UserCopyWith get copyWith => throw UnimplementedError();
}

class _User implements User {
  @override
  final List<String> names;

  @override
  final Map<String, dynamic> properties;

  @override
  final Set<int> ages;

  _User({
    required this.names,
    required this.properties,
    required this.ages,
  });

  @override
  $UserCopyWith get copyWith => _$UserCopyWithImpl(this);

  @override
  String toString() {
    return "User(names: $names, properties: $properties, ages: $ages)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! User) return false;
    if (!const DeepCollectionEquality().equals(other.names, names)) return false;
    if (!const DeepCollectionEquality().equals(other.properties, properties)) return false;
    if (!const DeepCollectionEquality().equals(other.ages, ages)) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      names,
      properties,
      ages,
    );
  }
}

abstract class $UserCopyWith {
  User call({
    List<String> names,
    Map<String, dynamic> properties,
    Set<int> ages,
  });
}

class _$UserCopyWithImpl implements $UserCopyWith {
  final _$User __value;

  _$UserCopyWithImpl(this.__value);

  @override
  User call({
    Object? names = genq,
    Object? properties = genq,
    Object? ages = genq,
  }) {
    return User(
      names: names == genq ? __value.names : names as List<String>,
      properties: properties == genq ? __value.properties : properties as Map<String, dynamic>,
      ages: ages == genq ? __value.ages : ages as Set<int>,
    );
  }
}