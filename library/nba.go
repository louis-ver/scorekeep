package library

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

var teamMap = map[string]string{
	"SAC Kings":         "Sacramento Kings",
	"DET Pistons":       "Detroit Pistons",
	"OKC Thunder":       "Oklahoma City Thunder",
	"ORL Magic":         "Orlando Magic",
	"PHI 76ers":         "Philadelphia 76ers",
	"TOR Raptors":       "Toronto Raptors",
	"LA Clippers":       "Los Angeles Clippers",
	"ATL Hawks":         "Atlanta Hawks",
	"MEM Grizzlies":     "Memphis Grizzlies",
	"BOS Celtics":       "Boston Celtics",
	"WAS Wizards":       "Washington Wizards",
	"MIA Heat":          "Miami Heat",
	"LA Lakers":         "Los Angeles Lakers",
	"NY Knicks":         "New York Knicks",
	"MIN Timberwolves":  "Minnesota Timberwolves",
	"CHI Bulls":         "Chicago Bulls",
	"CLE Cavaliers":     "Cleveland Cavaliers",
	"POR Trail Blazers": "Portland Trail Blazers",
	"BKN Nets":          "Brooklyn Nets",
	"DEN Nuggets":       "Denver Nuggets",
	"DAL Mavericks":     "Dallas Mavericks",
	"HOU Rockets":       "Houston Rockets",
	"IND Pacers":        "Indiana Pacers",
	"PHX Suns":          "Phoenix Suns",
	"SA Spurs":          "San Antonio Spurs",
	"NO Pelicans":       "New Orleans Pelicans",
	"UTA Jazz":          "Utah Jazz",
	"GS Warriors":       "Golden State Warriors",
}

type nba struct {
	host   string
	apikey string
}

func InitializeNBA(host, apiKey string) Leaguer {
	return &nba{host: host, apikey: apiKey}
}

func (n *nba) GetLeagueInformation() League {
	return League{
		Name:    "National Basketball Association",
		Acronym: "NBA",
		Sport:   "Basketball",
	}
}

func (n *nba) GetScores(date string) []Game {
	return n.scrapeScores(date)
}

func (n *nba) scrapeScores(date string) []Game {
	var games []Game
	c := colly.NewCollector()

	c.OnHTML(".EventCard__eventCardContainer--3hTGN", func(e *colly.HTMLElement) {
		teams := e.ChildTexts(".EventCard__teamName--JweK5")
		scores := e.ChildTexts(".EventCard__scoreColumn--2JZbq")
		clock := e.ChildText(".EventCard__clockColumn--3lEPz")
		homeTeam := teams[0]
		homeScore, _ := strconv.Atoi(scores[0])
		awayTeam := teams[1]
		awayScore, _ := strconv.Atoi(scores[1])
		var currentPeriod, timeRemaining string
		currentPeriod = clock
		if !strings.Contains(clock, "Final") {
			times := strings.Split(clock, " ")
			currentPeriod = times[1]
			timeRemaining = times[0]
		}
		game := Game{
			Home: Team{
				Name:  teamMap[homeTeam],
				Score: homeScore,
			},
			Away: Team{
				Name:  teamMap[awayTeam],
				Score: awayScore,
			},
			CurrentPeriodOrdinal:  currentPeriod,
			TimeRemainingInPeriod: timeRemaining,
		}
		games = append(games, game)
	})

	c.Visit(fmt.Sprintf("%s/events/date/%s", n.host, date))

	return games
}
