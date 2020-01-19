package library

type Leaguer interface {
	GetScores(date string, favorites []string) []Game
	GetLeagueInformation() League
}

type League struct {
	Name    string
	Acronym string
	Sport   string
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
