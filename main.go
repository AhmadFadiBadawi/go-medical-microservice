package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Result defines the standardized structure for our medical report.
// The `json:"..."` tags tell Go how to format the keys in the JSON output.
type Result struct {
	PatientID string `json:"patient_id"`
	HeartRate int    `json:"heart_rate"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

func main() {
	// The "/check-pulse" route acts as our diagnostic endpoint.
	http.HandleFunc("/check-pulse", func(w http.ResponseWriter, r *http.Request) {
		
		// 1. EXTRACTION: Get Patient ID and BPM from the URL parameters.
		// Example URL: /check-pulse?id=P-100&bpm=120
		id := r.URL.Query().Get("id")
		bpmStr := r.URL.Query().Get("bpm")
		
		// Convert the BPM string to an integer so we can run logic on it.
		bpm, _ := strconv.Atoi(bpmStr)

		// 2. CLINICAL LOGIC: Evaluate the heart rate.
		status := "Normal"
		msg := "Patient is stable."

		if bpm > 100 {
			status = "High (Tachycardia)"
			msg = "Alert: Patient requires immediate review."
		} else if bpm < 60 && bpm > 0 {
			status = "Low (Bradycardia)"
			msg = "Patient is resting or athletic."
		} else if bpm <= 0 {
			status = "Error"
			msg = "Invalid heart rate data provided."
		}

		// 3. RESPONSE: Create the Result record and send it back as JSON.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Result{
			PatientID: id,
			HeartRate: bpm,
			Status:    status,
			Message:   msg,
		})
	})

	// Start the server on port 8080.
	http.ListenAndServe(":8080", nil)
}