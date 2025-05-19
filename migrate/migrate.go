package main

import (
	"github.com/har-sat/go-crud/initializers"
	"github.com/har-sat/go-crud/models"
)

func init() {
	initializers.LoadEnvVariales()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}