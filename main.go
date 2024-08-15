package main

import (
	"books/internal/core/domain"
	bookrepository "books/internal/core/repositories/book"
	"books/internal/core/services"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase() *gorm.DB {

	dbURL := "postgres://postgres:password@localhost/postgres?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	db.AutoMigrate(&domain.Book{})

	return db
}

func main() {

	db := connectDatabase()

	// routes.InitRoutes().Run(":3000")

	repo := bookrepository.NewBookRepository(db)
	book := services.NewBookService(repo)

	err := book.CreateBook(domain.Book{
		Title:  "Title book",
		Author: "Gabriel",
		Desc:   "An awesome book",
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
