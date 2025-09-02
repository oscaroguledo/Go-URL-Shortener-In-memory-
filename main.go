package main

import (
	"url-shortener/core"
	logger "url-shortener/core/middleware"
	"url-shortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(logger.Logger())
	r.Use(gin.Recovery())

	if err := core.InitDB(); err != nil {
		logger.Fatal("Error initializing database:", err)
		return
	}
	defer func() {
		if err := core.Close(); err != nil {
			logger.Fatal("Error closing database:", err)
		}
	}()

	// Initialize routes
	initGroup := r.Group("/")
	routes.InitRoutes(initGroup)

	// Replace with your proxy IP or CIDR range
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// routes.RegisterURLRoutes(r)

	if err := r.Run(":8080"); err != nil {
		logger.Fatal("Server failed to start:", err)
	}

	logger.Info("Server running on http://localhost:8080")
}
