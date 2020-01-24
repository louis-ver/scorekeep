package library

import (
	"github.com/gocolly/colly"
	"strconv"
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

type data struct {
	Api api `json:"api"`
}

type api struct {
	Games []nbagame `json:"games"`
}

type nbagame struct {
	Home          nbateam `json:"hTeam"`
	Away          nbateam `json:"vTeam"`
	CurrentPeriod string  `json:"currentPeriod"`
	Status        string  `json:"statusGame"`
	Clock         string  `json:"clock"`
}

type nbateam struct {
	ID       string `json:"teamId"`
	FullName string `json:"fullName"`
	Name     string `json:"nickName"`
	Score    score  `json:"score"`
}

type score struct {
	Points string `json:"points"`
}

func (n *nba) GetScores(date string, favorites []string) []Game {
	return scrapeScores()
}

func scrapeScores() []Game {
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
		game := Game{
			Home: Team{
				Name:  teamMap[homeTeam],
				Score: homeScore,
			},
			Away: Team{
				Name:  teamMap[awayTeam],
				Score: awayScore,
			},
			CurrentPeriodOrdinal:  "",
			TimeRemainingInPeriod: clock,
		}
		games = append(games, game)
	})

	c.Visit("https://www.thescore.com/nba/events/date/2020-01-23")

	return games
}
