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

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if strings.HasSuffix(event.Name, ".genq.dart") {
				continue
			}

			if !strings.HasSuffix(event.Name, ".dart") {
				// only dart files
				continue
			}

			if event.Op&fsnotify.Write != 0 || event.Op&fsnotify.Create != 0 || event.Op&fsnotify.Rename != 0 {
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

  // Before anything is put into the jobs channel, we need to:
  // 1. Spawn a number for workers, which will drain the jobs channel. This needs
  //    to be done before we start putting jobs into the channel, otherwise sending
  //    to the jobs channel will block.
  // 2. Call messageProcessor.start(), which will start a goroutine to process the
  //    messages channel (the results of the jobs). This also needs to be done before
  //    we start putting jobs into the channel, otherwise sending to the messages channel (when a job is done)
  //    will block.
  // After all of that setup, we can start putting jobs into the jobs channel.
  // The job will get picked up by a worker, processed, and the result will be sent to the messages channel.
  // Finally, we call messageProcessor.wait(), which will close the jobs channel once all jobs have been added and received, 
  // and wait for all messages to be processed.
	workerCount := runtime.NumCPU()
	messageProcessor := newMessageProcessor(format)

	for w := 0; w < workerCount; w++ {
		go worker(messageProcessor.jobs, messageProcessor.messages)
	}

	messageProcessor.start()

	// Walk the directory and add jobs to the jobs channel
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		messageProcessor.addJob(job{path: path})
		return nil
	})

	messageProcessor.wait()
}

func ReadString(p string) (string, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

type messageProcessor struct {
	jobWg     *sync.WaitGroup
	messageWg *sync.WaitGroup
	messages  chan resWithPath
	jobs      chan job
	format    bool
}

func newMessageProcessor(format bool) *messageProcessor {
	return &messageProcessor{jobWg: &sync.WaitGroup{}, messageWg: &sync.WaitGroup{}, messages: make(chan resWithPath), format: format, jobs: make(chan job)}
}

func (m *messageProcessor) addJob(job job) {
	m.jobWg.Add(1)
	m.jobs <- job
}

func (m *messageProcessor) wait() {
	close(m.jobs)
	m.jobWg.Wait()

	close(m.messages)
	m.messageWg.Wait()
}

func (m *messageProcessor) start() {
	m.messageWg.Add(1)
	go func() {
		defer m.messageWg.Done()
		genCount := 0
		dartArgs := []string{"format"}
		for i := range m.messages {
			m.jobWg.Done()
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

		if m.format {
			fmt.Printf("ℹ️ Running 'dart format' on %d files\n", len(dartArgs)-1)
			exec.Command("dart", dartArgs...).Run()
		}
	}()
}
