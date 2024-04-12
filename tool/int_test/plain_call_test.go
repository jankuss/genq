package inttest

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestPlainCall(t *testing.T) {
	setup()
	defer teardown()

	addGenqTestFile("sut/test.dart", "PlainCallTest")
	addGenqTestFile("sut/abc.dart", "PlainCallTest")

	out, error := exec.Command(os.Getenv("GENQ_PATH")).Output()
	if error != nil {
		t.Fatalf("Did not expect error: %s", error)
	}

	if !strings.Contains(string(out), "test.genq.dart") {
		t.Fatalf("Expected output to contain 'test.genq.dart'")
	}

	if !strings.Contains(string(out), "abc.genq.dart") {
		t.Fatalf("Expected output to contain 'test.genq.dart'")
	}

	if _, err := os.Stat("sut/test.genq.dart"); os.IsNotExist(err) {
		t.Fatalf("Expected file 'sut/test.genq.dart' to exist")
	}

	if _, err := os.Stat("sut/abc.genq.dart"); os.IsNotExist(err) {
		t.Fatalf("Expected file 'sut/abc.genq.dart' to exist")
	}

	content, err := os.ReadFile("sut/test.genq.dart")
	if err != nil {
		t.Fatalf("Did not expect error: %s", err)
	}

	if len(content) == 0 {
		t.Fatalf("Expected file 'sut/test.genq.dart' to contain content")
	}

	content, err = os.ReadFile("sut/abc.genq.dart")
	if err != nil {
		t.Fatalf("Did not expect error: %s", err)
	}

	if len(content) == 0 {
		t.Fatalf("Expected file 'sut/abc.genq.dart' to contain content")
	}
}
