package pkg

type Leaguer interface {
	GetScores(date string) []Game
	GetLeagueInformation() League
}

type League struct {
	Name    string
	Acronym string
	Sport   string
	Icon    string
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
