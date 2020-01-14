package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/louis-ver/scorekeep/util"
)

type nhlapi struct {
	address string
}

func initialize() *nhlapi {
	return &nhlapi{address: "https://statsapi.web.nhl.com/api/v1"}
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
	resp, err := http.Get(fmt.Sprintf("%s%s", n.address, "/teams"))
	if err != nil {
		log.Fatal(err)
	}
	var teams nhlteams
	err = util.DecodeJSON(resp, &teams)
	if err != nil {
		log.Fatal(err)
	}
	var teamNames []string
	for _, element := range teams.Teams {
		teamNames = append(teamNames, element.FullName)
	}
	return teamNames
}

type Game struct {
	Home Team
	Away Team
}

type Team struct {
	Name  string
	Score int
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

func (n *nhlapi) GetScores(date string) []Game {
	resp, err := http.Get(fmt.Sprintf("%s%s?startDate=%s&endDate=%s", n.address, "/schedule", date, date))
	if err != nil {
		panic(err)
	}
	var nhldates nhldates
	util.DecodeJSON(resp, &nhldates)
	if err != nil {
		panic(err)
	}
	var scores []Game
	games := nhldates.Dates[0]
	for _, element := range games.Games {
		home := Team{Name: element.Teams.Home.Team.FullName, Score: element.Teams.Home.Score}
		away := Team{Name: element.Teams.Away.Team.FullName, Score: element.Teams.Away.Score}
		scores = append(scores, Game{Home: home, Away: away})
	}
	return scores
}
