import 'package:genq_test/json/json_collections.dart';
import 'package:test/test.dart';

void main() {
  test('collections fromJson', () {
    final value = $JsonCollectionsFromJson({
      'somethingList': [
        {'name': 'A'},
        {'name': 'B'},
      ],
      'somethingSet': [
        {'name': 'A'},
        {'name': 'B'},
      ],
      'somethingSetNullableValue': [
        {'name': 'A'},
        null,
        {'name': 'B'},
      ],
      'somethingMap': {
        'a': {'name': 'A'},
        'b': {'name': 'B'},
      },
      'somethingUriMap': {
        'https://example.com': {'name': 'A'},
        'https://example.org': {'name': 'B'},
      },
    });

    expect(value.somethingList,
        equals([Something(name: 'A'), Something(name: 'B')]));
    expect(value.somethingSet,
        equals({Something(name: 'A'), Something(name: 'B')}));
    expect(value.somethingSetNullableValue,
        equals({Something(name: 'A'), null, Something(name: 'B')}));
    expect(value.somethingMap,
        equals({'a': Something(name: 'A'), 'b': Something(name: 'B')}));
    expect(
      value.somethingUriMap,
      equals({
        Uri.parse('https://example.com'): Something(name: 'A'),
        Uri.parse('https://example.org'): Something(name: 'B'),
      }),
    );
  });

  test('collections toJson', () {
    final value = $JsonCollectionsToJson(JsonCollections(
      somethingList: [Something(name: 'A'), Something(name: 'B')],
      somethingSet: {Something(name: 'A'), Something(name: 'B')},
      somethingSetNullableValue: {
        Something(name: 'A'),
        null,
        Something(name: 'B')
      },
      somethingMap: {'a': Something(name: 'A'), 'b': Something(name: 'B')},
      somethingUriMap: {
        Uri.parse('https://example.com'): Something(name: 'A'),
        Uri.parse('https://example.org'): Something(name: 'B'),
      },
    ));

    expect(
      value,
      equals({
        'somethingList': [
          {'name': 'A'},
          {'name': 'B'},
        ],
        'somethingSet': [
          {'name': 'A'},
          {'name': 'B'},
        ],
        'somethingSetNullableValue': [
          {'name': 'A'},
          null,
          {'name': 'B'},
        ],
        'somethingMap': {
          'a': {'name': 'A'},
          'b': {'name': 'B'},
        },
        'somethingUriMap': {
          'https://example.com': {'name': 'A'},
          'https://example.org': {'name': 'B'},
        },
      }),
    );

    expect(value['somethingList'], isA<List>());
    expect(value['somethingSet'], isA<List>());
    expect(value['somethingSetNullableValue'], isA<List>());
    expect(value['somethingMap'], isA<Map>());
    expect(value['somethingUriMap'], isA<Map>());
  });
}
