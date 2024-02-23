import 'package:genq/genq.dart';

part 'input.genq.dart';

@genq
class User<T> with _$User<T> {
  factory User({
    required T data,
  }) = _User<T>;
}
