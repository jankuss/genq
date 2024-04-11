package inttest

import (
	"os/exec"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
  setup()
  defer teardown()

  out, err := exec.Command("genq", "--version").Output()
  if err != nil {
    t.Fatal(err)
  }

  if !strings.Contains(string(out), "version: ") {
    t.Fatalf("Expected output to contain 'version: '")
  }

  if !strings.Contains(string(out), "commit: ") {
    t.Fatalf("Expected output to contain 'commit: '")
  }

  if !strings.Contains(string(out), "date: ") {
    t.Fatalf("Expected output to contain 'date: '")
  }
}
