package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestGetExoplanetByIDHandler(t *testing.T) {
	// Prepare a request with GET method and an ID parameter
	req, err := http.NewRequest("GET", "/exoplanets/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a new HTTP router and attach the handler
	router := httprouter.New()
	router.GET("/exoplanets/:id", GetExoplanetByID)

	// Serve the request to the handler
	router.ServeHTTP(rr, req)

	// Check if the status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response body is not empty
	if rr.Body.Len() == 0 {
		t.Error("handler returned empty body, expected an exoplanet")
	}
}

func TestUpdateExoplanetHandler(t *testing.T) {
	// Prepare a request body for updating an exoplanet
	requestBody := []byte(`{"id":1,"name":"Updated Exoplanet","description":"Description of Updated Exoplanet","distance":200,"radius":3,"mass":4.5,"type":"GasGiant"}`)

	// Create a request with PUT method, request body, and an ID parameter
	req, err := http.NewRequest("PUT", "/exoplanets/1", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a new HTTP router and attach the handler
	router := httprouter.New()
	router.PUT("/exoplanets/:id", UpdateExoplanet)

	// Serve the request to the handler
	router.ServeHTTP(rr, req)

	// Check if the status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response body is not empty
	if rr.Body.Len() == 0 {
		t.Error("handler returned empty body, expected updated exoplanet details")
	}
}

func TestDeleteExoplanetHandler(t *testing.T) {
	// Prepare a request with DELETE method and an ID parameter
	req, err := http.NewRequest("DELETE", "/exoplanets/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a new HTTP router and attach the handler
	router := httprouter.New()
	router.DELETE("/exoplanets/:id", DeleteExoplanet)

	// Serve the request to the handler
	router.ServeHTTP(rr, req)

	// Check if the status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response body is empty (delete handler usually doesn't return content)
	if rr.Body.Len() != 0 {
		t.Error("handler returned non-empty body, expected empty body")
	}
}

func TestFuelEstimationHandler(t *testing.T) {
	// Prepare a request with GET method and query parameters
	req, err := http.NewRequest("GET", "/fuel?id=1&crew_capacity=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a new HTTP router and attach the handler
	router := httprouter.New()
	router.GET("/fuel", FuelEstimation)

	// Serve the request to the handler
	router.ServeHTTP(rr, req)

	// Check if the status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response body is not empty
	if rr.Body.Len() == 0 {
		t.Error("handler returned empty body, expected fuel estimation")
	}
}
