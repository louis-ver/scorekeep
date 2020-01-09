package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/louis-ver/scorekeep/nhl"
)

func main() {
	nhlTeamHandler := func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(nhl.GetTeams())
	}

	http.HandleFunc("/teams", nhlTeamHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
