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
