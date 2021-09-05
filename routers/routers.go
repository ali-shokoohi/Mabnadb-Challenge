package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shokoohi/mabnadp-challenge/controllers"
)

// GetRouter - Return a Gin router client
func GetRouter() *gin.Engine {
	// Create new router
	r := gin.New()
	// Route for "/" endpoint
	r.Any("/", controllers.HomePage)
	// Routes for /instruments... endpoints
	r.GET("/instruments", controllers.ReturnAllInstruments)
	r.GET("/instruments/:id", controllers.ReturnSingleInstrument)
	r.POST("/instruments", controllers.CreateSingleInstrument)
	r.DELETE("/instruments/:id", controllers.DeleteSingleInstrument)
	// Routes for /trades... endpoints
	r.GET("/trades", controllers.ReturnAllTrades)
	r.GET("/trades/:id", controllers.ReturnSingleTrade)
	r.POST("/trades", controllers.CreateSingleTrade)
	r.DELETE("/trades/:id", controllers.DeleteSingleTrade)
	//return router
	return r
}
