package controllers

import (
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
