package hospital

import (
	"encoding/json"
	"net/http"
)

type SearchDoctorsRequest struct {
	Department string `json:"department,omitempty"`
	Symptom    string `json:"symptom,omitempty"`
}

func (handler *Handler) SearchDoctors(w http.ResponseWriter, r *http.Request) {
	var req SearchDoctorsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Just ignore parsing errors for demo, or we can handle it
	}

	response := map[string]interface{}{
		"success": true,
		"data": []map[string]interface{}{
			{
				"doctor_id":  "doc_101",
				"name":       "Dr. Anisur Rahman",
				"department": "Gastroenterology",
				"degree":     "MBBS, MD (Gastro)",
				"fee":        "1500 BDT",
			},
			{
				"doctor_id":  "doc_102",
				"name":       "Dr. Kamal Hossain",
				"department": "Gastroenterology",
				"degree":     "MBBS, FCPS",
				"fee":        "1200 BDT",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
