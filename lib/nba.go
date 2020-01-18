package lib

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type nba struct {
	host string
	apikey string
}

func InitializeNBA(host, apiKey string) *nba {
	return &nba{host: host, apikey: apiKey}
}

type data struct {
	Api api `json:"api"`
}

type api struct {
	Games []nbagame `json:"games"`
}

type nbagame struct {
	Home nbateam `json:"hTeam"`
	Away nbateam `json:"vTeam"`
}

type nbateam struct {
	ID string `json:"teamId"`
	FullName string `json:"fullName"`
	Name string `json:"nickName"`
	Score score `json:"score"`
}

type score struct {
	Points string `json:"points"`
}

func (n *nba) GetScores(date string) []Game {
	games := fmt.Sprintf("/games/date/%s", date)
	req, err := http.NewRequest("GET", n.host + games, nil)
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
		scores = append(
			scores,
			Game{
				Home: Team{Name: element.Home.FullName, Score: homeScore},
				Away: Team{Name: element.Away.FullName, Score: awayScore},
			})
	}
	return scores
}