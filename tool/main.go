package main

import (
	"fmt"
	"genq/cli"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// main with arguments
func main() {
  for _, arg := range os.Args {
    if arg == "--version" {
      fmt.Println("=========================================================")
      fmt.Println("genq - Blazingly fast code generator for Dart and Flutter")
      fmt.Println("=========================================================")

      fmt.Println("version:", version)
      fmt.Println("commit:", commit)
      fmt.Println("date:", date)
      os.Exit(0)
    }
  }

	cli.Execute()
}
