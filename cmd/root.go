package cmd

import (
	"fmt"
	"io"
	"os"
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
		w := tabwriter.NewWriter(os.Stdout, 10, 10, 3, ' ', tabwriter.DiscardEmptyColumns)
		printScores(w, games)
		w.Flush()
	},
}

func printScores(w io.Writer, games []api.Game) {
	fmt.Fprintln(w, "AWAY\tSCORE\tHOME\tSCORE")
	for _, element := range games {
		fmt.Fprintf(w, "%s\t%d\t%s\t%d\n", element.Away.Name, element.Away.Score, element.Home.Name, element.Home.Score)
	}
}

func Execute() {
	rootCmd.Flags().StringVarP(&date, "date", "d", "", "Date to fetch scores (YYYY-mm-dd")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
