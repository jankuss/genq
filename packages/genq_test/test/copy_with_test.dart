import 'package:genq_test/address.dart';
import 'package:genq_test/user.dart';
import 'package:test/test.dart';

void main() {
  test('copyWith creates a new class with changes applied', () {
    final user = User(
      username: 'Jan',
      email: 'jan@kuss.dev',
      age: null,
    );

    final updatedUser = user.copyWith(username: 'Jan Kuß');

    expect(updatedUser.username, 'Jan Kuß');
    expect(updatedUser.email, 'jan@kuss.dev');
    expect(updatedUser.age, null);
    expect(user, isNot(equals(updatedUser)));
  });

  test('copyWith can unset nullable values', () {
    final user = User(
      username: 'Jan',
      email: 'jan@kuss.dev',
      age: 55,
      address: Address(
        street: 'Kussstraße 42',
        zipCode: 12345,
        city: 'Kussstadt',
      ),
    );

    expect(user.age, isNotNull);
    expect(user.address, isNotNull);
    final updatedUser = user.copyWith(age: null, address: null);
    expect(updatedUser.age, null);
    expect(updatedUser.address, null);
  });

  test('copyWith can copy deeply', () {
    final user = User(
      username: 'Jan',
      email: 'jan@kuss.dev',
      age: 55,
      address: Address(
        street: 'Kussstraße 42',
        zipCode: 12345,
        city: 'Kussstadt',
      ),
    );

    final updatedUser = user.copyWith(
      username: 'Jannik',
      address: user.address?.copyWith(
        street: 'Kussstraße 43',
        zipCode: 12346,
        city: 'Kussdorf',
      ),
    );

    expect(updatedUser.username, 'Jannik');
    expect(updatedUser.address?.street, 'Kussstraße 43');
    expect(updatedUser.address?.zipCode, 12346);
    expect(updatedUser.address?.city, 'Kussdorf');
  });

  test('empty copyWith creates a new instance that is equal', () {
    final user = User(
      username: 'Jan',
      email: 'jan@kuss.dev',
      age: null,
    );

    final updatedUser = user.copyWith();
    expect(identical(user, updatedUser), false);
    expect(user, equals(updatedUser));
  });
}
