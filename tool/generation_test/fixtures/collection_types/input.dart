import 'package:genq/genq.dart';

part 'input.genq.dart';

@genq
class User with _$User {
  factory User({
    required List<String> names,
    required Map<String, dynamic> properties,
    required Set<int> ages,
  }) = _User;
}
