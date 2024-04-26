import 'package:genq/genq.dart';

part 'input.genq.dart';

sealed class State {}

@genq
class LoadingState extends State with _$LoadingState {
  const factory LoadingState() = _LoadingState;
}

@genq
class SuccessState extends State with _$SuccessState {
  const factory SuccessState({
    required String name,
    required int age,
  }) = _SuccessState;
}

void main() {
  final State state = LoadingState();

  final result = switch(state) {
    LoadingState() => 'Loading',
    SuccessState(name: 'John', age: 30) => 'A 30 year old John',
    SuccessState(name: final name, age: 1) => '$name is just 1 year old',
    SuccessState(name: final name, age: final age) => '$name is $age years old',
  };

  print(result);
}
