package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

// TestReturnSingleTrade - Unit test for ReturnSingleTrade controller
func TestReturnSingleTrade(t *testing.T) {
	// Get router object from our router package
	router := routers.GetRouter()
	// Create a http request
	request, err := http.NewRequest("GET", "/trades/1", nil)
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 200, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))

	// Checks response type via convert response body type from []bytes to the Trade model type
	var result models.Trade
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Can't convert response body type to Trade model type")
}

// TestCreateSingleTrade - Unit test for CreateSingleTrade controller
func TestCreateSingleTrade(t *testing.T) {
	// Create request body for create a instrument
	instrument := map[string]string{"name": "Test"}
	instrumentBody, err := json.Marshal(instrument)
	assert.Nil(t, err, "Create request body's error must be Nil!")
	// Get router object from our router package
	router := routers.GetRouter()
	// Create a http request
	request, err := http.NewRequest("POST", "/instruments", bytes.NewBuffer(instrumentBody))
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 201, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))
	// Checks response type via convert response body type from []bytes to the Instrument model type
	var resultInstrument models.Instrument
	err = json.Unmarshal(response.Body.Bytes(), &resultInstrument)
	assert.Nil(t, err, "Can't convert response body type to Instrument model type")
	// Result values must be equal to request's value
	assert.Equal(t, instrument["name"], resultInstrument.Name, fmt.Sprintf("Response must be: %#v! But It's: %#v !!!", instrument, resultInstrument))
	// Create request body for create a trade
	trade := models.Trade{
		Open: 1, Close: 1, High: 1, Low: 1, InstrumentID: resultInstrument.ID, DateEn: time.Now(),
	}
	tradeBody, err := json.Marshal(trade)
	assert.Nil(t, err, "Create request body's error must be Nil!")
	// Get router object from our router package
	router = routers.GetRouter()
	// Create a http request
	request, err = http.NewRequest("POST", "/trades", bytes.NewBuffer(tradeBody))
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response = httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 201, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))
	// Checks response type via convert response body type from []bytes to the Trade model type
	var result models.Trade
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Can't convert response body type to Trade model type")
	// Result values must be equal to request's value
	assert.Equal(t, trade.Open, result.Open, fmt.Sprintf("Response must be: %#v! But It's: %#v !!!", trade, result))
	assert.Equal(t, trade.Close, result.Close, fmt.Sprintf("Response must be: %#v! But It's: %#v !!!", trade, result))
	assert.Equal(t, trade.High, result.High, fmt.Sprintf("Response must be: %#v! But It's: %#v !!!", trade, result))
	assert.Equal(t, trade.Low, result.Low, fmt.Sprintf("Response must be: %#v! But It's: %#v !!!", trade, result))
}
