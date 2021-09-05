package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shokoohi/mabnadp-challenge/databases"
	"gitlab.com/shokoohi/mabnadp-challenge/models"
)

// ReturnAllTrades - Return all of trades from the database
func ReturnAllTrades(c *gin.Context) {
	// Get database client
	db := databases.GetPostgres()
	// Find all trade objects
	var trades []models.Trade
	db.Find(&trades)
	// Send response to request
	c.JSON(http.StatusOK, trades)
}

// ReturnSingleTrade - Return A single trade objects from database via its ID
func ReturnSingleTrade(c *gin.Context) {
	// Get ID from params
	id, exists := c.Params.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, "You must send id in params!")
	}
	// Get database client
	db := databases.GetPostgres()
	// Find all instrument objects
	var trade models.Trade
	db.First(&trade, id)
	// Send response to request
	c.JSON(http.StatusOK, trade)
}

// CreateSingleTrade - Create A single trade objects to database
func CreateSingleTrade(c *gin.Context) {
	// Create a instance of our model with request data
	var trade models.Trade
	err := c.BindJSON(&trade)
	if err != nil {
		log.Println("Error: ", err, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, "Send a valid instrument's structure!")
		return
	}
	// Sumbit to the database
	db := databases.GetPostgres()
	db.Create(&trade)
	// Send response to the request
	c.JSON(http.StatusCreated, trade)
}

// DeleteSingleTrade - Delete a single trade from database via its ID
func DeleteSingleTrade(c *gin.Context) {
	// Get ID from params
	id, exists := c.Params.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, "You must send id in params!")
	}
	// Get database client
	db := databases.GetPostgres()
	// Find all instrument objects
	var trade models.Trade
	db.Preload("Trades").First(&trade, id)
	// Delete that instrument object
	db.Delete(&trade)
	// Send response to the request
	c.JSON(http.StatusOK, trade)
}
