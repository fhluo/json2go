package main

import (
	"github.com/fhluo/json2go/pkg/downloaders"
	"github.com/spf13/cobra"
)

var nsisCmd = &cobra.Command{
	Use:   "nsis",
	Short: "setup nsis",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		installNSIS(nsisVersion, nsisPath)
	},
}

var (
	nsisPath    string
	nsisVersion string
)

func init() {
	rootCmd.AddCommand(nsisCmd)

	nsisCmd.Flags().StringVarP(&nsisPath, "path", "p", ".", "path to extract")
	nsisCmd.Flags().StringVarP(&nsisVersion, "version", "v", downloaders.DefaultNSISVersion, "nsis version")
}

func installNSIS(version string, installPath string) {
	install(downloaders.NewNSISDownloader(version), installPath)
}
