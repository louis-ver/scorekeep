package nhl

import (
	"fmt"
	"net/http"

	"github.com/louis-ver/scorekeep/util"
)

type TeamsDTO struct {
	Teams []TeamDTO `json:"teams"`
}

type TeamDTO struct {
	ID       int    `json:"id"`
	FullName string `json:"name"`
	Name     string `json:"teamName"`
	Location string `json:"locationName"`
}

const teamsEndpoint = "/teams"

func GetTeams() TeamsDTO {
	resp, err := http.Get(fmt.Sprintf("%s%s", ApiURL, teamsEndpoint))
	if err != nil {
		panic(err)
	}
	var teams TeamsDTO
	err = util.DecodeJSON(resp, &teams)
	if err != nil {
		panic(err)
	}
	return teams
}
