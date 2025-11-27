package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"os"
)

// Read the CSV file
func ReadCsvFile(filename string) ([][]string, error) {

	// Open the CSV file
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("error opening CSV file")
	}
	defer csvFile.Close()

	// Parse the CSV file
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("error reading CSV records")
	}

	return records[1:], nil
}

func convertToJSON(records [][]string) ([]byte, error) {
	// Assuming the first row contains the headers
	headers := records[0]

	var jsonData []map[string]string

	// Iterate through the records and create a map for each row
	for _, row := range records[1:] {
		recordMap := make(map[string]string)
		for i, value := range row {
			recordMap[headers[i]] = value
		}
		jsonData = append(jsonData, recordMap)
	}

	// Convert the map to JSON
	return json.MarshalIndent(jsonData, "", "  ")
}
