package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetPassengers(w http.ResponseWriter, r *http.Request) {
	var passengers []TitanicJSON
	db.Find(&passengers)
	jsonEncoded, _ := json.Marshal(&passengers)

	// Respond with the JSON data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonEncoded)

}

func GetPassengerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var passengers []TitanicJSON
	result := db.First(&passengers, vars["passengerId"])
	jsonEncoded, _ := json.Marshal(result)
	// Respond with the JSON data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsonEncoded)
}

func CSVPassengerHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ReadCsvFile("titanic.csv")
	// Convert CSV to JSON
	jsonData, err := convertToJSON(records)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Respond with the JSON data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func HistogramHandler(w http.ResponseWriter, r *http.Request) {

	records, err := ReadCsvFile("titanic.csv")

	// Extract Fare data from CSV
	var fares []FareData
	for _, record := range records[1:] {
		fare, err := strconv.ParseFloat(record[9], 64) // Assuming the Fare column is at index 9
		if err != nil {
			fmt.Println("Error parsing Fare:", err)
			continue
		}
		fares = append(fares, FareData{Fare: fare})
	}

	// Sort fares
	sort.Slice(fares, func(i, j int) bool {
		return fares[i].Fare < fares[j].Fare
	})

	// Calculate percentiles
	percentiles := []float64{25, 50, 75, 90, 95, 99}
	histogram := make(map[float64]int)
	totalFares := len(fares)

	for _, percentile := range percentiles {
		index := int(percentile / 100 * float64(totalFares))
		histogram[percentile] = index
	}

	// Marshal the histogram to JSON
	histogramJSON, err := json.Marshal(histogram)
	if err != nil {
		http.Error(w, "Error converting histogram to JSON", http.StatusInternalServerError)
		return
	}

	// Respond with the JSON data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(histogramJSON)
}
