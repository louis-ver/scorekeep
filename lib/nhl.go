package lib

import (
	"fmt"
	"log"
	"net/http"
)

type nhlapi struct {
	host string
}

func InitializeNHL(host string) League {
	return &nhlapi{host: host}
}

type nhldates struct {
	Dates []nhldate `json:"dates"`
}

type nhldate struct {
	Games []nhlgame `json:"games"`
}

type nhlgame struct {
	ID int `json:"gamePk"`
}

func (n *nhlapi) GetScores(date string, favorites []string) []Game {
	resp, err := http.Get(fmt.Sprintf("%s%s?startDate=%s&endDate=%s", n.host, "/schedule", date, date))
	if err != nil {
		panic(err)
	}
	var nhldates nhldates
	decodeJSON(resp, &nhldates)
	if err != nil {
		panic(err)
	}
	var scores []Game
	games := nhldates.Dates[0]
	// TODO: Rewrite this in parallel
	for _, element := range games.Games {
		game := n.GetGameState(&element)

		scores = append(scores, game)
	}
	return scores
}

type nhlgamedetail struct {
	CurrentPeriodOrdinal string   `json:"currentPeriodOrdinal"`
	PeriodTimeRemaining  string   `json:"currentPeriodTimeRemaining"`
	Teams                nhlteams `json:"teams"`
}

type nhlteams struct {
	Home nhlteam `json:"home"`
	Away nhlteam `json:"away"`
}

type nhlteam struct {
	TeamInfo nhlteaminfo `json:"team"`
	Goals    int         `json:"goals"`
}

type nhlteaminfo struct {
	Name string `json:"name"`
}

func (n *nhlapi) GetGameState(game *nhlgame) Game {
	resp, err := http.Get(fmt.Sprintf("%s/game/%d/linescore", n.host, game.ID))
	if err != nil {
		log.Fatal(err)
	}
	var nhlgamedetail nhlgamedetail
	decodeJSON(resp, &nhlgamedetail)
	homeTeam := Team{
		Name:  nhlgamedetail.Teams.Home.TeamInfo.Name,
		Score: nhlgamedetail.Teams.Home.Goals,
	}
	awayTeam := Team{
		Name:  nhlgamedetail.Teams.Away.TeamInfo.Name,
		Score: nhlgamedetail.Teams.Away.Goals,
	}
	return Game{
		Home:                  homeTeam,
		Away:                  awayTeam,
		CurrentPeriodOrdinal:  nhlgamedetail.CurrentPeriodOrdinal,
		TimeRemainingInPeriod: nhlgamedetail.PeriodTimeRemaining,
	}
}
