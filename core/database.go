package core

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var settings *SettingStruct

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}
	fmt.Println("Database pool closed")
	return nil
}
