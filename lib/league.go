package lib

type League interface {
	GetTeams() []string
	GetScore() []Game
}

type Game struct {
	Home Team
	Away Team
}

type Team struct {
	Name  string
	Score int
}