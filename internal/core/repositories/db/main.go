package db

import (
	"books/internal/core/domain"

	"log"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm.Open failed: %v", err)
	}

	// Install the OpenTelemetry plugin for GORM
	if err := db.Use(otelgorm.NewPlugin(
		otelgorm.WithDBName("books-db"),
	)); err != nil {
		log.Fatalf("failed to register otelgorm plugin: %v", err)
	}

	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("failed to connect to gorm: %v", err)
	// }

	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
		//logger.Log.Fatal("Error connecting to database: ", err.Error())
	}

	if err := db.AutoMigrate(&domain.Book{}); err != nil {
		log.Fatal("Error auto migrate: ", err.Error())
		//logger.Log.Fatal("Error auto migrate: ", err.Error())
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal("Error auto migrate: ", err.Error())
		//logger.Log.Fatal("Error auto migrate: ", err.Error())
	}

	return db
}
