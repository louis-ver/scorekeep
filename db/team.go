package db

type Team struct {
	ID         int
	ExternalID int
	LeagueID   int
	Link       string
	Location   string
	Name       string
}

func GetTeamsForLeague(league int) []Team {
	db := Open()
	var teams []Team
	db.Where("league = ?", league).Find(&teams)
	return teams
}
