import 'package:genq_test/user.dart';
import 'package:test/test.dart';

void main() {
  test('toString returns a string representation', () {
    final user1 = User(
      username: 'Max',
      email: 'max@mustermann.de',
      age: 52,
    );

    expect(
        user1.toString(),
        equals(
            "User(username: Max, email: max@mustermann.de, age: 52, address: null)"));
  });
}
