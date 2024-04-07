import 'package:genq_test/json/json_address.dart';
import 'package:genq_test/json/json_user.dart';
import 'package:genq_test/json/json_user_status.dart';
import 'package:test/test.dart';

void main() {
  test('fromJson parses json', () {
    final value = $JsonUserFromJson({
      'full_name': 'Test',
      'age': 42,
      'registered': true,
      'status': 'active',
      'address': {
        'street': 'Main Street',
        'city': 'Springfield',
        'state': 'Illinois',
        'zip': '62701',
      },
    });

    expect(
      value,
      equals(
        JsonUser(
          name: 'Test',
          age: 42,
          registered: true,
          status: JsonUserStatus.active,
          address: JsonAddress(
            street: 'Main Street',
            city: 'Springfield',
            state: 'Illinois',
            zip: '62701',
          ),
        ),
      ),
    );
  });

  test('fromJson parses when nullable field is absent', () {
    final value = $JsonUserFromJson({
      'full_name': 'Super',
      'age': 10,
      'registered': false,
      'status': 'INACTIVE',
    });

    expect(
      value,
      equals(
        JsonUser(
          name: 'Super',
          age: 10,
          registered: false,
          status: JsonUserStatus.inactive,
        ),
      ),
    );
  });

  test('fromJson parses when nullable field is null', () {
    final value = $JsonUserFromJson({
      'full_name': 'Super',
      'age': 10,
      'registered': false,
      'status': 'INACTIVE',
      'address': null,
    });

    expect(
      value,
      equals(
        JsonUser(
          name: 'Super',
          age: 10,
          registered: false,
          status: JsonUserStatus.inactive,
        ),
      ),
    );
  });

  test('throws when JSON cannot be parsed, due to invalid type', () {
    expect(() {
      $JsonUserFromJson({
        'full_name': 'Super',
        'age': 10,
        'registered': 123,
        'status': 'INACTIVE',
      });
    }, throwsA(anything));
  });

  test('throws when JSON cannot be parsed due to invalid enum', () {
    expect(() {
      $JsonUserFromJson({
        'full_name': 'Super',
        'age': 10,
        'registered': false,
        'status': 'iNaCtIvE',
      });
    }, throwsA(anything));
  });
}
