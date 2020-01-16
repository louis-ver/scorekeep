package api

type League struct {
	Name    string
	Acronym string
	Sport   string
}

func GetSupportedLeagues() []League {
	return []League{
		League{Name: "National Hockey League", Acronym: "NHL", Sport: "Hockey"},
		League{Name: "National Basketball Association", Acronym: "NBA", Sport: "Basketball"},
	}
}
