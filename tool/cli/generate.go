package cli

import (
	"fmt"
	"genq/generation"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

func pre(str string) string {
	return ">  " + str
}

func PrintError(path string, error generation.GenerateResultError) {
	fmt.Println()
	fmt.Printf("Error when parsing: %s:%d\n", path, error.Context.Line+1)
	fmt.Println(pre(error.Context.LineString))

	str := ""
	for i := 0; i < error.Context.PosInLine-1; i++ {
		str = str + " "
	}
	str = str + "^"

	fmt.Println(pre(str))
	fmt.Println(pre(error.Error.Err.Error()))
}

type resWithPath struct {
	path string
	res  generation.GenerateResult
}

func generate(path string, format bool) {
	fmt.Println("ℹ️ Generating code for:", path)
	messages := make(chan resWithPath)

	wg := sync.WaitGroup{}
	totalCount := 0
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		totalCount++
		go func() {
			defer wg.Done()
			res, err := generation.Generate(path, ReadString)

			if err != nil {
				return
			}

      if len(res.WriteString) > 0 {
        os.WriteFile(res.OutputFile, []byte(strings.Join(res.WriteString, "\n")), 0644)
      }

			messages <- resWithPath{
				path: path,
				res:  res,
			}
		}()

		return nil
	})

	go func() {
		wg.Wait()
		close(messages)
	}()

	resCount := 0
  genCount := 0
	args := []string{"format"}
	fmt.Print("\033[s")
	for i := range messages {
    genCount += i.res.GenCount

		resCount++
		fmt.Print("\033[u\033[K")
		fmt.Printf("📝 Generated %s", i.res.OutputFile)

		if len(i.res.Errors) > 0 {
			for _, error := range i.res.Errors {
				PrintError(i.path, error)
			}
		}

		if !i.res.Noop {
			args = append(args, i.res.OutputFile)
		}
	}

	println()

  fmt.Printf("✅ Generated %d data classes\n", genCount)

	if format {
		fmt.Printf("ℹ️ Running 'dart format' on %d files\n", len(args)-1)
		exec.Command("dart", args...).Run()
	}
}

func ReadString(p string) (string, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
