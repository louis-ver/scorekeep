package nhl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DatesDTO struct {
	Dates []DateDTO `json:"dates"`
}

type DateDTO struct {
	Games []GameDTO `json:"games"`
}

type GameDTO struct {
	Teams ScheduleTeamDTO `json:"teams"`
}

type ScheduleTeamDTO struct {
	Away ScheduleRecordDTO `json:"away"`
	Home ScheduleRecordDTO `json:"home"`
}

type ScheduleRecordDTO struct {
	Score int     `json:"score"`
	Team  TeamDTO `json:"team"`
}

const schedulesEndpoint = "/schedule"

func GetScoresForDate(date string) DatesDTO {
	resp, err := http.Get(fmt.Sprintf("%s%s?startDate=%s&endDate=%s", ApiURL, schedulesEndpoint, date, date))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var dates DatesDTO
	err = decoder.Decode(&dates)
	if err != nil {
		panic(err)
	}
	return dates
}
