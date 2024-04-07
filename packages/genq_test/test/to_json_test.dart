import 'package:genq_test/json/json_address.dart';
import 'package:genq_test/json/json_user.dart';
import 'package:genq_test/json/json_user_status.dart';
import 'package:test/test.dart';

void main() {
  test('returns json representation of JsonUser', () {
    final value = $JsonUserToJson(JsonUser(
      name: 'Test',
      age: 42,
      registered: true,
      status: JsonUserStatus.active,
      address: JsonAddress(
        street: 'Street 1',
        zip: '12345',
        city: 'City',
        state: 'State',
      ),
    ));

    expect(
      value,
      equals({
        'full_name': 'Test',
        'age': 42,
        'registered': true,
        'status': 'active',
        'address': {
          'street': 'Street 1',
          'city': 'City',
          'state': 'State',
          'zip': '12345',
        },
      }),
    );
  });

  test('returns json representation with null value', () {
    final value = $JsonUserToJson(JsonUser(
      name: 'Test',
      age: 42,
      registered: true,
      status: JsonUserStatus.active,
      address: null,
    ));

    expect(
      value,
      equals({
        'full_name': 'Test',
        'age': 42,
        'registered': true,
        'status': 'active',
        'address': null,
      }),
    );
  });
}
