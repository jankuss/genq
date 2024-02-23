import 'package:genq/genq.dart';

part "input.genq.dart";

@genq
class Test with _$Test {
  factory Test({
    required int param1,
  }) = _Test;
}
