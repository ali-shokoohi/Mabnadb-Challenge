package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shokoohi/mabnadp-challenge/routers"
)

func TestHomePage(t *testing.T) {
	// Get router object from our router package
	router := routers.GetRouter()
	// Create a http request
	request, err := http.NewRequest("GET", "/", nil)
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 200, response.Result().StatusCode, "Response code be: "+fmt.Sprint(200)+
		"! But It's: "+fmt.Sprint(response.Result().StatusCode)+"!!!")

	// Define excepted value to check response with this
	expected := map[string]string{
		"status":  "ok",
		"message": "Welcome!",
	}
	// Convert response body type from []bytes to map[string]string
	var result map[string]string
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Unmarshal json error must be Nil!")
	// Check the response body must be equal to expected value
	assert.Equal(t, expected, result, fmt.Sprintf("Response must be: %#v! But It's: %#v !!!", expected, result))

}
