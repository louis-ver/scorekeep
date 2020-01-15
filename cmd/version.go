package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Scorekeep",
	Long:  "All software has versions. This is Scorekeep's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scorekeep v0.1")
	},
}
