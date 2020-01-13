package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/louis-ver/scorekeep/db"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/leagues", func(c *gin.Context) {
		c.JSON(http.StatusOK, db.GetLeagues())
	})

	r.GET("/leagues/:id", func(c *gin.Context) {
		leagueID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, db.GetLeague(leagueID))
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
