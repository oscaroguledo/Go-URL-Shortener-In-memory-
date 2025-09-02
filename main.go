package main

import (
	"log"
	"url-shortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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
