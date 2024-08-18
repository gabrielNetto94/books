package db

import (
	"books/internal/core/domain"
	"books/internal/pkg/env"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	dbURL := env.GetVariable("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	db.AutoMigrate(&domain.Book{})

	return db
}
