import 'package:genq_test/json/json_default_values.dart';
import 'package:test/test.dart';

void main() {
  test('default values are applied when null', () {
    final value = $TestDefaultValuesFromJson({
      'registered': true,
    });

    expect(
      value,
      equals(
        TestDefaultValues(
          name: 'Supername',
          age: 999,
          registered: true,
        ),
      ),
    );
  });

  test('default values are not applied when not null', () {
    final value = $TestDefaultValuesFromJson({
      'name': 'Test',
      'age': 42,
      'registered': true,
    });

    expect(
      value,
      equals(
        TestDefaultValues(
          name: 'Test',
          age: 42,
          registered: true,
        ),
      ),
    );
  });
}
