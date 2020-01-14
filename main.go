package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/favorites", func(c *gin.Context) {
		nhl := initialize()
		nhl.AddFavorite("Calgary Flames")
	})

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
		config := GetConfig()
		switch leagueName := c.Param("league_name"); leagueName {
		case "nhl":
			nhl := initialize()
			c.JSON(http.StatusOK, nhl.GetScores(date, config.Favorites.NHL))
		default:
			c.Status(http.StatusNotFound)
		}
	})

	return r
}

func main() {
	r := setupRouter()
	touchConfigFile()
	r.Run(":8080")
}
