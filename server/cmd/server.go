package main

import (
	"net/http"

	"main.go/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	router.Use(cors.New(config))

	// Dummy route
	router.GET("/dummy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Server 1",
		})
	})

	// Start server
	router.Run(":8080")
}
