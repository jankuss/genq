import 'package:genq_test/test_classes.dart';
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

  test('== returns true for same collections', () {
    final user = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);
    final collectionTests1 = CollectionTests(
      users: [user],
      userMap: {"1": user},
      userSet: {
        user,
      },
    );
    final collectionTests2 = CollectionTests(
      users: [user],
      userMap: {"1": user},
      userSet: {
        user,
      },
    );

    expect(collectionTests1, equals(collectionTests2));
  });

  test('== returns false for modified list', () {
    final user = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);
    final collectionTests1 = CollectionTests(
      users: [user],
      userMap: {"1": user},
      userSet: {
        user,
      },
    );
    final collectionTests2 = CollectionTests(
      users: [user, user],
      userMap: {"1": user},
      userSet: {
        user,
      },
    );

    expect(collectionTests1, isNot(equals(collectionTests2)));
  });

  test('== returns false for modified map', () {
    final user = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);
    final collectionTests1 = CollectionTests(
      users: [user],
      userMap: {"1": user},
      userSet: {
        user,
      },
    );
    final collectionTests2 = CollectionTests(
      users: [user],
      userMap: {"1": user, "2": user},
      userSet: {
        user,
      },
    );

    expect(collectionTests1, isNot(equals(collectionTests2)));
  });

  test('== returns false for modified set', () {
    final user = User(username: 'Jan', email: 'jan@kuss.dev', age: 42);
    final user2 = User(username: 'Jannik', email: 'jan@kuss.dev', age: 42);
    final collectionTests1 = CollectionTests(
      users: [user],
      userMap: {"1": user},
      userSet: {
        user,
      },
    );
    final collectionTests2 = CollectionTests(
      users: [user],
      userMap: {"1": user},
      userSet: {
        user,
        user2,
      },
    );

    expect(collectionTests1, isNot(equals(collectionTests2)));
  });
}
