package nhl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Dates struct {
	Date []Date `json:"dates"`
}

type Date struct {
	ScheduleTeam []ScheduleTeam `json:"games"`
}

type ScheduleTeam struct {
	Game Game `json:"teams"`
}

type Game struct {
	Away Record `json:"away"`
	Home Record `json:"home"`
}

type Record struct {
	Score int  `json:"score"`
	Team  Team `json:"team"`
}

const schedulesEndpoint = "/schedule"

func GetScoresForDate(date string) Dates {
	resp, err := http.Get(fmt.Sprintf("%s%s?startDate=%s&endDate=%s", ApiURL, schedulesEndpoint, date, date))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var dates Dates
	err = decoder.Decode(&dates)
	if err != nil {
		panic(err)
	}
	return dates
}
