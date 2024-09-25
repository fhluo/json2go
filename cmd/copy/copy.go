package main

import (
	"fmt"
	"github.com/fhluo/json2go/pkg/util"
	"github.com/lmittmann/tint"
	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"time"
)

var (
	verbose     bool
	source      string
	destination string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().StringVarP(&source, "src", "s", "", "source directory")
	rootCmd.Flags().StringVarP(&destination, "dst", "d", "", "destination directory")

	_ = rootCmd.MarkFlagRequired("src")
	_ = rootCmd.MarkFlagRequired("dst")

	_ = rootCmd.MarkFlagDirname("src")
	_ = rootCmd.MarkFlagDirname("dst")
}

var rootCmd = &cobra.Command{
	Use:   "copy [path...]",
	Short: "Copy files and directories",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger := slog.New(util.DiscardHandler)
		if verbose {
			logger = slog.New(
				tint.NewHandler(os.Stderr, &tint.Options{
					Level:      slog.LevelInfo,
					TimeFormat: time.TimeOnly,
				}),
			)
		}
		slog.SetDefault(logger)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			Copy(source, destination)
		}

		for path := range slices.Values(args) {
			Copy(filepath.Join(source, path), filepath.Join(destination, path))
		}
	},
}

func Copy(src string, dst string) {
	err := cp.Copy(src, dst)
	if err != nil {
		fmt.Printf("Failed to copy files from %s to %s\n", src, dst)
		slog.Error("failed to copy files", "err", err, "src", src, "dst", dst)
		os.Exit(1)
	}

	slog.Info("copied", "src", src, "dst", dst)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
