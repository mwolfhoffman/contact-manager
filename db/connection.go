package db

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mwolfhoffman/contact-manager/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./db/dev.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB.AutoMigrate(&models.Contact{})
}
