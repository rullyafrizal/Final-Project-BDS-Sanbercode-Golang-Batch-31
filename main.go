package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/config"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/routes"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/seeds"
)

var err error

func main() {
	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectPostgres()
	sqlDb, err := db.DB()

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	err = seeds.Seed(db)

	if err != nil {
		log.Fatal("Error seeding database")
	}

	routes.SetupRouter(db)

	defer sqlDb.Close()
}
