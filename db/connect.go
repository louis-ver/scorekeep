package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

const (
	dbUser     = "postgres"
	dbPassword = "password"
	dbName     = "postgres"
)

func Open() *gorm.DB {
	parameters := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword)
	db, err := gorm.Open("postgres", parameters)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
