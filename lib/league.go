package lib

type League interface {
	GetScores(date string, favorites []string) []Game
	GetTeams() []string
}

type Game struct {
	Home Team
	Away Team
}

type Team struct {
	Name  string
	Score int
}