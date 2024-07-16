package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"main.go/cors"
)

func main() {
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	router.Use(cors.New(config))

	// Setup proxy
	target, _ := url.Parse("http://localhost:8080")
	proxy := httputil.NewSingleHostReverseProxy(target)
	router.Any("/*proxyPath", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("proxyPath")
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// Start server
	router.Run(":8081")
}
