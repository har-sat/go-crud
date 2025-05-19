package main

import (
	"github.com/gin-gonic/gin"
	"github.com/har-sat/go-crud/controllers"
	"github.com/har-sat/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariales()
	initializers.ConnectToDatabase()
}
func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.GetAllPosts)

	r.GET("/posts/:id", controllers.GetPost)
	r.POST("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	
	r.Run()
}
