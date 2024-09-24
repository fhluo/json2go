package main

import (
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of json2go",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("json2go v0.5.0")
	},
}

func init() {
	versionCmd.SetOut(os.Stdout)

	rootCmd.AddCommand(versionCmd)
}
