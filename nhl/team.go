package nhl

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var teams TeamsDTO
	err = decoder.Decode(&teams)
	if err != nil {
		panic(err)
	}
	return teams
}
