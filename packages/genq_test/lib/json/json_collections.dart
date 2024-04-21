import 'package:genq/genq.dart';
import 'package:genq_test/json/json_account.dart';

part 'json_collections.genq.dart';

@Genq(json: true)
class JsonCollections with _$JsonCollections {
  factory JsonCollections({
    required List<Something> somethingList,
    required Set<Something> somethingSet,
    required Set<Something?> somethingSetNullableValue,
    required Map<String, Something> somethingMap,
    required Map<Uri, Something> somethingUriMap,
  }) = _JsonCollections;
}

@Genq(json: true)
class Something with _$Something {
  factory Something({
    required String name,
  }) = _Something;
}
