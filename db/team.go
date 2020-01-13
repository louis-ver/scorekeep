package db

type Team struct {
	ID         int
	ExternalID int
	LeagueID   int
	Link       string
	Location   string
	Name       string
	TeamName   string
}

func GetTeamsForLeague(league int) []Team {
	db := Open()
	var teams []Team
	db.Where("league_id = ?", league).Find(&teams)
	return teams
}
