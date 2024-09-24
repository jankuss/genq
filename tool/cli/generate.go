package cli

import (
	"fmt"
	"genq/generation"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/farmergreg/rfsnotify"
	"gopkg.in/fsnotify.v1"
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
	err  error
}

type job struct {
	path string
}

func processJob(job job) resWithPath {
	res, err := generation.Generate(job.path, ReadString)
	if err != nil {
		return resWithPath{
			err: err,
		}
	}

	if len(res.WriteString) > 0 {
		os.WriteFile(res.OutputFile, []byte(strings.Join(res.WriteString, "\n")), 0644)
	}

	return resWithPath{
		path: job.path,
		res:  res,
	}
}

func worker(jobs <-chan job, results chan<- resWithPath) {
	for j := range jobs {
		res := processJob(j)
		results <- res
	}
}

func generateWatchMode(path string, format bool) {
	// Do one initial generation
	generate(path, format)

	// Watch for changes
	fmt.Println("👀 Watching for changes in:", path)

	watcher, err := rfsnotify.NewWatcher()
	if err != nil {
		fmt.Println("❌ Error watching for changes:", err)
		return
	}

	watcher.AddRecursive(path)
	lastGeneratedMap := make(map[string]time.Time)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if strings.HasSuffix(event.Name, ".genq.dart") {
				continue
			}

			if strings.HasSuffix(event.Name, ".dart~") {
				// Ignore temporary files
				continue
			}

			if event.Op&fsnotify.Write != 0 || event.Op&fsnotify.Create != 0 || event.Op&fsnotify.Rename != 0 {
				lastGeneratedMap[event.Name] = time.Now()

				i := processJob(job{path: event.Name})
				if i.err != nil {
					continue
				}

				if !i.res.Noop {
					fmt.Printf("📝 Generated %s\n", i.res.OutputFile)
				}

				if len(i.res.Errors) > 0 {
					for _, error := range i.res.Errors {
						PrintError(i.path, error)
					}
				}

				if !i.res.Noop && format {
					fmt.Printf("ℹ️ Running 'dart format' on 1 file\n")
					args := []string{"format", i.res.OutputFile}
					exec.Command("dart", args...).Run()
				}
			}
		case err := <-watcher.Errors:
			fmt.Println("❌ Error watching for changes:", err)
			return
		}
	}
}

func generate(path string, format bool) {
	fmt.Println("ℹ️ Generating code for:", path)

	// Worker count is CPU count
	workerCount := runtime.NumCPU()
	jobs := make(chan job)
	messages := make(chan resWithPath)

	// Spawn {CPU count} workers to process jobs
	for w := 0; w < workerCount; w++ {
		go worker(jobs, messages)
	}

	wg := sync.WaitGroup{}
	msgWg := sync.WaitGroup{}

	msgWg.Add(1)
	go func() {
		defer msgWg.Done()
		genCount := 0
		dartArgs := []string{"format"}
		for i := range messages {
			wg.Done()
			genCount += i.res.GenCount

			if i.err != nil {
				continue
			}

			if !i.res.Noop {
				fmt.Printf("📝 Generated %s\n", i.res.OutputFile)
			}

			if len(i.res.Errors) > 0 {
				for _, error := range i.res.Errors {
					PrintError(i.path, error)
				}
			}

			if !i.res.Noop {
				dartArgs = append(dartArgs, i.res.OutputFile)
			}
		}

		fmt.Printf("✅ Generated %d data classes\n", genCount)

		if format {
			fmt.Printf("ℹ️ Running 'dart format' on %d files\n", len(dartArgs)-1)
			exec.Command("dart", dartArgs...).Run()
		}
	}()

	// Collect jobs in slice
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		jobs <- job{path: path}
		return nil
	})
	close(jobs)

	go func() {
		// After all jobs produced a result, close the messages channel
		wg.Wait()
		close(messages)
	}()

	msgWg.Wait()
}

func ReadString(p string) (string, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
