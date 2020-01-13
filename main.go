package main

import (
	"net/http"
	"strconv"

	"github.com/louis-ver/scorekeep/dto"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/louis-ver/scorekeep/service"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/leagues", func(c *gin.Context) {
		var leagues []dto.LeagueDTO
		leagues = service.GetLeagues()
		c.JSON(http.StatusOK, leagues)
	})

	r.GET("/leagues/:id", func(c *gin.Context) {
		leagueID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(http.StatusNotFound)
		} else {
			league := service.GetLeague(leagueID)
			c.JSON(http.StatusOK, league)
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
