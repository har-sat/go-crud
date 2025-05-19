package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/har-sat/go-crud/initializers"
	"github.com/har-sat/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"post": post})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")
	result := initializers.DB.Find(&post, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"post": post})
}


func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body string
		Title string
	}
	c.Bind(&body)

	var post models.Post
	result := initializers.DB.Find(&post, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
	}

	initializers.DB.Model(&post).Updates(models.Post{Body: body.Body, Title: body.Title})

	c.IndentedJSON(http.StatusOK, gin.H{"post": "updated"})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	
	result := initializers.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
	}
	
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("post with id = %v deleted successfully", id)})
}