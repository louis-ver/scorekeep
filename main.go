package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/leagues", func(c *gin.Context) {
		c.JSON(http.StatusOK, []string{"nhl"})
	})

	r.GET("/leagues/:league_name", func(c *gin.Context) {
		switch leagueName := c.Param("league_name"); leagueName {
		case "nhl":
			nhl := initialize()
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
			nhl := initialize()
			c.JSON(http.StatusOK, nhl.GetScores(date, []string{favorites}))
		default:
			c.Status(http.StatusNotFound)
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
