package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/louis-ver/scorekeep/pkg"
	"github.com/spf13/cobra"
)

var (
	date string
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
	Icon  string
	Games []game
}

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "The easiest way to track scores across leagues",
	Long:  "Scorekeep is a CLI that enables you to track game scores across the most popular professional sports leagues",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan League)

		go fetchScores(League{
			Name:  "nhl",
			Icon:  "🏒",
			Games: nil,
		}, date, c)
		go fetchScores(League{
			Name:  "nba",
			Icon:  "🏀",
			Games: nil,
		}, date, c)

		l1, l2 := <-c, <-c
		printScores([]League{l1, l2})

	},
}

func fetchScores(league League, date string, c chan League) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	resp, err := http.Get(fmt.Sprintf("%s/leagues/%s/scores?date=%s", GetConfig().ServerUrl, league.Name, date))
	if err != nil {
		log.Fatal(err)
	}
	var games []game
	pkg.DecodeJSON(resp, &games)

	var favorites []game
	var regulars []game

	for _, game := range games {
		if teamInFavorites(game.Away.Name, league.Name) || teamInFavorites(game.Home.Name, league.Name) {
			favorites = append(favorites, game)
		} else {
			regulars = append(regulars, game)
		}
	}
	games = append(favorites, regulars...)

	l := League{
		Name:  league.Name,
		Icon:  league.Icon,
		Games: games,
	}

	c <- l
}

func teamInFavorites(team, league string) bool {
	config := GetConfig()
	favorites := config.GetFavorites(league)
	return pkg.StringInSlice(team, favorites)

}

func printScores(leagues []League) {
	w := tabwriter.NewWriter(os.Stdout, 10, 10, 3, ' ', 0)
	for _, league := range leagues {
		leagueUpperCase := strings.ToUpper(league.Name)
		var periodName string
		switch leagueUpperCase {
		case "NATIONAL HOCKEY LEAGUE":
			periodName = "PERIOD"
		default:
			periodName = "QUARTER"
		}
		if len(league.Games) > 0 {
			fmt.Fprintf(w, "%s %s\t\t\t\t\t\n", league.Icon, leagueUpperCase)
			fmt.Fprintf(w, "AWAY\tSCORE\tHOME\tSCORE\t%s\tTIME\n", periodName)
			for _, element := range league.Games {
				fmt.Fprintf(w, "%s\t%d\t%s\t%d\t%s\t%s\n", element.Away.Name, element.Away.Score, element.Home.Name, element.Home.Score, element.CurrentPeriodOrdinal, element.TimeRemainingInPeriod)
			}
		}
	}
	w.Flush()
}

func Execute() {
	rootCmd.Flags().StringVarP(&date, "date", "d", "", "Date to fetch scores (YYYY-mm-dd)")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
