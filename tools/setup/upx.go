package main

import (
	"github.com/fhluo/json2go/pkg/downloaders"
	"github.com/spf13/cobra"
)

var upxCmd = &cobra.Command{
	Use:   "upx",
	Short: "setup upx",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		installUPX(upxVersion, githubToken, upxPath)
	},
}

var (
	upxPath    string
	upxVersion string
)

func init() {
	rootCmd.AddCommand(upxCmd)

	upxCmd.Flags().StringVarP(&upxPath, "path", "p", ".", "path to extract upx")
	upxCmd.Flags().StringVarP(&upxVersion, "version", "v", downloaders.DefaultUPXVersion, "upx version")
}

func installUPX(version string, githubToken string, installPath string) {
	downloader := downloaders.NewUPXDownloader(version, githubToken)
	install(downloader, installPath)
}
