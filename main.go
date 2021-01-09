package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Item Model
type Item struct {
	gorm.Model
	Name        string
	Description string
	Upvotes     int
}

func main() {
	db, error := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if error != nil {
		panic("Connection to database can't be established")
	}

	db.AutoMigrate(&Item{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/", postItem)
	r.GET("/upvote", upvote)
	r.Run(":8080")
}

func upvote(c *gin.Context) {
	c.JSON(200, gin.H{"message": "upvoted"})
}

func postItem(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Post Item"})
}
