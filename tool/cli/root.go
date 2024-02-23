package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "genq",
	Short: "genq is a blazingly fast code generator for Dart and Flutter",
	Run: func(cmd *cobra.Command, args []string) {
		inputPath, err := cmd.Flags().GetString("input")
		if err != nil {
			panic(err)
		}

		format, err := cmd.Flags().GetBool("format")
		if err != nil {
			panic(err)
		}

		fmt.Println("   __ _  ___ _ __   __ _ ")
		fmt.Println("  / _` |/ _ \\ '_ \\ / _` |")
		fmt.Println(" | (_| |  __/ | | | (_| |")
		fmt.Println("  \\__, |\\___|_| |_|\\__, |")
		fmt.Println("   __/ |              | |")
		fmt.Println("  |___/               |_|")
		fmt.Println()
		fmt.Println("===================================================")
		fmt.Println("Blazingly fast code generator for Dart and Flutter")
		fmt.Println("===================================================")
		fmt.Println()

		generate(inputPath, format)

		watch, err := cmd.Flags().GetBool("watch")
		if err != nil {
			panic(err)
		}

		if watch {
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				panic(err)
			}
			defer watcher.Close()
			go func() {
				for {
					select {
					case event, ok := <-watcher.Events:
						if !ok {
							return
						}
						if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) || event.Has(fsnotify.Remove) || event.Has(fsnotify.Rename) {
							generate(event.Name, format)
						}
					case err, ok := <-watcher.Errors:
						if !ok {
							return
						}
						log.Println("error:", err)
					}
				}
			}()

			err = watcher.Add(inputPath)
			if err != nil {
				panic(err)
			}
			fmt.Println("Watching for changes...")
			<-make(chan struct{})
		}
	},
}

func Execute() {
	rootCmd.Flags().StringP("input", "i", ".", "The input file or directory to generate code from")
	rootCmd.Flags().BoolP("watch", "w", false, "Watch for changes and regenerate code")
	rootCmd.Flags().BoolP("format", "f", false, "Format the generated code with dart format")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
