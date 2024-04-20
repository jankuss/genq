import 'package:genq/genq.dart';

part 'const_lat_lng.genq.dart';

@genq
class ConstLatLng with _$ConstLatLng {
  const factory ConstLatLng({
    required double lat,
    required double lng,
  }) = _ConstLatLng;
}
