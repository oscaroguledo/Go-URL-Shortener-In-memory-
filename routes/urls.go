package routes

import (
	"url-shortener/services"

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
func CreateURLHandler(c *gin.Context) {
	// Implementation for creating a short URL
	var req struct {
		OriginalURL string `json:"original_url" binding:"required,url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	shortCode, err := services.CreateShortURL(req.OriginalURL)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create short URL"})
		return
	}

	c.JSON(201, gin.H{
		"original_url": req.OriginalURL,
		"short_code":   shortCode,
		"short_url":    c.Request.Host + "/" + shortCode,
	})
}

func GetURLHandler(c *gin.Context) {
	// Implementation for retrieving the original URL
	shortCode := c.Param("code")
	originalURL, err := services.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(404, gin.H{"error": "Short URL not found"})
		return
	}

	c.JSON(200, gin.H{
		"original_url": originalURL,
		"short_code":   shortCode,
		"short_url":    c.Request.Host + "/" + shortCode,
	})
}

func DeleteURLHandler(c *gin.Context) {
	// Implementation for deleting a short URL
	shortCode := c.Param("code")
	if err := services.DeleteShortURL(shortCode); err != nil {
		c.JSON(404, gin.H{"error": "Short URL not found"})
		return
	}

	c.JSON(200, gin.H{
		"message":    "Short URL deleted successfully",
		"short_code": shortCode,
	})
}
func ListURLsHandler(c *gin.Context) {
	// Implementation for listing all URLs
	urls, err := services.GetAllURLs()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve URLs"})
		return
	}

	c.JSON(200, gin.H{
		"urls": urls,
	})
}

// Routes registers the health endpoint(s)
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/", RootHandler)
	r.GET("/health", HealthHandler)
	r.POST("/urls", CreateURLHandler)
	r.GET("/urls:code", GetURLHandler)
	r.DELETE("/urls:code", DeleteURLHandler)
	r.GET("/urls", ListURLsHandler)
}
