part of 'json_etf.dart';

mixin _$ETF {
  ISIN get isin => throw UnimplementedError();
  String get name => throw UnimplementedError();
  double get price => throw UnimplementedError();

  $ETFCopyWith get copyWith => throw UnimplementedError();
}

class _ETF implements ETF {
  @override
  final ISIN isin;

  @override
  final String name;

  @override
  final double price;

  _ETF({
    required this.isin,
    required this.name,
    required this.price,
  });

  @override
  $ETFCopyWith get copyWith => _$ETFCopyWithImpl(this);

  @override
  String toString() {
    return "ETF(isin: $isin, name: $name, price: $price)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! ETF) return false;
    if (!identical(other.isin, isin) && other.isin != isin) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.price, price) && other.price != price) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      isin,
      name,
      price,
    );
  }
}

abstract class $ETFCopyWith {
  ETF call({
    ISIN isin,
    String name,
    double price,
  });
}

class _$ETFCopyWithImpl implements $ETFCopyWith {
  final _$ETF __value;

  _$ETFCopyWithImpl(this.__value);

  @override
  ETF call({
    Object? isin = genq,
    Object? name = genq,
    Object? price = genq,
  }) {
    return ETF(
      isin: isin == genq ? __value.isin : isin as ISIN,
      name: name == genq ? __value.name : name as String,
      price: price == genq ? __value.price : price as double,
    );
  }
}

ETF $ETFFromJson(Map<String, dynamic> json) {
  return ETF(
    isin: ISIN.fromJson(json['isin']),
    name: json['name'] as String,
    price: (json['price'] as num).toDouble(),
  );
}

Map<String, dynamic> $ETFToJson(ETF obj) {
  return {
    'isin': ISIN.toJson(obj.isin),
    'name': obj.name,
    'price': obj.price,
  };
}
