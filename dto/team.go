package dto

type TeamDTO struct {
	ID       int    `json:"id"`
	Location string `json:"location"`
	Name     string `json:"name"`
	TeamName string `json:"team_name"`
}
