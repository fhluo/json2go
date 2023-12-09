package cmd

import (
	"os"

	"github.com/fhluo/json2go/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of json2go",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("json2go v" + version.Get().String())
	},
}

func init() {
	versionCmd.SetOut(os.Stdout)

	rootCmd.AddCommand(versionCmd)
}
