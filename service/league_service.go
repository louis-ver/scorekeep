package service

import (
	"github.com/louis-ver/scorekeep/converter"
	"github.com/louis-ver/scorekeep/db"
	"github.com/louis-ver/scorekeep/dto"
)

func GetLeagues() []dto.LeagueDTO {
	leagues := db.GetLeagues()

	var leaguesDTO []dto.LeagueDTO
	for _, element := range leagues {
		var leagueDTO dto.LeagueDTO
		converter.LeagueToLeagueDTO(&leagueDTO, element, nil)
		leaguesDTO = append(leaguesDTO, leagueDTO)
	}
	return leaguesDTO
}

func GetLeague(id int) dto.LeagueDTO {
	league := db.GetLeague(id)
	teams := db.GetTeamsForLeague(id)
	var leagueDTO dto.LeagueDTO
	converter.LeagueToLeagueDTO(&leagueDTO, league, teams)

	return leagueDTO
}
