package inttest

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestInputParam(t *testing.T) {
	setup()
	defer teardown()

	addGenqTestFile("sut/input_param.dart", "InputParamTest")
	addGenqTestFile("sut/input_param2.dart", "InputParamTest")
	out, err := exec.Command("genq", "--input", "sut/input_param.dart").Output()
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(out), "input_param.genq.dart") {
		t.Fatalf("Expected output to contain 'input_param.genq.dart'")
	}

	if strings.Contains(string(out), "input_param2.genq.dart") {
		t.Fatalf("Did not expected output to contain 'input_param2.genq.dart'")
	}

  if _, err := os.Stat("sut/input_param.genq.dart"); os.IsNotExist(err) {
    t.Fatalf("Expected file 'sut/input_param.genq.dart' to exist")
  }

  if _, err := os.Stat("sut/input_param2.genq.dart"); !os.IsNotExist(err) {
    t.Fatalf("Expected file 'sut/input_param2.genq.dart' to not exist")
  }
}
