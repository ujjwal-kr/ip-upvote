package main

import (
	"fmt"
	"strconv"

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
	ID          int
}

// Upvoters strict
type Upvoters struct {
	gorm.Model
	IP     string
	ItemID int
}

// DB export
var DB *gorm.DB

func main() {
	db := Init()

	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Upvoters{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		var items []Item
		db.Where(&Item{}).Find(&items)
		c.HTML(200, "index.tmpl", gin.H{"items": items})
	})

	r.POST("/", postItem)
	r.GET("/upvote/:id", upvote)
	r.Run(":8080")
	// port := os.Getenv("PORT")
	// r.Run(":" + port)
}

func upvote(c *gin.Context) {
	db := GetDB()
	ip := c.GetHeader("x-forwarded-for")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}
	item := &Item{}
	db.Where(&Item{ID: id}).First(&item)
	if len(item.Name) <= 0 {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}
	upvoter := &Upvoters{}
	db.Where(&Upvoters{IP: ip, ItemID: item.ID}).First(&upvoter)
	if len(upvoter.IP) > 0 {
		c.HTML(200, "err.tmpl", gin.H{"message": "already upvoted, for god's sake dont try to spam"})
		return
	}
	db.Create(&Upvoters{IP: ip, ItemID: item.ID})
	item.Upvotes = item.Upvotes + 1
	db.Save(&item)
	c.Redirect(302, "/")
	return
}

func postItem(c *gin.Context) {
	db := GetDB()
	item := &Item{}
	c.Bind(&item)

	if len(item.Name) <= 0 || len(item.Description) <= 0 {
		c.JSON(422, gin.H{"error": "validation"})
		return
	}

	db.Create(&item)
	c.Redirect(302, "/")
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
