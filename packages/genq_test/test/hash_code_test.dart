import 'package:genq_test/user.dart';
import 'package:test/test.dart';

void main() {
  test('hashCode is equal for equal objects', () {
    final user1 = User(
      username: 'Jan',
      email: 'max@mustermann.de',
      age: 52,
    );
    final user2 = User(
      username: 'Jan',
      email: 'max@mustermann.de',
      age: 52,
    );

    expect(user1.hashCode, equals(user2.hashCode));
  });

  test('hashCode is not equal for differing objects', () {
    final user1 = User(
      username: 'Max',
      email: 'max@mustermann.de',
      age: 52,
    );
    final user2 = User(
      username: 'Max',
      email: 'max@mustermann.de',
      age: 42,
    );

    expect(user1.hashCode, isNot(equals(user2.hashCode)));
  });
}
