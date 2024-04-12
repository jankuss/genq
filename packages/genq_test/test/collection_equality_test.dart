import 'package:genq_test/collections.dart';
import 'package:genq_test/user.dart';
import 'package:test/test.dart';

void main() {
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
