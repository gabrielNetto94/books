package db

import (
	"books/internal/core/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
		//logger.Log.Fatal("Error connecting to database: ", err.Error())
	}

	if err := db.AutoMigrate(&domain.Book{}); err != nil {
		log.Fatal("Error auto migrate: ", err.Error())
		//logger.Log.Fatal("Error auto migrate: ", err.Error())
	}

	return db
}
