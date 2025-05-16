package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil {
		log.Fatalf("Connect to Database: %v", err)
	}
}