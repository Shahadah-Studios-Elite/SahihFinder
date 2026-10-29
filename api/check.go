package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RequestData struct {
	Text string `json:"text"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req RequestData
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := strings.ToLower(req.Text)
	status := "Verified Format"
	message := "This text does not trigger automated safety flags. Remember to always cross-check chains of narrators on Sunnah.com."

	if strings.Contains(input, "prophet said") && !strings.Contains(input, "sahih") && !strings.Contains(input, "bukhari") && !strings.Contains(input, "muslim") {
		status = "Flagged"
		message = "Warning: This text claims to be a saying of the Prophet but does not reference a verified collection (like Bukhari or Muslim). Please verify this source carefully."
	} else if strings.Contains(input, "forward this") || strings.Contains(input, "share to 10 people") {
		status = "Flagged"
		message = "Warning: Chain-letter text detected (e.g., 'share to 10 people'). These are usually false or constructed internet rumors. Do not spread without strict verification."
	}

	res := ResponseData{
		Status:  status,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

