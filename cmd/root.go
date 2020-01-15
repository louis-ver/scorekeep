package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/louis-ver/scorekeep/api"
	"github.com/spf13/cobra"
)

var (
	date string
)

var rootCmd = &cobra.Command{
	Use:   "scorekeep",
	Short: "The best way to track scores across leagues",
	Long:  "Scorekeep is a CLI that enables you to track game scores across all professional sports leagues",
	Run: func(cmd *cobra.Command, args []string) {
		nhl := api.Initialize()
		games := nhl.GetScores(date, []string{})
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.DiscardEmptyColumns)
		fmt.Fprintln(w, "AWAY\t\tHOME\t\t")
		fmt.Fprintln(w, formatScore(games))
	},
}

func formatScore(games []api.Game) string {
	var str strings.Builder
	for _, element := range games {
		str.WriteString(fmt.Sprintf("%s\t%d\t%s\t%d\n", element.Away.Name, element.Away.Score, element.Home.Name, element.Home.Score))
	}
	return str.String()
}

func Execute() {
	rootCmd.Flags().StringVarP(&date, "date", "d", "", "Date to fetch scores (YYYY-mm-dd")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
