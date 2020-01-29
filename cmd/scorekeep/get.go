package main

import (
	"fmt"
	"github.com/louis-ver/scorekeep/pkg"
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
	Icon    string `json:"Icon"`
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
		leagues := fetchSupportedLeagues()
		printLeagues(leagues)
	},
}

func fetchSupportedLeagues() []league {
	resp, err := http.Get(GetConfig().ServerUrl + "/leagues")
	if err != nil {
		log.Fatal(err)
	}
	var leagues []league
	pkg.DecodeJSON(resp, &leagues)

	return leagues
}

func printLeagues(leagues []league) {
	w := tabwriter.NewWriter(os.Stdout, 10, 10, 3, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintln(w, "NAME\tACRONYM\tSPORT\tICON")
	for _, element := range leagues {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", element.Name, element.Acronym, element.Sport, element.Icon)
	}
	w.Flush()
}
