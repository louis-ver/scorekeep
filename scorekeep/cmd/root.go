package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/louis-ver/scorekeep/library"
	"github.com/spf13/cobra"
)

var (
	date string
)

const (
	ScorekeepServerURL = "http://localhost:8080"
)

type game struct {
	Home                  team
	Away                  team
	CurrentPeriodOrdinal  string
	TimeRemainingInPeriod string
}

type team struct {
	Name  string
	Score int
}

type League struct {
	Name  string
	Games []game
}

var rootCmd = &cobra.Command{
	Use:   "scorekeep",
	Short: "The easiest way to track scores across leagues",
	Long:  "Scorekeep is a CLI that enables you to track game scores across the most popular professional sports leagues",
	Run: func(cmd *cobra.Command, args []string) {
		nhlGames := fetchScores("nhl")
		nbaGames := fetchScores("nba")

		printScores([]League{
			{
				Name:  "National Hockey League",
				Games: nhlGames,
			},
			{
				Name:  "National Basketball Association",
				Games: nbaGames,
			},
		})

	},
}

func fetchScores(league string) []game {
	resp, err := http.Get(fmt.Sprintf("%s/leagues/%s/scores", ScorekeepServerURL, league))
	if err != nil {
		log.Fatal(err)
	}
	var games []game
	library.DecodeJSON(resp, &games)

	return games
}

func printScores(leagues []League) {
	w := tabwriter.NewWriter(os.Stdout, 10, 10, 3, ' ', 0)
	for _, league := range leagues {
		leagueUpperCase := strings.ToUpper(league.Name)
		var periodName string
		switch leagueUpperCase {
		case "NHL":
			periodName = "PERIOD"
		default:
			periodName = "QUARTER"
		}
		fmt.Fprintf(w, "%s\t\t\t\t\t\n", leagueUpperCase)
		fmt.Fprintf(w, "AWAY\tSCORE\tHOME\tSCORE\t%s\tTIME REMAINING\n", periodName)
		for _, element := range league.Games {
			fmt.Fprintf(w, "%s\t%d\t%s\t%d\t%s\t%s\n", element.Away.Name, element.Away.Score, element.Home.Name, element.Home.Score, element.CurrentPeriodOrdinal, element.TimeRemainingInPeriod)
		}
	}
	w.Flush()
}

func Execute() {
	rootCmd.Flags().StringVarP(&date, "date", "d", "", "Date to fetch scores (YYYY-mm-dd")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
