package main

import (
	"fmt"

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

// DB export
var DB *gorm.DB

func main() {
	db := Init()

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

// Init func
func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: ", err)
	}
	DB = db
	return DB
}

// GetDB func
func GetDB() *gorm.DB {
	return DB
}
