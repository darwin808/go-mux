package db

import (
	"log"

	"github.com/darwin808/greatest-todo/cmd/greatest-todo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:12345678@localhost:5432/greatest-db"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Accounts{})

	return db
}
