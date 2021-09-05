package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shokoohi/mabnadp-challenge/models"
	"gitlab.com/shokoohi/mabnadp-challenge/routers"
)

// TestReturnAllTrades - Unit test for ReturnAllTrades controller
func TestReturnAllTrades(t *testing.T) {
	// Get router object from our router package
	router := routers.GetRouter()
	// Create a http request
	request, err := http.NewRequest("GET", "/trades", nil)
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 200, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))

	// Checks response type via convert response body type from []bytes to the []Trade model type
	var result []models.Trade
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Can't convert response body type to []Trade model type")
}
