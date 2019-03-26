package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database //database

func init() {

	err := godotenv.Load()

	if err != nil && os.Getenv("APP_ENV") == "development" {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Connect to Database
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Print(err)
	}

	db = session.DB(dbName)
}

// GetDB returns a handle to the DB Session object
func GetDB() *mgo.Database {
	return db
}
