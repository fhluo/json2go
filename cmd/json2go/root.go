package main

import (
	"bytes"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	gen "github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/json2go"
	"github.com/spf13/cobra"
)

func readInput() ([]byte, error) {
	if input == "" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(input)
}

func generate(data []byte) (*gen.File, error) {
	file := gen.NewFile(packageName)

	code, err := json2go.Options{
		TypeName:     typeName,
		AllCapsWords: acronyms,
	}.GenerateCodeFromBytes(data)
	if err != nil {
		return nil, err
	}

	file.Add(code)
	return file, nil
}

func writeOutput(output string, file *gen.File) error {
	buffer := new(bytes.Buffer)
	if err := file.Render(buffer); err != nil {
		return err
	}

	if output == "" {
		_, err := io.Copy(os.Stdout, buffer)
		return err
	}

	// make sure output directory exists
	if err := os.MkdirAll(filepath.Dir(output), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(output, buffer.Bytes(), os.ModePerm)
}

var rootCmd = &cobra.Command{
	Use:   "json2go",
	Short: "Generate Go type definitions from JSON",
	Run: func(cmd *cobra.Command, args []string) {
		// read input
		data, err := readInput()
		if err != nil {
			slog.Error("failed to read input file", "err", err)
			os.Exit(1)
		}

		// generate code
		file, err := generate(data)
		if err != nil {
			slog.Error("failed to generate code", "err", err)
			os.Exit(1)
		}

		// write output
		err = writeOutput(output, file)
		if err != nil {
			slog.Error("failed to write output file", "err", err)
			os.Exit(1)
		}
	},
}

var (
	input       string
	output      string
	packageName string
	typeName    string
	acronyms    []string
)

func init() {
	rootCmd.Flags().StringVarP(&input, "input", "i", "", "input file (default: stdin)")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "output file (default: stdout)")
	rootCmd.Flags().StringVarP(&packageName, "package", "p", "main", "package name")
	rootCmd.Flags().StringVarP(&typeName, "type", "t", "T", "type name")
	rootCmd.Flags().StringSliceVarP(&acronyms, "acronyms", "a", nil, "specify acronyms")

	if err := rootCmd.MarkFlagFilename("input", "json"); err != nil {
		slog.Error("failed to mark input flag filename", "err", err)
		os.Exit(1)
	}

	if err := rootCmd.MarkFlagFilename("output", "go"); err != nil {
		slog.Error("failed to mark input flag filename", "err", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
