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
	// Routes for /instruments endpoint
	r.GET("/instruments", controllers.ReturnAllInstruments)
	//return router
	return r
}