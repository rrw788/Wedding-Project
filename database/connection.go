package database

import (
	"log"
	"weddingproject/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/weddingdb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Migrate schema for models
	if err := DB.AutoMigrate(&models.Client{}, &models.Event{}); err != nil {
		log.Fatal("Failed to migrate database schema: ", err)
	}

	log.Println("Connected to MySQL database and migrated schema")
}
