import 'package:genq/genq.dart';

part 'json_balance.genq.dart';

@Genq(json: true)
class BalanceSnapshot with _$BalanceSnapshot {
  const factory BalanceSnapshot({
    required BigInt balance,
    required DateTime timestamp,
    required Uri uri,
  }) = _BalanceSnapshot;
}
