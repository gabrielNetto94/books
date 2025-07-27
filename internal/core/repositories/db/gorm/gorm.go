package gorm

import (
	"books/internal/core/domain"
	"books/internal/core/repositories/db"
	"context"
	"fmt"
	"log"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseRepositoryImpl struct {
	db *gorm.DB
}

func connectDatabase(dsn string) *gorm.DB {

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

func NewDatabaseRepository(dsn string) db.DatabaseRepositoryInterface {
	db := connectDatabase(dsn)
	return &DatabaseRepositoryImpl{db}
}

func (d *DatabaseRepositoryImpl) Create(ctx context.Context, data any) error {

	res := d.db.WithContext(ctx).Create(data)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *DatabaseRepositoryImpl) Find(ctx context.Context, data any) error {

	resp := d.db.WithContext(ctx).Find(data)
	if resp.Error != nil {
		return resp.Error
	}

	if resp.RowsAffected == 0 {
		return fmt.Errorf("data not found")
	}

	return nil
}
