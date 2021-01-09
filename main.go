package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
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
