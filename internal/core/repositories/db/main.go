package db

import (
	"books/internal/core/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	dbURL := "postgres://postgres:password@localhost/postgres?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	db.AutoMigrate(&domain.Book{})

	return db
}
