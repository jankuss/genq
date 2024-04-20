import 'package:genq_test/json/json_balance.dart';
import 'package:test/test.dart';

void main() {
  test('BalanceSnapshot.fromJson correctly parses datetime, bigint & uri', () {
    final value = $BalanceSnapshotFromJson({
      'timestamp': '2021-01-01T00:00:00Z',
      'balance': '12345678901234567890',
      'uri': 'https://example.com',
    });

    expect(value.timestamp, equals(DateTime.utc(2021, 1, 1)));
    expect(value.balance, equals(BigInt.parse('12345678901234567890')));
    expect(value.uri, equals(Uri.parse('https://example.com')));
  });

  test('BalanceSnapshot.toJson correctly serializes datetime and bigint', () {
    final value = $BalanceSnapshotToJson(BalanceSnapshot(
      timestamp: DateTime.utc(2021, 1, 1),
      balance: BigInt.parse('12345678901234567890'),
      uri: Uri.parse('https://example.com'),
    ));

    expect(
      value,
      equals({
        'timestamp': '2021-01-01T00:00:00.000Z',
        'balance': '12345678901234567890',
        'uri': 'https://example.com',
      }),
    );
  });
}
