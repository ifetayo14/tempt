package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func StartDB() {
	_, err := gorm.Open(postgres.Open(os.Getenv("PG_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: ", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
