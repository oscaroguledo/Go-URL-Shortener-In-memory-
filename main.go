package main

import (
	"log"
	"url-shortener/core"
	"url-shortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if err := core.InitDB(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err := core.Close(); err != nil {
			log.Println("Error closing database:", err)
		}
	}()

	// Initialize routes
	initGroup := r.Group("/")
	routes.InitRoutes(initGroup)

	// Replace with your proxy IP or CIDR range
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// routes.RegisterURLRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}

	log.Println("Server running on http://localhost:8080")
}
