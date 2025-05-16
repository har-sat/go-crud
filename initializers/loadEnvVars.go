package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariales() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env")
	}
}