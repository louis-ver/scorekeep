package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/louis-ver/scorekeep/nhl"
)

func main() {
	nhlTeamHandler := func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(nhl.GetTeams())
	}

	nhlScheduleHandler := func(w http.ResponseWriter, req *http.Request) {
		date := req.URL.Query().Get("date")
		if date == "" {
			date = time.Now().Format("2006-01-02")
		}
		json.NewEncoder(w).Encode(nhl.GetScoresForDate(date))
	}

	http.HandleFunc("/nhl/teams", nhlTeamHandler)
	http.HandleFunc("/nhl/schedule", nhlScheduleHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
