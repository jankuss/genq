import 'package:genq/genq.dart';

part 'input.genq.dart';

@genq
class SomeState with _$SomeState {
  const factory SomeState.initial() = Initial;
  const factory SomeState.loading() = Loading;
  const factory SomeState.loaded({ required String data }) = Loaded;
  const factory SomeState.error({ required String message }) = Error;
}

