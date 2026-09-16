package hospital

import (
	"encoding/json"
	"net/http"
)

type CheckTreatmentCostRequest struct {
	TreatmentName string `json:"treatment_name"`
}

func (handler *Handler) CheckTreatmentCost(w http.ResponseWriter, r *http.Request) {
	var req CheckTreatmentCostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Ignore error for demo
	}

	response := map[string]interface{}{
		"success":      true,
		"service_name": "MRI of Brain",
		"cost_bdt":     "8000",
		"remarks":      "Please come empty stomach if contrast is required.",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
