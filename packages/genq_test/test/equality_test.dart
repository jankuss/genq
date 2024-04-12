import 'package:genq_test/address.dart';
import 'package:genq_test/user.dart';
import 'package:test/test.dart';

void main() {
  test('== is true for equal objects', () {
    final user1 = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);
    final user2 = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);

    expect(user1, equals(user2));
  });

  test('== is false for not equal objects (differ in age)', () {
    final user1 = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);
    final user2 = User(username: 'Jan', email: 'jan@kuss.dev', age: 52);

    expect(user1, isNot(equals(user2)));
  });

  test('== is false for not equal objects (differ in email)', () {
    final user1 = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);
    final user2 = User(username: 'Jan', email: 'jan@kuss.de', age: 52);

    expect(user1, isNot(equals(user2)));
  });

  test('== is false for not equal objects (differ in name and email)', () {
    final user1 = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);
    final user2 = User(username: 'Jannik', email: 'jannik@kuss.de', age: 52);

    expect(user1, isNot(equals(user2)));
  });

  test('== is true for identical instance', () {
    final user1 = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);

    expect(user1, equals(user1));
  });

  test('== is false for different type', () {
    final user = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);
    final address =
        Address(street: 'Kussstraße 42', zipCode: 12345, city: 'Kussstadt');

    expect(user, isNot(equals(address)));
  });

  test('== is true for deeply copied', () {
    final address =
        Address(street: 'Kussstraße 42', zipCode: 12345, city: 'Kussstadt');
    final user =
        User(username: 'Jan', email: 'jan@kuss.dev', age: 42, address: address);
    final user2 = user.copyWith(address: address.copyWith());

    expect(user, equals(user2));
  });
}
