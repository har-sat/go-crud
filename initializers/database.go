package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//caps cuz should be available outside this package
var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil {
		log.Fatalf("Connect to Database: %v", err)
	}
}