package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shokoohi/mabnadp-challenge/models"
	"gitlab.com/shokoohi/mabnadp-challenge/routers"
)

// TestReturnAllInstruments - Unit test for ReturnAllInstruments controller
func TestReturnAllInstruments(t *testing.T) {
	// Get router object from our router package
	router := routers.GetRouter()
	// Create a http request
	request, err := http.NewRequest("GET", "/instruments", nil)
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 200, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))

	// Checks response type via convert response body type from []bytes to the []Instrument model type
	var result []models.Instrument
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Can't convert response body type to []Instrument model type")
}

// TestReturnSingleInstrument - Unit test for ReturnSingleInstrument controller
func TestReturnSingleInstrument(t *testing.T) {
	// Get router object from our router package
	router := routers.GetRouter()
	// Create a http request
	request, err := http.NewRequest("GET", "/instruments/1", nil)
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 200, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))

	// Checks response type via convert response body type from []bytes to the []Instrument model type
	var result models.Instrument
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Can't convert response body type to Instrument model type")
}

// TestCreateSingleInstrument - Unit test for CreateSingleInstrument and DeleteSingleInstrument controllers
func TestCreateDeleteSingleInstrument(t *testing.T) {
	// Create request body
	instrument := map[string]string{"name": "Test"}
	requestBody, err := json.Marshal(instrument)
	assert.Nil(t, err, "Create request body's error must be Nil!")
	// Get router object from our router package
	router := routers.GetRouter()
	// Create a http request
	request, err := http.NewRequest("POST", "/instruments", bytes.NewBuffer(requestBody))
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 201, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))
	// Checks response type via convert response body type from []bytes to the Instrument model type
	var result models.Instrument
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Can't convert response body type to Instrument model type")
	// Result values must be equal to request's value
	assert.Equal(t, instrument["name"], result.Name, fmt.Sprintf("Response must be: %#v! But It's: %#v !!!", instrument, result))
	// Create a http request
	request, err = http.NewRequest("DELETE", fmt.Sprintf("/instruments/%d", result.ID), nil)
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response = httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 200, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))
	// Checks response type via convert response body type from []bytes to the Instrument model type
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Can't convert response body type to Instrument model type")
	// Result values must be equal to created object's value
	assert.Equal(t, instrument["name"], result.Name, fmt.Sprintf("Response must be: %#v! But It's: %#v !!!", instrument, result))
}
