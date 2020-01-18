package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louis-ver/scorekeep/lib"
	"net/http"
	"os"
	"time"
)

func main() {
	r := gin.Default()
	nhl := lib.InitializeNHL(os.Getenv("NHL_API_HOST"))
	nba := lib.InitializeNBA(os.Getenv("NBA_API_HOST"), os.Getenv("NBA_API_KEY"))
	r.GET("/leagues/:league_name", func(c *gin.Context) {
		switch leagueName := c.Param("league_name"); leagueName {
		case "nhl":
			c.JSON(http.StatusOK, nhl.GetTeams())
		default:
			c.Status(http.StatusNotFound)
		}
	})

	r.GET("/leagues/:league_name/scores", func(c *gin.Context) {
		date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
		favorites := c.DefaultQuery("favorites", "")
		switch leagueName := c.Param("league_name"); leagueName {
		case "nhl":
			c.JSON(http.StatusOK, nhl.GetScores(date, []string{favorites}))
		case "nba":
			c.JSON(http.StatusOK, nba.GetScores(date))
		default:
			c.Status(http.StatusNotFound)
		}
	})

	r.Run(":8080")
}
