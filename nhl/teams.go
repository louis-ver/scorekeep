package nhl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var nhlServer = "https://statsapi.web.nhl.com/api/v1"

type Teams struct {
	Team []Team `json:"teams"`
}

type Team struct {
	ID       int    `json:"id"`
	FullName string `json:"name"`
	Name     string `json:"teamName"`
	Location string `json:"locationName"`
}

func GetTeams() Teams {
	resp, err := http.Get(fmt.Sprintf("%s/teams", nhlServer))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var teams Teams
	err = decoder.Decode(&teams)
	if err != nil {
		panic(err)
	}
	return teams
}
