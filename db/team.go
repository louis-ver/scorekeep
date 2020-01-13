package db

type Team struct {
	ID         int
	ExternalID int
	League     int
	Link       string
	Location   string
	Name       string
}

func getTeamsForLeague(league int) []Team {
	db := Open()
	var teams []Team
	db.Where("league = ?", league).Find(&teams)
	return teams
}
