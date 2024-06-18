package main

import (
	"books/models/book"
	"books/routes"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase() *gorm.DB {
	dbURL := "postgres://postgres:password@db:5432?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	db.AutoMigrate(&book.Book{})

	return db
}

func main() {

	routes.InitRoutes().Run(":3000")

	// db := connectDatabase()

	// router.GET("/books", func(ctx *gin.Context) {

	// 	var books []models.Book
	// 	res := db.Find(&books)

	// 	if res.Error != nil {
	// 		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
	// 			ctx.JSON(404, gin.H{"message": res.Error.Error()})
	// 			return
	// 		}
	// 		ctx.JSON(500, gin.H{"error": res.Error.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(200, books)
	// })

	// router.GET("/books/:id", func(ctx *gin.Context) {

	// 	id := ctx.Param("id")
	// 	var user models.Book
	// 	res := db.First(&user, id)

	// 	if res.Error != nil {
	// 		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
	// 			ctx.JSON(404, gin.H{"message": res.Error.Error()})
	// 			return

	// 		}
	// 		ctx.JSON(500, gin.H{"error": res.Error.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(200, user)
	// })

	// router.POST("/book", func(ctx *gin.Context) {

	// 	var req models.Book
	// 	if err := ctx.BindJSON(&req); err != nil {
	// 		ctx.JSON(400, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	res := db.Create(&models.Book{
	// 		Title:  req.Title,
	// 		Author: req.Author,
	// 		Desc:   req.Desc,
	// 	})
	// 	if res.Error != nil {
	// 		ctx.JSON(400, gin.H{
	// 			"error": res.Error.Error(),
	// 		})
	// 		return
	// 	}

	// 	ctx.JSON(200, res.RowsAffected)
	// })

}
