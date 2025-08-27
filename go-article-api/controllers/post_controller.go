package controllers

import (
	"go-article-api/config"
	"go-article-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func FindPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func GetLimitedOffsetedPost(c *gin.Context) {
	var posts []models.Post

	// Get query parameters with default values
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	// limit := c.Param("limit")
	// offset := c.Param("offset")

	// Convert to integers
	limitInt, err1 := strconv.Atoi(limit)
	offsetInt, err2 := strconv.Atoi(offset)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit or offset"})
		return
	}

	// Apply pagination
	config.DB.Limit(limitInt).Offset(offsetInt).Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if len(post.Title) < 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title must be at least 20 characters"})
		return
	}
	if len(post.Content) < 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content must be at least 200 characters"})
		return
	}
	if len(post.Category) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category must be at least 3 characters"})
		return
	}
	validStatus := map[string]bool{"Published": true, "Draft": true, "Trashed": true}
	if !validStatus[post.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status must be Published, Draft, or Trashed"})
		return
	}
	//

	config.DB.Create(&post)
	c.JSON(http.StatusCreated, post)
}

func EditPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if len(post.Title) < 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title must be at least 20 characters"})
		return
	}
	if len(post.Content) < 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content must be at least 200 characters"})
		return
	}
	if len(post.Category) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category must be at least 3 characters"})
		return
	}
	validStatus := map[string]bool{"Published": true, "Draft": true, "Trashed": true}
	if !validStatus[post.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status must be Published, Draft, or Trashed"})
		return
	}
	//

	config.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// soft delete menggunakan status "Trashed"
	post.Status = "Trashed"
	config.DB.Save(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post status changed to Trashed"})
}
