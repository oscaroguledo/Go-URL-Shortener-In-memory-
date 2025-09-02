package routes

import (
	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "Url Shortener API",
		"status":  "Running",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"urls":    "/",
			"shorten": "/shorten",
		},
		"docs": map[string]string{
			"Swagger": "/docs",
			"ReDoc":   "/redoc",
		},
	})
}

func HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "healthy",
		"service": "Url Shortener API",
		"version": "1.0.0",
		"redis":   "connected",
	})
}

// Routes registers the health endpoint(s)
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/", RootHandler)
	r.GET("/health", HealthHandler)
}
