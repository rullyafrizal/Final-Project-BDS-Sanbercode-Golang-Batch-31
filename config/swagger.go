package config

import (
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/docs"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
)

func SetupSwagger() {
	docs.SwaggerInfo.Title = "Blog API"
	docs.SwaggerInfo.Description = "This is a Blog Rest API Documentation."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost")
	docs.SwaggerInfo.Schemes = []string{}

	if utils.Getenv("APP_ENV", "local") == "local" {
		docs.SwaggerInfo.Schemes = []string{"http"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}
}