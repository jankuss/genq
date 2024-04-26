part of 'input.dart';

mixin _$SomeState {
  dynamic get copyWith => throw UnimplementedError();
}

class Initial implements SomeState {
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
  final Initial value;

  _$InitialCopyWithImpl(this.value);

  @override
  Initial call() {
    return Initial();
  }
}
class Loading implements SomeState {
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
  final Loading value;

  _$LoadingCopyWithImpl(this.value);

  @override
  Loading call() {
    return Loading();
  }
}
class Loaded implements SomeState {
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
  final Loaded value;

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
class Error implements SomeState {
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
  final Error value;

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