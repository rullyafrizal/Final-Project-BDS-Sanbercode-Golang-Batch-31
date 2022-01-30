package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/config"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/docs"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/routes"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/seeds"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

var err error

func main() {
	// if os.Getenv("APP_ENV") == "local" {

	// }
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//programmatically set swagger info
	swaggerHost := os.Getenv("SWAGGER_HOST")

	if swaggerHost == "" {
		swaggerHost = "localhost:8080"
	}

	docs.SwaggerInfo.Title = "Swagger Blog API"
	docs.SwaggerInfo.Description = "This is a Blog Rest API Docs."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = swaggerHost
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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
