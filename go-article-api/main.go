package main

import (
	"go-article-api/config"
	"go-article-api/controllers"
	"go-article-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Post{})

	r.GET("/posts", controllers.GetPosts)
	r.POST("/posts", controllers.CreatePost)

	r.Run(":8081")
}
