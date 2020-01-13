package dto

type LeagueDTO struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Acronym string    `json:"acronym"`
	Sport   string    `json:"sport"`
	Teams   []TeamDTO `json:"teams,omitempty"`
}
