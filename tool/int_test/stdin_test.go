package inttest
/* 
import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestStin(t *testing.T) {
	setup()
  defer teardown()

	cmd := exec.Command(os.Getenv("GENQ_PATH"), "--stdin", "sut/stdin.dart")

	buffer := bytes.Buffer{}
	buffer.Write([]byte(`@genq
class TestClazz with _$TestClazz {
  factory TestClazz({
    required String name,
    required int age,
    required List<String> friends,
  });
}`))

	cmd.Stdin = &buffer

	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

  if _, err := os.Stat("sut/stdin.genq.dart"); os.IsNotExist(err) {
    t.Fatalf("Expected file 'sut/stdin.genq.dart' to exist")
  }
} */
