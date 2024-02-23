package templates

import (
	"testing"
)

func TestDefaultTemplateMixin(t *testing.T) {
	actual := []string{}

  parsed := getFirstGenqClass(`
  @genq
  class TestClazz {
    factory TestClazz({
      required String name,
      required int age,
      required List<String> friends,
    });
  }`);
	actual = templateMixin(actual, parsed)

	expect := getNormalized(`
mixin _$TestClazz {
  String get name => throw UnimplementedError();
  int get age => throw UnimplementedError();
  List<String> get friends => throw UnimplementedError();

  $TestClazzCopyWith get copyWith => throw UnimplementedError();
}`)

  compare(t, actual, expect)
}

