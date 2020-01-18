package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louis-ver/scorekeep/lib"
	"log"
	"net/http"
	"os"
	"time"
)

type httperror struct {
	HttpCode int `json:"-"`
	Message string `json:"error"`
}

func main() {
	r := gin.Default()
	nhlhost := os.Getenv("NHL_API_HOST")
	if nhlhost == "" {
		log.Fatal("$NHL_API_HOST cannot be empty.")
	}
	nbahost, nbaapikey := os.Getenv("NBA_API_HOST"), os.Getenv("NBA_API_KEY")
	if nbahost == "" {
		log.Fatal("$NBA_API_HOST cannot be empty")
	}
	if nbaapikey == "" {
		log.Fatal("$NBA_API_KEY cannot be empty")
	}
	nhl := lib.InitializeNHL(nhlhost)
	nba := lib.InitializeNBA(nbahost, nbaapikey)

	r.GET("/leagues/:league_name/scores", func(c *gin.Context) {
		date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
		favorites := c.DefaultQuery("favorites", "")
		switch leagueName := c.Param("league_name"); leagueName {
		case "nhl":
			scoresResponse(c, nhl, date, []string{favorites})
		case "nba":
			scoresResponse(c, nba, date, []string{favorites})
		default:
			serverError(c, httperror{Message: "league not supported by scorekeep", HttpCode: http.StatusNotFound})
		}
	})

	r.Run(":8080")
}

func scoresResponse(c *gin.Context, l lib.League, date string, favorites []string) {
	c.JSON(http.StatusOK, l.GetScores(date, favorites))
}

func serverError(c *gin.Context, err httperror) {
	c.JSON(err.HttpCode, err)
}