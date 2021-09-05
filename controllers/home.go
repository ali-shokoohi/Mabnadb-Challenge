package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePage - Controller for website' home page
func HomePage(c *gin.Context) {
	result := map[string]string{
		"status":  "ok",
		"message": "Welcome!",
	}
	c.JSON(http.StatusOK, result)
}
