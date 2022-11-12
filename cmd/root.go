package cmd

import (
	"fmt"
	"github.com/fhluo/json2go/internal/def"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var rootCmd = &cobra.Command{
	Use:   "json2go",
	Short: "Generating a Go type definition from JSON.",
	Run: func(cmd *cobra.Command, args []string) {
		def.SetAcronyms(acronyms...)

		paths, err := expand(args)
		if err != nil {
			log.Fatalln(err)
		}

		var wg sync.WaitGroup
		wg.Add(len(paths))

		_ = os.MkdirAll(output, os.ModePerm)

		for _, path := range paths {
			path := path
			go func() {
				data, err := os.ReadFile(path)
				if err != nil {
					log.Println(err)
					return
				}

				base := filepath.Base(path)
				name := strings.TrimSuffix(base, filepath.Ext(path))

				data, err = Generate(data, output, fmt.Sprintf("Generated from %s", base), def.ToCamelCase(name))
				if err != nil {
					log.Println(err)
					return
				}

				dst := filepath.Join(output, def.ToSnakeCase(name)+".go")
				err = os.WriteFile(dst, data, os.ModePerm)
				if err != nil {
					log.Println(err)
					return
				}

				fmt.Printf("generated: %s => %s\n", path, dst)
				wg.Done()
			}()
		}

		wg.Wait()
	},
}

var (
	output   string
	acronyms []string
)

func init() {
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "output folder")
	rootCmd.Flags().StringSliceVarP(&acronyms, "acronyms", "a", nil, "specify acronyms")

	err := rootCmd.MarkFlagDirname("output")
	if err != nil {
		log.Fatalln(err)
	}

	err = rootCmd.MarkFlagRequired("output")
	if err != nil {
		log.Fatalln(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func expand(paths []string) ([]string, error) {
	result := make([]string, 0, len(paths))
	for _, path := range paths {
		matches, err := filepath.Glob(path)
		if err != nil {
			return nil, err
		}
		result = append(result, matches...)
	}
	return result, nil
}
