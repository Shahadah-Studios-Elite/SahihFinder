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
	if r.Method != "POST" {
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
	
	status := "Needs Manual Check"
	message := "This text cannot be automatically verified. Please cross-check this information manually on Sunnah.com or ask a trusted local scholar."

	hasTrustedSource := strings.Contains(input, "bukhari") || 
		strings.Contains(input, "muslim") || 
		strings.Contains(input, "tirmidhi") || 
		strings.Contains(input, "abu dawud") || 
		strings.Contains(input, "nasa'i") || 
		strings.Contains(input, "ibn majah")

	containsAuthorityWord := strings.Contains(input, "said") || 
		strings.Contains(input, "prophet") || 
		strings.Contains(input, "told")

	if strings.Contains(input, "forward this") || 
		strings.Contains(input, "share to") || 
		strings.Contains(input, "10 people") {
		status = "Flagged"
		message = "Warning: Chain-letter formatting detected. This resembles known internet rumors. Do not spread without strict verification."
	} else if containsAuthorityWord && !hasTrustedSource {
		status = "Flagged"
		message = "Warning: An unverified transmission or command statement was detected without referencing a reliable collection (like Bukhari or Muslim). Please check this claim carefully."
	} else if hasTrustedSource {
		status = "Source Referenced"
		message = "A classical Sunni Hadith collection is referenced in the text. Make sure to double-check the specific chapter and volume number to confirm accuracy."
	}

	res := ResponseData{
		Status:  status,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

