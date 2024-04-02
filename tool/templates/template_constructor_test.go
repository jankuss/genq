package templates

import "testing"

func TestDefaultTemplateConstructor(t *testing.T) {
	actual := []string{}

	parsed := getFirstGenqClass(`
  @genq
  class TestClazz {
    factory TestClazz({
      required String name,
      required int age,
      required List<String> friends,
    });
  }`)

	actual = templateConstructor(actual, parsed)
	expect := getNormalized(`
class _TestClazz implements TestClazz {
  @override
  final String name;

  @override
  final int age;

  @override
  final List<String> friends;

  _TestClazz({
    required this.name,
    required this.age,
    required this.friends,
  });

  @override
  $TestClazzCopyWith get copyWith => _$TestClazzCopyWithImpl(this);

  @override
  String toString() {
    return "TestClazz(name: $name, age: $age, friends: $friends)";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! TestClazz) return false;
    if (!identical(other.name, name) && other.name != name) return false;
    if (!identical(other.age, age) && other.age != age) return false;
    if (!const DeepCollectionEquality().equals(other.friends, friends)) return false;
    return true;
  }

  @override
  int get hashCode {
    return Object.hash(
      runtimeType,
      name,
      age,
      friends,
    );
  }
}`)

	compare(t, actual, expect)
}

func TestEmpty(t *testing.T) {
	actual := []string{}

	parsed := getFirstGenqClass(`
  @genq
  class TestClazz {
    factory TestClazz();
  }`)

	actual = templateConstructor(actual, parsed)
	expect := getNormalized(`
class _TestClazz implements TestClazz {
  _TestClazz();

  @override
  $TestClazzCopyWith get copyWith => _$TestClazzCopyWithImpl(this);

  @override
  String toString() {
    return "TestClazz()";
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    if (other is! TestClazz) return false;
    return true;
  }

  @override
  int get hashCode {
    return runtimeType.hashCode;
  }
}`)

	compare(t, actual, expect)
}
