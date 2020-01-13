package db

type League struct {
	ID      int
	Name    string
	Acronym string
	Sport   string
}

func GetLeagues() []League {
	db := Open()
	defer db.Close()
	var leagues []League
	db.Find(&leagues)
	return leagues
}

func GetLeague(id int) League {
	db := Open()
	defer db.Close()
	var league League
	db.First(&league, id)
	return league
}
