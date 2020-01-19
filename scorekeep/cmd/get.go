package cmd

import (
	"fmt"
	"github.com/louis-ver/scorekeep/library"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(leaguesCmd)
	rootCmd.AddCommand(getCmd)
}

type league struct {
	Name    string `json:"Name"`
	Acronym string `json:"Acronym"`
	Sport   string `json:"Sport"`
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
		resp, err := http.Get(ScorekeepServerURL + "/leagues")
		if err != nil {
			log.Fatal(err)
		}
		var leagues []league
		library.DecodeJSON(resp, &leagues)
		printLeagues(leagues)
	},
}

func printLeagues(leagues []league) {
	w := tabwriter.NewWriter(os.Stdout, 10, 10, 3, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintln(w, "NAME\tACRONYM\tSPORT")
	for _, element := range leagues {
		fmt.Fprintf(w, "%s\t%s\t%s\t\n", element.Name, element.Acronym, element.Sport)
	}
	w.Flush()
}
