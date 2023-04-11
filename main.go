package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// parts struct to seed record parts data
type parts struct {
	gorm.Model         // This is a gorm.Model struct that contains ID, CreatedAt, UpdatedAt, DeletedAt fields
	PartName   string  `json:"partname"`
	PartType   string  `json:"parttype"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
}

func main() {
	// Set up database connection
	db, err := gorm.Open(sqlite.Open("parts.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&parts{}) // what does this do? I think it creates the table if it doesn't exist already??? Or no?

	// Initialize Gin router
	r := gin.Default()

	// Define your routes and their handlers :: this seems to be wokring and I was getting json data back
	r.GET("/api/parts", func(c *gin.Context) {
		// Handle GET /parts request
		var parts []parts

		// Get all parts from database
		db.Find(&parts) // find all parts

		// Return all parts as JSON
		c.JSON(200, parts)

	})

	// Is this working?
	r.POST("/api/parts", func(c *gin.Context) {
		// Handle POST /parts request
		var newPart parts

		// Bind the JSON body of the request to the newPart struct
		var err = c.BindJSON(&newPart)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		// Create a new part record in the database
		result := db.Create(&newPart)
		if result.Error != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		// Return the newly created part as JSON
		c.JSON(200, newPart)
	})

	// Start the server
	r.Run(":8080")
}
