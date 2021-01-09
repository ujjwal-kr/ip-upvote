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
	Name        string `json:"name"`
	Description string `json:"description"`
	Upvotes     int    `json:"upvotes"`
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
	db := GetDB()
	item := &Item{}
	c.BindJSON(&item)
	if len(item.Name) <= 0 || len(item.Description) <= 0 {
		c.JSON(422, gin.H{"error": "validation"})
	} else {
		db.Create(&item)
		c.JSON(201, item)
	}
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
