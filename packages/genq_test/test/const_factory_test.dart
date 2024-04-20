import 'package:genq_test/const_lat_lng.dart';
import 'package:test/test.dart';

void main() {
  test('two normal instances are identical', () {
    expect(identical(ConstLatLng(lat: 1, lng: 2), ConstLatLng(lat: 1, lng: 2)), isFalse);
  });

  test('two const instances are identical', () {
    expect(identical(const ConstLatLng(lat: 1, lng: 2), const ConstLatLng(lat: 1, lng: 2)), isTrue);
  });
}
