package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type parts struct {
	gorm.Model
	PartName   string  `json:"partname"`
	PartType   string  `json:"parttype"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("parts.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	db.AutoMigrate(&parts{})

	r := gin.Default() 	// Initialize Gin router

	// Define routes and their handlers
	r.GET("/api/parts", func(c *gin.Context) {
		var parts []parts
		db.Find(&parts)
		c.JSON(200, parts)

	})

	r.POST("/api/parts", func(c *gin.Context) {
		var newPart parts
		var err = c.BindJSON(&newPart)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		result := db.Create(&newPart)
		if result.Error != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(200, newPart)
	})

	// Start the server
	r.Run(":8080")
}
