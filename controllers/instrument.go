package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shokoohi/mabnadp-challenge/databases"
	"gitlab.com/shokoohi/mabnadp-challenge/models"
)

// ReturnAllInstruments - Return all Instrument objects from database
func ReturnAllInstruments(c *gin.Context) {
	// Get database client
	db := databases.GetPostgres()
	// Find all instrument objects
	var instruments []models.Instrument
	db.Preload("Trades").Find(&instruments)
	// Send response to request
	c.JSON(http.StatusOK, instruments)
}

// ReturnSingleInstrument - Return A single instrument objects from database via its ID
func ReturnSingleInstrument(c *gin.Context) {
	// Get ID from params
	id, exists := c.Params.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, "You must send id in params!")
	}
	// Get database client
	db := databases.GetPostgres()
	// Find all instrument objects
	var instrument models.Instrument
	db.Preload("Trades").First(&instrument, id)
	// Send response to request
	c.JSON(http.StatusOK, instrument)
}

// CreateSingleInstrument - Create A single instrument objects to database
func CreateSingleInstrument(c *gin.Context) {
	// Create a instance of our model with request data
	var instrument models.Instrument
	err := c.BindJSON(&instrument)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Send a valid instrument's structure!")
		return
	}
	// Sumbit to the database
	db := databases.GetPostgres()
	db.Create(&instrument)
	// Send response to the request
	c.JSON(http.StatusOK, instrument)
}
