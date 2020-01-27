package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louis-ver/scorekeep/library"
	"net/http"
	"os"
	"time"
)

type httperror struct {
	HttpCode int    `json:"-"`
	Message  string `json:"error"`
}

func main() {
	nhlHost := os.Getenv("NHL_API_HOST")
	nbaHost, nbaApiKey := os.Getenv("NBA_API_HOST"), os.Getenv("NBA_API_KEY")

	apis := map[string]library.Leaguer{
		"nhl": library.InitializeNHL(nhlHost),
		"nba": library.InitializeNBA(nbaHost, nbaApiKey),
	}

	r := gin.Default()

	r.GET("/leagues", func(c *gin.Context) {
		leaguesResponse(c, apis)
	})

	r.GET("/leagues/:league_name/scores", func(c *gin.Context) {
		date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))

		leagueApi, ok := apis[c.Param("league_name")]
		if !ok {
			serverError(c, httperror{Message: "league not supported by scorekeep", HttpCode: http.StatusNotFound})
		} else {
			scoresResponse(c, leagueApi, date)
		}
	})

	r.Run(":8080")
}

func scoresResponse(c *gin.Context, l library.Leaguer, date string) {
	c.JSON(http.StatusOK, l.GetScores(date))
}

func leaguesResponse(c *gin.Context, leagueMap map[string]library.Leaguer) {
	var supportedLeagues []library.League
	for _, leaguer := range leagueMap {
		supportedLeagues = append(supportedLeagues, leaguer.GetLeagueInformation())
	}
	c.JSON(http.StatusOK, supportedLeagues)
}

func serverError(c *gin.Context, err httperror) {
	c.JSON(err.HttpCode, err)
}
