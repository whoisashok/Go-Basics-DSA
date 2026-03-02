package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Log struct {
	Service   string    `json:"service"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type Metrics struct {
	TotalLogs int
	ErrorLogs int
}

const (
	workerCount    = 5
	errorThreshold = 3
)

var (
	metricsMap  = make(map[string]*Metrics)
	logsChan    = make(chan Log, 1000)
	metricsChan = make(chan Log, 1000)
)

func LogMonitor(logs chan string) {
	// Start worker pool
	for i := 0; i < workerCount; i++ {
		go worker(i)
	}
	// Metrics processor
	go metricsProcessor()
	// Alert engine
	go alertEngine()

	http.HandleFunc("/logs", logHandler)
	println("Log Monitor is running on http://localhost:8080/logs")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	var log []Log
	json.NewDecoder(r.Body).Decode(&log)
	for _, log := range log {
		logsChan <- log
	}
	w.WriteHeader(http.StatusOK)
}

// Worker pool to process logs concurrently
func worker(id int) {
	for logEntry := range logsChan {
		// Simulate processing
		fmt.Printf("Worker %d processing %s log\n", id, logEntry.Service)
		// Fan-out to metrics processor
		metricsChan <- logEntry
	}
}

func metricsProcessor() {
	for logEntry := range metricsChan {
		m := metricsMap[logEntry.Service]
		m.TotalLogs++
		if logEntry.Level == "ERROR" {
			m.ErrorLogs++
		}
	}
}

func alertEngine() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		for service, m := range metricsMap {
			if m.ErrorLogs > errorThreshold {
				fmt.Printf("🚨 ALERT: %s has %d errors\n", service, m.ErrorLogs)
			}
		}
	}
}
