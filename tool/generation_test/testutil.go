package generationtest

import (
	"fmt"
	. "genq/generation"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func readInput(path string) string {
	b, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Trim(string(b), "\n")
}

func readOutput(path string) []string {
	b, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(strings.Trim(string(b), "\n"), "\n")
}

func testGenOutput(t *testing.T, folder string) {
	var read ReadContent = func(path string) (string, error) {
		return readInput(fmt.Sprintf("./fixtures/%s/input.dart", folder)), nil
	}

	res, err := Generate("./input.dart", read,)
	if err != nil {
		t.Fatalf("Did not expect error: %s", err)
	}

	if len(res.Errors) > 0 {
		for _, err := range res.Errors {
			t.Fatalf("Error at line %d, pos %d: %s", err.Context.Line, err.Context.PosInLine, err.Error.Err.Error())
		}
	}

	if diff := cmp.Diff(readOutput(fmt.Sprintf("./fixtures/%s/input.genq.dart", folder)), res.WriteString); diff != "" {
		println("== GENERATED OUTPUT WAS ==")
		println(strings.Join(res.WriteString, "\n"))
		println("==========================")

		if os.Getenv("UPDATE_FIXTURES") == "1" {
			os.WriteFile(fmt.Sprintf("./fixtures/%s/input.genq.dart", folder), []byte(strings.Join(res.WriteString, "\n")), 0644)
		}

		t.Fatalf(diff)
	}

	cmd := exec.Command("dart", "analyze", fmt.Sprintf("./fixtures/%s/input.genq.dart", folder))
	stdout, err := cmd.Output()
	if err != nil {
		println(string(stdout))
		t.Fatal(err)
	}
}
