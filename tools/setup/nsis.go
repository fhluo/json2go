package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"os"
)

var nsisCmd = &cobra.Command{
	Use:   "nsis",
	Short: "setup nsis",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		name := getNSISDownloadFilename(nsisVersion)
		url := getNSISDownloadURL(nsisVersion)

		switch _, err := os.Stat(name); {
		case os.IsNotExist(err):
			if err = download(url, name); err != nil {
				slog.Error("failed to download NSIS", err)
				os.Exit(1)
			}
		case err == nil:
			slog.Info("file already exists, skip download")
		default:
			slog.Error("failed to stat file", err)
			os.Exit(1)
		}

		if err := extract(context.Background(), name, nsisPath); err != nil {
			slog.Error("failed to extract files", err, "name", name)
			os.Exit(1)
		}
	},
}

var (
	nsisPath    string
	nsisVersion string
)

func init() {
	rootCmd.AddCommand(nsisCmd)

	nsisCmd.Flags().StringVarP(&nsisPath, "path", "p", ".", "path to extract")
	nsisCmd.Flags().StringVarP(&nsisVersion, "version", "v", "3.08", "nsis version")
}

func getNSISDownloadURL(version string) string {
	return fmt.Sprintf("https://sourceforge.net/projects/nsis/files/NSIS%%203/%[1]s/nsis-%[1]s.zip/download", version)
}

func getNSISDownloadFilename(version string) string {
	return fmt.Sprintf("nsis-%s.zip", version)
}
