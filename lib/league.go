package lib

type League interface {
	GetScores(date string, favorites []string) []Game
}

type Game struct {
	Home                  Team
	Away                  Team
	CurrentPeriodOrdinal  string
	TimeRemainingInPeriod string
}

type Team struct {
	Name  string
	Score int
}
