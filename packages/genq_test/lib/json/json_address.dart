import 'package:genq/genq.dart';

part 'json_address.genq.dart';

@Genq(json: true)
class JsonAddress with _$JsonAddress {
  factory JsonAddress({
    required String street,
    required String city,
    required String state,
    required String zip,
  }) = _JsonAddress;
}
