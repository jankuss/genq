import 'package:genq_test/json/json_etf.dart';
import 'package:test/test.dart';

void main() {
  test('correctly uses custom fromJson', () {
    final value = $ETFFromJson({
      'isin': 'US0378331005',
      'name': 'AAPL',
      'price': 123.45,
    });

    expect(value, ETF(
      isin: ISIN('US0378331005'),
      name: 'AAPL',
      price: 123.45,
    ));
  });

  test('correctly uses custom toJson', () {
    final value = $ETFToJson(
      ETF(
        isin: ISIN('US0378331005'),
        name: 'AAPL',
        price: 123.45,
      ),
    );

    expect(value, {
      'isin': 'US0378331005',
      'name': 'AAPL',
      'price': 123.45,
    });
  });
}
