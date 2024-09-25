package inttest

import (
	"fmt"
	"os"
	"testing"
)

func addFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)

	if err != nil {
		panic(err)
	}
}

func addGenqTestFile(path string, name string) {
	addFile(path, fmt.Sprintf(`@genq
class %s with _$%s {
  factory %s({
    required String name,
    required int age,
    required List<String> friends,
  });
}`, name, name, name))
}

func setup() {
	err := os.Mkdir("sut", 0755)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	err := os.RemoveAll("sut")
	if err != nil {
		panic(err)
	}
}

func assertFileWithContent(t *testing.T, path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Did not expect error: %s", err)
	}

	if len(content) == 0 {
		t.Fatalf("Expected file '%s' to contain content", path)
	}
}
