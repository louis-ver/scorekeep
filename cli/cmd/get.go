package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(leaguesCmd)
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get various types of content",
	Long: `Get various types of content.
		
	Get requires a subcommand, e.g. ` + "`scorekeep get leagues`.",
	Run: nil,
}

var leaguesCmd = &cobra.Command{
	Use:   "leagues",
	Short: "Get all leagues",
	Long:  "Get all leagues supported by Scorekeep",
	Run: func(cmd *cobra.Command, args []string) {
		printSupportedLeagues()
	},
}

func printSupportedLeagues() {
	w := tabwriter.NewWriter(os.Stdout, 10, 10, 3, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintln(w, "NAME\tACRONYM\tSPORT")
	// for _, element := range cli.GetSupportedLeagues() {
	// 	fmt.Fprintf(w, "%s\t%s\t%s\t\n", element.Name, element.Acronym, element.Sport)
	// }
	w.Flush()
}
