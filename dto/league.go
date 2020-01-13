package dto

type LeagueDTO struct {
	ID      int
	Name    string
	Acronym string
	Sport   string
	Teams   []TeamDTO
}
