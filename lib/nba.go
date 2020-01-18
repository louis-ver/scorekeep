package lib

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type nba struct {
	host   string
	apikey string
}

func InitializeNBA(host, apiKey string) League {
	return &nba{host: host, apikey: apiKey}
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
	games := fmt.Sprintf("/games/date/%s", date)
	req, err := http.NewRequest("GET", n.host+games, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-RapidAPI-Key", n.apikey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var nbaGames data
	decodeJSON(resp, &nbaGames)
	var scores []Game
	for _, element := range nbaGames.Api.Games {
		homeScore, err := strconv.Atoi(element.Home.Score.Points)
		if err != nil {
			homeScore = 0
		}
		awayScore, err := strconv.Atoi(element.Away.Score.Points)
		if err != nil {
			homeScore = 0
		}
		convertToFinishedGame(&element)
		scores = append(
			scores,
			Game{
				Home:                  Team{Name: element.Home.FullName, Score: homeScore},
				Away:                  Team{Name: element.Away.FullName, Score: awayScore},
				CurrentPeriodOrdinal:  nbaCurrentPeriodToGameCurrentPeriod(element.CurrentPeriod),
				TimeRemainingInPeriod: element.Clock,
			})
	}
	return scores
}

func convertToFinishedGame(nbaGame *nbagame) {
	if (nbaCurrentPeriodToGameCurrentPeriod(nbaGame.CurrentPeriod) == "4th" || nbaCurrentPeriodToGameCurrentPeriod(nbaGame.CurrentPeriod) == "OT") && nbaGame.Clock == "" {
		nbaGame.Clock = "Final"
	}
}

func nbaCurrentPeriodToGameCurrentPeriod(nbaCurrentPeriod string) string {
	currentPeriod, err := strconv.Atoi(nbaCurrentPeriod[:1])
	if err != nil {
		log.Fatal(err)
	}
	if currentPeriod == 1 {
		return "1st"
	} else if currentPeriod == 2 {
		return "2nd"
	} else if currentPeriod == 3 {
		return "3rd"
	} else if currentPeriod == 4 {
		return "4th"
	} else if currentPeriod >= 5 {
		return "OT"
	} else {
		return ""
	}
}
