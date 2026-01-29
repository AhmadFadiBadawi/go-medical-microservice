package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// The "Patient Data" structure
type Result struct {
	HeartRate int    `json:"heart_rate"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

func main() {
	http.HandleFunc("/check-pulse", func(w http.ResponseWriter, r *http.Request) {
		// 1. Get the heart rate from the URL (e.g., ?bpm=85)
		bpmStr := r.URL.Query().Get("bpm")
		bpm, _ := strconv.Atoi(bpmStr)

		// 2. Simple Medical Logic
		status := "Normal"
		msg := "Patient is stable."
		if bpm > 100 {
			status = "High (Tachycardia)"
			msg = "Alert: Patient requires immediate review."
		} else if bpm < 60 && bpm > 0 {
			status = "Low (Bradycardia)"
			msg = "Patient is resting or athletic."
		}

		// 3. Send the JSON response
		json.NewEncoder(w).Encode(Result{
			HeartRate: bpm,
			Status:    status,
			Message:   msg,
		})
	})

	http.ListenAndServe(":8080", nil)
}