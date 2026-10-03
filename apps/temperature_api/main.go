package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func randomTemperature() float64 {
	// Random temperature between 15.0 and 30.0 degrees, one decimal place
	value := 15.0 + rand.Float64()*15.0
	return float64(int(value*10)) / 10
}

func writeJSON(w http.ResponseWriter, resp TemperatureResponse) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /temperature?location=<location>
func handleTemperatureByLocation(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		http.Error(w, `{"error":"location query parameter is required"}`, http.StatusBadRequest)
		return
	}

	resp := TemperatureResponse{
		Value:       randomTemperature(),
		Unit:        "celsius",
		Timestamp:   time.Now().UTC(),
		Location:    location,
		Status:      "active",
		SensorType:  "temperature",
		Description: "Temperature reading for " + location,
	}
	writeJSON(w, resp)
}

// GET /temperature/<sensor_id>
func handleTemperatureByID(w http.ResponseWriter, r *http.Request) {
	sensorID := strings.TrimPrefix(r.URL.Path, "/temperature/")
	if sensorID == "" {
		http.Error(w, `{"error":"sensor_id is required"}`, http.StatusBadRequest)
		return
	}

	resp := TemperatureResponse{
		Value:       randomTemperature(),
		Unit:        "celsius",
		Timestamp:   time.Now().UTC(),
		Status:      "active",
		SensorID:    sensorID,
		SensorType:  "temperature",
		Description: "Temperature reading for sensor " + sensorID,
	}
	writeJSON(w, resp)
}

func main() {
	http.HandleFunc("/temperature/", handleTemperatureByID)
	http.HandleFunc("/temperature", handleTemperatureByLocation)

	log.Println("temperature-api listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
