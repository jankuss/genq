import 'package:genq_test/json/json_numeric.dart';
import 'package:test/test.dart';

void main() {
  test("fromJson handles numeric values", () {
    final json = {
      'intField': 42,
      'doubleField': 42.42,
      'intFieldNullable': 42,
      'doubleFieldNullable': 42.42,
      'numField': 42,
      'numFieldNullable': 42,
    };

    final value = $JsonNumericFromJson(json);

    expect(value.intField, equals(42));
    expect(value.doubleField, equals(42.42));
    expect(value.intFieldNullable, equals(42));
    expect(value.doubleFieldNullable, equals(42.42));
    expect(value.numField, equals(42));
    expect(value.numFieldNullable, equals(42));
  });

  test("fromJson handles numeric values of different types", () {
    final json = {
      'intField': 42.0,
      'doubleField': 42,
      'intFieldNullable': 42.0,
      'doubleFieldNullable': 42,
      'numField': 42.0,
      'numFieldNullable': 42.0,
    };

    final value = $JsonNumericFromJson(json);

    expect(value.intField, equals(42));
    expect(value.doubleField, equals(42.0));
    expect(value.intFieldNullable, equals(42));
    expect(value.doubleFieldNullable, equals(42.0));
    expect(value.numField, equals(42));
    expect(value.numFieldNullable, equals(42));
  });

  test("fromJson handles nullable values", () {
    final json = {
      'intField': 42,
      'doubleField': 42.0,
      'intFieldNullable': null,
      'doubleFieldNullable': null,
      'numField': 42,
      'numFieldNullable': null,
    };

    final value = $JsonNumericFromJson(json);

    expect(value.intField, equals(42));
    expect(value.doubleField, equals(42.0));
    expect(value.intFieldNullable, isNull);
    expect(value.doubleFieldNullable, isNull);
    expect(value.numField, equals(42));
    expect(value.numFieldNullable, isNull);
  });

  test("fromJson and then toJson has the actual dart values", () {
    final json = {
      'intField': 42.0,
      'doubleField': 42,
      'intFieldNullable': 42.0,
      'doubleFieldNullable': 42,
      'numField': 42,
      'numFieldNullable': 42.0,
    };

    final value = $JsonNumericFromJson(json);
    final actualJson = $JsonNumericToJson(value);

    expect(actualJson['intField'], equals(42));
    expect(actualJson['doubleField'], equals(42.0));
    expect(actualJson['intFieldNullable'], equals(42));
    expect(actualJson['doubleFieldNullable'], equals(42.0));
    expect(actualJson['numField'], equals(42));
    expect(actualJson['numFieldNullable'], equals(42.0));
  });
}
