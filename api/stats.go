package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Year string `json:"year"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	currentYear := time.Now().Format("2006")

	res := Response{
		Year: currentYear,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
