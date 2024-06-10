package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

/*
Below are powershell commands and one can use cURL as well for same
NOTE : Constraints types are used per given test sheet

GET Request for ListExoplanets:
Invoke-WebRequest -Uri "http://localhost:8080/exoplanets" -Method GET

GET Request for GetExoplanetByID:
Invoke-WebRequest -Uri "http://localhost:8080/exoplanets/{ID}" -Method GET

POST Request for AddExoplanet:
Invoke-WebRequest -Uri "http://localhost:8080/exoplanets" -Method POST -Body (@{ name = "New Exoplanet"; description = "Description of New Exoplanet"; distance = 100; radius = 2; type = "Terrestrial" } | ConvertTo-Json) -ContentType "application/json"

PUT Request for UpdateExoplanet:
Invoke-WebRequest -Uri "http://localhost:8080/exoplanets/{ID}" -Method PUT -Body (@{ name = "Updated Exoplanet"; description = "Description of Updated Exoplanet"; distance = 200; radius = 3; type = "Terrestrial" } | ConvertTo-Json) -ContentType "application/json"

DELETE Request for DeleteExoplanet:
Invoke-WebRequest -Uri "http://localhost:8080/exoplanets/{ID}" -Method DELETE

GET Request for FuelEstimation:
Invoke-WebRequest -Uri "http://localhost:8080/fuel?id={ID}&crew_capacity={capacity_int}" -Method GET
*/

// Exoplanet struct representing an exoplanet
type Exoplanet struct {
	PlanetID   int     `json:"id"`
	PlanetName string  `json:"name"`
	PlanetDesc string  `json:"description"`
	Distance   int     `json:"distance"`
	Radius     float64 `json:"radius"`
	Mass       float64 `json:"mass,omitempty"`
	Type       string  `json:"type"`
}

var (
	exoplanets         = make(map[int]*Exoplanet)
	exoplanetsByRadius = make(map[float64]*Exoplanet)
	exoplanetsByMass   = make(map[float64]*Exoplanet)
)

func main() {
	router := httprouter.New()

	router.POST("/exoplanets", AddExoplanet)
	router.GET("/exoplanets", ListExoplanets)
	router.GET("/exoplanets/:id", GetExoplanetByID)
	router.PUT("/exoplanets/:id", UpdateExoplanet)
	router.DELETE("/exoplanets/:id", DeleteExoplanet)
	router.GET("/fuel", FuelEstimation)

	log.Fatal(http.ListenAndServe(":8080", router))
}
func AddExoplanet(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var newPlanet Exoplanet
	err := json.NewDecoder(req.Body).Decode(&newPlanet)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if planet name already exists
	if planetExists(newPlanet.PlanetName) {
		http.Error(writer, "planet with the same name already exists", http.StatusConflict)
		return
	}

	// Validate exoplanet
	if err := ValidateExoplanet(newPlanet); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the new exoplanet
	newPlanet.PlanetID = len(exoplanets) + 1
	exoplanets[newPlanet.PlanetID] = &newPlanet
	exoplanetsByRadius[newPlanet.Radius] = &newPlanet
	exoplanetsByMass[newPlanet.Mass] = &newPlanet

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(newPlanet)
	fmt.Println(&exoplanets)
}

// Check if planet with the same name already exists
func planetExists(name string) bool {
	for _, planet := range exoplanets {
		if planet.PlanetName == name {
			return true
		}
	}
	return false
}

// ListExoplanets lists all exoplanets
func ListExoplanets(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var exoplanetList []*Exoplanet
	for _, planet := range exoplanets {
		exoplanetList = append(exoplanetList, planet)
		fmt.Println(planet)
	}
	json.NewEncoder(writer).Encode(exoplanetList)
}

// GetExoplanetByID retrieves an exoplanet by ID
func GetExoplanetByID(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := parseID(params)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	planet, ok := exoplanets[id]
	if !ok {
		http.NotFound(writer, req)
		return
	}
	json.NewEncoder(writer).Encode(planet)
}

// UpdateExoplanet updates an existing exoplanet
func UpdateExoplanet(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := parseID(params)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedPlanet Exoplanet
	err = json.NewDecoder(req.Body).Decode(&updatedPlanet)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if planet exists
	planet, ok := exoplanets[id]
	if !ok {
		http.NotFound(writer, req)
		return
	}

	// Validate updated planet
	if err := ValidateExoplanet(updatedPlanet); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Update planet data
	planet.PlanetName = updatedPlanet.PlanetName
	planet.PlanetDesc = updatedPlanet.PlanetDesc
	planet.Distance = updatedPlanet.Distance
	planet.Radius = updatedPlanet.Radius
	planet.Mass = updatedPlanet.Mass
	planet.Type = updatedPlanet.Type

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(planet)
}

// DeleteExoplanet deletes an exoplanet by ID
func DeleteExoplanet(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := parseID(params)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if planet exists
	if _, ok := exoplanets[id]; !ok {
		http.NotFound(writer, req)
		return
	}

	// Delete planet from maps
	delete(exoplanetsByRadius, exoplanets[id].Radius)
	delete(exoplanetsByMass, exoplanets[id].Mass)
	delete(exoplanets, id)
	fmt.Printf("Deleted %s", exoplanets[id].PlanetName)
	writer.WriteHeader(http.StatusOK)
}

// FuelEstimation estimates fuel required for a journey to an exoplanet
func FuelEstimation(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	params := req.URL.Query()
	idStr := params.Get("id")
	crewCapacityStr := params.Get("crew_capacity")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id parameter", http.StatusBadRequest)
		return
	}
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		http.Error(writer, "invalid crew_capacity parameter", http.StatusBadRequest)
		return
	}

	// Check if planet exists
	planet, ok := exoplanets[id]
	if !ok {
		http.NotFound(writer, req)
		return
	}

	// Logic to calculate gravity for each type
	var gravity float64
	if planet.Type == "GasGiant" {
		gravity = 0.5 / (planet.Radius * planet.Radius)
	} else {
		gravity = planet.Mass / (planet.Radius * planet.Radius)
	}

	fuel := float64(planet.Distance) / (gravity * gravity) * float64(crewCapacity)
	fmt.Printf("Fuel Estimation %f", fuel)
	json.NewEncoder(writer).Encode(map[string]float64{"fuel_cost": fuel})
}

// Validates exoplanet data
func ValidateExoplanet(planet Exoplanet) error {
	if planet.PlanetName == "" || planet.PlanetDesc == "" {
		return fmt.Errorf("name and description cannot be empty")
	}
	if planet.Distance < 10 || planet.Distance > 1000 {
		return fmt.Errorf("distance must be between 10 and 1000 light years")
	}
	if planet.Radius < 0.1 || planet.Radius > 10 {
		return fmt.Errorf("radius must be between 0.1 and 10 Earth-radius units")
	}
	if strings.EqualFold(planet.Type, "Terrestrial") && planet.Mass < 0.1 || planet.Mass > 10 {
		return fmt.Errorf("mass for Terrestial Planets must be provided")
	}
	if !strings.EqualFold(planet.Type, "GasGiant") && !strings.EqualFold(planet.Type, "Terrestrial") {
		return fmt.Errorf("planet type %s is not supported yet. Supported planet types are Terrestrial and GasGiant", planet.Type)
	}
	return nil
}

// parseID parses the ID parameter from request parameters
func parseID(params httprouter.Params) (int, error) {
	idStr := params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid id parameter: %s", err)
	}
	return id, nil
}
