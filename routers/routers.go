package routers

import (
	"github.com/gin-gonic/gin"
)

// GetRouter - Return a Gin router client
func GetRouter() *gin.Engine {
	// Create new router
	r := gin.New()
	//return router
	return r
}
