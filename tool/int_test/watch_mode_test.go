package inttest

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestWatchMode(t *testing.T) {
	setup()
	defer teardown()

	addGenqTestFile("sut/changing_file.dart", "Class1")
	addGenqTestFile("sut/changing_file2.dart", "Class2")

	cmd := exec.Command(os.Getenv("GENQ_PATH"), "--watch", "--input", "sut")
	r, w, _ := os.Pipe()
	cmd.Stdout = w

	go func() {
		cmd.Run()
	}()

	// TODO: Is there a better way instead of sleeping an arbitrary amount?
	time.Sleep(2 * time.Second)
	addGenqTestFile("sut/changing_file.dart", "Class3")
	addGenqTestFile("sut/changing_file2.dart", "Class4")
	addGenqTestFile("sut/changing_file3.dart", "Class5")
	err := os.MkdirAll("sut/subdir", 0755)
	if err != nil {
		t.Fatalf("Did not expect error: %s", err)
	}
	time.Sleep(2 * time.Second)

	addGenqTestFile("sut/subdir/changing_file4.dart", "Class6")
	time.Sleep(2 * time.Second)

	cmd.Process.Kill()
	w.Close()
	out, _ := io.ReadAll(r)

	content := string(out)
	if !strings.Contains(content, "Watching for changes in: sut") {
		t.Fatalf("Expected output to contain info about watching, was: %s", content)
	}

	if strings.Count(content, "Generated sut/changing_file.genq.dart") > 2 {
		t.Fatalf("Expected output to contain 2x 'sut/changing_file.genq.dart', was: %s", content)
	}

	if strings.Count(content, "Generated sut/changing_file2.genq.dart") > 2 {
		t.Fatalf("Expected output to contain 2x 'sut/changing_file2.genq.dart', was %s", content)
	}

	if strings.Count(content, "Generated sut/changing_file3.genq.dart") > 1 {
		t.Fatalf("Expected output to contain 1x 'sut/changing_file3.genq.dart' was %s", content)
	}

	if strings.Count(content, "Generated sut/subdir/changing_file4.genq.dart") > 1 {
		t.Fatalf("Expected output to contain 1x 'sut/subdir/changing_file4.genq.dart' was %s", content)
	}

	assertFileWithContent(t, "sut/changing_file.genq.dart")
	assertFileWithContent(t, "sut/changing_file2.genq.dart")
	assertFileWithContent(t, "sut/changing_file3.genq.dart")
	assertFileWithContent(t, "sut/subdir/changing_file4.genq.dart")
}
