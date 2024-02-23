import 'package:genq/genq.dart';

part 'input.genq.dart';

@genq
class ShoppingCart with _$ShoppingCart {
  factory ShoppingCart({
    required String ownerName,
    required List<String> items,
  }) = _ShoppingCart;
}
