package src

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDb(connString string) *gorm.DB {
	var DB *gorm.DB
	var err error
	DB, err = gorm.Open(sqlite.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB.AutoMigrate(&Contact{})
	return DB
}
