package cli

import (
	"fmt"
	"os"

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

		_, err = cmd.Flags().GetBool("version")
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
	},
}

func Execute() {
	rootCmd.Flags().StringP("input", "i", ".", "The input file or directory to generate code from")
	rootCmd.Flags().BoolP("format", "f", false, "Format the generated code with dart format")
	rootCmd.Flags().Bool("version", false, "Prints the version of genq")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
