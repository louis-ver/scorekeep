package converter

import (
	"github.com/louis-ver/scorekeep/db"
	"github.com/louis-ver/scorekeep/dto"
)

func TeamToTeamDTO(teamDTO *dto.TeamDTO, team db.Team) {
	teamDTO.ID = team.ID
	teamDTO.Location = team.Location
	teamDTO.Name = team.Name
	teamDTO.TeamName = team.TeamName
}

func LeagueToLeagueDTO(leagueDTO *dto.LeagueDTO, league db.League, teams []db.Team) {
	leagueDTO.ID = league.ID
	leagueDTO.Name = league.Name
	leagueDTO.Sport = league.Sport
	leagueDTO.Acronym = league.Acronym

	var teamsDTO []dto.TeamDTO
	for _, element := range teams {
		var teamDTO dto.TeamDTO
		TeamToTeamDTO(&teamDTO, element)
		teamsDTO = append(teamsDTO, teamDTO)
	}
	leagueDTO.Teams = teamsDTO
}
