package main

import (
	"go-article-api/config"
	"go-article-api/controllers"
	"go-article-api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default()) // <-- Add this line
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Post{})

	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.FindPost)
	r.GET("/posts/paginate", controllers.GetLimitedOffsetedPost)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.EditPost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run(":8081")
}
