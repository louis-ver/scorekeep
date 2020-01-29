package pkg

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

var teamMap = map[string]string{
	"ATL Hawks":         "Atlanta Hawks",
	"BKN Nets":          "Brooklyn Nets",
	"BOS Celtics":       "Boston Celtics",
	"CHA Hornets":       "Charlotte Hornets",
	"CHI Bulls":         "Chicago Bulls",
	"CLE Cavaliers":     "Cleveland Cavaliers",
	"DAL Mavericks":     "Dallas Mavericks",
	"DEN Nuggets":       "Denver Nuggets",
	"DET Pistons":       "Detroit Pistons",
	"GS Warriors":       "Golden State Warriors",
	"HOU Rockets":       "Houston Rockets",
	"IND Pacers":        "Indiana Pacers",
	"LA Clippers":       "Los Angeles Clippers",
	"LA Lakers":         "Los Angeles Lakers",
	"MEM Grizzlies":     "Memphis Grizzlies",
	"MIA Heat":          "Miami Heat",
	"MIL Bucks":         "Milwaukee Bucks",
	"MIN Timberwolves":  "Minnesota Timberwolves",
	"NO Pelicans":       "New Orleans Pelicans",
	"NY Knicks":         "New York Knicks",
	"OKC Thunder":       "Oklahoma City Thunder",
	"ORL Magic":         "Orlando Magic",
	"PHI 76ers":         "Philadelphia 76ers",
	"PHX Suns":          "Phoenix Suns",
	"POR Trail Blazers": "Portland Trail Blazers",
	"SA Spurs":          "San Antonio Spurs",
	"SAC Kings":         "Sacramento Kings",
	"TOR Raptors":       "Toronto Raptors",
	"UTA Jazz":          "Utah Jazz",
	"WAS Wizards":       "Washington Wizards",
}

type nba struct {
	host   string
	apikey string
}

func InitializeNBA(host string) Leaguer {
	return &nba{host: host}
}

func (n *nba) GetLeagueInformation() League {
	return League{
		Name:    "National Basketball Association",
		Acronym: "NBA",
		Sport:   "Basketball",
		Icon:    "🏀",
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
		var currentPeriodOrdinal, timeRemainingInPeriod string
		if strings.Contains(clock, "AM") || strings.Contains(clock, "PM") {
			timeRemainingInPeriod = clock
		} else if !strings.Contains(clock, "Final") {
			times := strings.Split(clock, " ")
			timeRemainingInPeriod = times[0]
			if len(times) > 1 {
				currentPeriodOrdinal = times[1]
			}
		} else if strings.Contains(clock, "OT") {
			currentPeriodOrdinal = "5th"
			timeRemainingInPeriod = "Final (OT)"
		} else if strings.Contains(clock, "Final") {
			currentPeriodOrdinal = "4th"
			timeRemainingInPeriod = "Final"
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
			CurrentPeriodOrdinal:  currentPeriodOrdinal,
			TimeRemainingInPeriod: timeRemainingInPeriod,
		}
		games = append(games, game)
	})

	c.Visit(fmt.Sprintf("%s/events/date/%s", n.host, date))

	return games
}
