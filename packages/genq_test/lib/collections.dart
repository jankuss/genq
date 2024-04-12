import 'package:genq/genq.dart';
import 'package:genq_test/user.dart';

part 'collections.genq.dart';

@genq
class CollectionTests with _$CollectionTests {
  factory CollectionTests({
    required List<User> users,
    required Map<String, User> userMap,
    required Set<User> userSet,
  }) = _CollectionTests;
}

