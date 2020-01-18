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

type nhlteams struct {
	Teams []nhlteam `json:"teams"`
}

type nhlteam struct {
	ID       int    `json:"id"`
	FullName string `json:"name"`
	Name     string `json:"teamName"`
	Location string `json:"locationName"`
}

func (n *nhlapi) GetTeams() []string {
	resp, err := http.Get(fmt.Sprintf("%s%s", n.host, "/teams"))
	if err != nil {
		log.Fatal(err)
	}
	var teams nhlteams
	err = decodeJSON(resp, &teams)
	if err != nil {
		log.Fatal(err)
	}
	var teamNames []string
	for _, element := range teams.Teams {
		teamNames = append(teamNames, element.FullName)
	}
	return teamNames
}

type nhldates struct {
	Dates []nhldate `json:"dates"`
}

type nhldate struct {
	Games []nhlgame `json:"games"`
}

type nhlgame struct {
	Teams nhlscheduleteam `json:"teams"`
}

type nhlscheduleteam struct {
	Away nhlschedulerecord `json:"away"`
	Home nhlschedulerecord `json:"home"`
}

type nhlschedulerecord struct {
	Score int     `json:"score"`
	Team  nhlteam `json:"team"`
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
	for _, element := range games.Games {
		home := Team{Name: element.Teams.Home.Team.FullName, Score: element.Teams.Home.Score}
		away := Team{Name: element.Teams.Away.Team.FullName, Score: element.Teams.Away.Score}
		game := Game{Home: home, Away: away}
		homeResourceName := teamNameToResourceName(home.Name)
		awayResourceName := teamNameToResourceName(away.Name)
		if StringInSlice(homeResourceName, favorites) || StringInSlice(awayResourceName, favorites) {
			scores = append([]Game{game}, scores...)
		} else {
			scores = append(scores, Game{Home: home, Away: away})
		}
	}
	return scores
}
