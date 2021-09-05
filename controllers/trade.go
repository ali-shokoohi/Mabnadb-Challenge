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
