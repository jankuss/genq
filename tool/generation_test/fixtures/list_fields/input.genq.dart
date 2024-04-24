part of 'input.dart';

mixin _$ShoppingCart {
  String get ownerName => throw UnimplementedError();
  List<String> get items => throw UnimplementedError();

  $ShoppingCartCopyWith get copyWith => throw UnimplementedError();
}

class _ShoppingCart implements ShoppingCart {
  @override
  final String ownerName;

  @override
  final List<String> items;

  _ShoppingCart({
    required this.ownerName,
    required this.items,
  });

  @override
  $ShoppingCartCopyWith get copyWith => _$ShoppingCartCopyWithImpl(this);

  @override
  String toString() {
    return "ShoppingCart(ownerName: $ownerName, items: $items)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! ShoppingCart) return false;
    if (!identical(other.ownerName, ownerName) && other.ownerName != ownerName) return false;
    if (!const DeepCollectionEquality().equals(other.items, items)) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      ownerName,
      items,
    );
  }
}

abstract class $ShoppingCartCopyWith {
  ShoppingCart call({
    String ownerName,
    List<String> items,
  });
}

class _$ShoppingCartCopyWithImpl implements $ShoppingCartCopyWith {
  final _$ShoppingCart value;

  _$ShoppingCartCopyWithImpl(this.value);

  @override
  ShoppingCart call({
    Object? ownerName = genq,
    Object? items = genq,
  }) {
    return ShoppingCart(
      ownerName: ownerName == genq ? value.ownerName : ownerName as String,
      items: items == genq ? value.items : items as List<String>,
    );
  }
}
