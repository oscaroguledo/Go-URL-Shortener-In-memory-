package services

import (
	"url-shortener/core"
	"url-shortener/models"
)

func CreateShortURL(originalURL string) (string, error) {
	shortCode := core.GenerateShortCode(6)
	url := models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
	}
	if err := core.DB.Create(&url).Error; err != nil {
		return "", err
	}
	return shortCode, nil
}

func GetOriginalURL(shortCode string) (string, error) {
	var url models.URL
	if err := core.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return "", err
	}
	return url.OriginalURL, nil
}

func DeleteShortURL(shortCode string) error {
	if err := core.DB.Where("short_code = ?", shortCode).Delete(&models.URL{}).Error; err != nil {
		return err
	}
	return nil
}
func GetAllURLs() ([]models.URL, error) {
	var urls []models.URL
	if err := core.DB.Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}
