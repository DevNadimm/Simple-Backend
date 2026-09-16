package hospital

import (
	"encoding/json"
	"net/http"
)

type CheckAvailabilityRequest struct {
	DoctorID string `json:"doctor_id"`
	Date     string `json:"date"`
}

func (handler *Handler) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	var req CheckAvailabilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Ignore error for demo purposes
	}

	date := "2026-09-17"
	if req.Date != "" {
		date = req.Date
	}

	response := map[string]interface{}{
		"success":         true,
		"doctor_name":     "Dr. Anisur Rahman",
		"date":            date,
		"available_slots": []string{"10:00 AM", "12:30 PM", "04:00 PM"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type BookAppointmentRequest struct {
	DoctorID     string `json:"doctor_id"`
	Date         string `json:"date"`
	TimeSlot     string `json:"time_slot"`
	PatientName  string `json:"patient_name"`
	PatientPhone string `json:"patient_phone"`
}

func (handler *Handler) BookAppointment(w http.ResponseWriter, r *http.Request) {
	var req BookAppointmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Ignore error for demo purposes
	}

	date := "2026-09-17"
	if req.Date != "" {
		date = req.Date
	}
	timeSlot := "10:00 AM"
	if req.TimeSlot != "" {
		timeSlot = req.TimeSlot
	}

	response := map[string]interface{}{
		"success":    true,
		"booking_id": "APT-998877",
		"message":    "Appointment successfully booked.",
		"details": map[string]interface{}{
			"doctor":  "Dr. Anisur Rahman",
			"date":    date,
			"time":    timeSlot,
			"room_no": "204",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
