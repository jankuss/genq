part of 'input.dart';

mixin _$SomeState {

  $SomeStateCopyWith get copyWith => throw UnimplementedError();
}

class _SomeState implements SomeState {
  _SomeState();

  @override
  $SomeStateCopyWith get copyWith => _$SomeStateCopyWithImpl(this);

  @override
  String toString() {
    return "SomeState()";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! SomeState) return false;
    return true;
  }

  @override
  int get hashCode {
    return runtimeType.hashCode;
  }
}

abstract class $SomeStateCopyWith {
  SomeState call();
}

class _$SomeStateCopyWithImpl implements $SomeStateCopyWith {
  final _$SomeState value;

  _$SomeStateCopyWithImpl(this.value);

  @override
  SomeState call() {
    return SomeState();
  }
}

mixin _$Initial {

  $InitialCopyWith get copyWith => throw UnimplementedError();
}

class Initial with _$Initial implements _SomeState {
  const Initial();

  @override
  $InitialCopyWith get copyWith => _$InitialCopyWithImpl(this);

  @override
  String toString() {
    return "Initial()";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Initial) return false;
    return true;
  }

  @override
  int get hashCode {
    return runtimeType.hashCode;
  }
}

abstract class $InitialCopyWith {
  Initial call();
}

class _$InitialCopyWithImpl implements $InitialCopyWith {
  final _$Initial value;

  _$InitialCopyWithImpl(this.value);

  @override
  Initial call() {
    return Initial();
  }
}
mixin _$Loading {

  $LoadingCopyWith get copyWith => throw UnimplementedError();
}

class Loading with _$Loading implements _SomeState {
  const Loading();

  @override
  $LoadingCopyWith get copyWith => _$LoadingCopyWithImpl(this);

  @override
  String toString() {
    return "Loading()";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Loading) return false;
    return true;
  }

  @override
  int get hashCode {
    return runtimeType.hashCode;
  }
}

abstract class $LoadingCopyWith {
  Loading call();
}

class _$LoadingCopyWithImpl implements $LoadingCopyWith {
  final _$Loading value;

  _$LoadingCopyWithImpl(this.value);

  @override
  Loading call() {
    return Loading();
  }
}
mixin _$Loaded {
  String get data => throw UnimplementedError();

  $LoadedCopyWith get copyWith => throw UnimplementedError();
}

class Loaded with _$Loaded implements _SomeState {
  @override
  final String data;

  const Loaded({
    required this.data,
  });

  @override
  $LoadedCopyWith get copyWith => _$LoadedCopyWithImpl(this);

  @override
  String toString() {
    return "Loaded(data: $data)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Loaded) return false;
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

abstract class $LoadedCopyWith {
  Loaded call({
    String data,
  });
}

class _$LoadedCopyWithImpl implements $LoadedCopyWith {
  final _$Loaded value;

  _$LoadedCopyWithImpl(this.value);

  @override
  Loaded call({
    Object? data = genq,
  }) {
    return Loaded(
      data: data == genq ? value.data : data as String,
    );
  }
}
mixin _$Error {
  String get message => throw UnimplementedError();

  $ErrorCopyWith get copyWith => throw UnimplementedError();
}

class Error with _$Error implements _SomeState {
  @override
  final String message;

  const Error({
    required this.message,
  });

  @override
  $ErrorCopyWith get copyWith => _$ErrorCopyWithImpl(this);

  @override
  String toString() {
    return "Error(message: $message)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! Error) return false;
    if (!identical(other.message, message) && other.message != message) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      message,
    );
  }
}

abstract class $ErrorCopyWith {
  Error call({
    String message,
  });
}

class _$ErrorCopyWithImpl implements $ErrorCopyWith {
  final _$Error value;

  _$ErrorCopyWithImpl(this.value);

  @override
  Error call({
    Object? message = genq,
  }) {
    return Error(
      message: message == genq ? value.message : message as String,
    );
  }
}
