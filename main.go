package main

import (
	"os"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jcyrille972/go_api_test/models"
	"github.com/jcyrille972/go_api_test/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	mode := os.Getenv("EnvGinMode")
	gin.SetMode(mode)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	// Its great to version your API's

	v1 := r.Group("/api/v1")
	r.Use(cors.Default())
	routes.AddRoutes(v1)
	// Handle error response when a route is not defined
	r.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.Run()
}
