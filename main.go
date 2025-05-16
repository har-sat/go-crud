package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/har-sat/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariales()
	initializers.ConnectToDatabase()
}
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"hello": "World",
		})
	})

	r.Run()
}
