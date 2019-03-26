package main

import (
	c "github.com/kardasjan/deepa/controller"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8000" //failsafe
	}

	r := gin.Default()
	c := controller.NewController()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE", "HEAD"}

	r.Use(cors.New(config))

	routes := r.Group("/")
	{
		geovision = r.Group("/geovision") {
			geovision.POST("/decode", c.DecodeEmail)
		}
		ip150 = r.Group("/ip150") {
			geovision.POST("/decode", c.DecodeEmail)
		}
	}

	r.Run(":" + port)
}
