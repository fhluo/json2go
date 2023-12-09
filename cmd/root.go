package cmd

import (
	"bytes"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	gen "github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/pkg/json2go"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "json2go",
	Short: "Generate Go type definitions from JSON",
	Run: func(cmd *cobra.Command, args []string) {
		// read input
		var (
			data []byte
			err  error
		)

		if input == "" {
			data, err = io.ReadAll(os.Stdin)
		} else {
			data, err = os.ReadFile(input)
		}

		if err != nil {
			slog.Error("failed to read input file", "err", err)
			os.Exit(1)
		}

		// generate code
		file := gen.NewFile(packageName)

		code, err := json2go.Options{
			TypeName:     typeName,
			AllCapsWords: acronyms,
		}.GenerateCodeFromBytes(data)

		if err != nil {
			slog.Error("failed to declare type", "err", err)
			os.Exit(1)
		}
		file.Add(code)

		buffer := new(bytes.Buffer)
		if err = file.Render(buffer); err != nil {
			slog.Error("failed to render file", "err", err)
			os.Exit(1)
		}

		// write output
		if output == "" {
			_, _ = io.Copy(os.Stdout, buffer)
		} else {
			// make sure output directory exists
			_ = os.MkdirAll(filepath.Dir(output), os.ModePerm)

			err = os.WriteFile(output, buffer.Bytes(), os.ModePerm)
		}

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
