package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type SearchRequest struct {
	Query string `json:"query"`
}

type HadithMatch struct {
	Book   string `json:"book"`
	Number int    `json:"number"`
	Text   string `json:"text"`
	URL    string `json:"url"`
	Score  int    `json:"-"`
}

type SearchResponse struct {
	Results []HadithMatch `json:"results"`
}

type HadithAPIResponse struct {
	Hadiths []struct {
		HadithNumber int    `json:"hadithnumber"`
		Text         string `json:"text"`
	} `json:"hadiths"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || strings.TrimSpace(req.Query) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	words := strings.Fields(strings.ToLower(req.Query))
	var keywords []string
	for _, word := range words {
		clean := strings.Trim(word, ".,!?\"'")
		if len(clean) > 2 {
			keywords = append(keywords, clean)
		}
	}

	collections := []struct {
		Name   string
		Slug   string
		APIURL string
	}{
		{"Sahih al-Bukhari", "bukhari", "https://jsdelivr.net"},
		{"Sahih Muslim", "muslim", "https://jsdelivr.net"},
	}

	var allMatches []HadithMatch
	client := &http.Client{Timeout: 5 * time.Second}

	for _, col := range collections {
		if len(allMatches) >= 10 {
			break
		}
		resp, err := client.Get(col.APIURL)
		if err != nil {
			continue
		}
		var data HadithAPIResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err == nil {
			for _, h := range data.Hadiths {
				if h.Text == "" {
					continue
				}
				lowerText := strings.ToLower(h.Text)
				score := 0
				for _, kw := range keywords {
					if strings.Contains(lowerText, kw) {
						score++
					}
				}
				if score > 0 {
					if strings.Contains(lowerText, strings.ToLower(req.Query)) {
						score += 5
					}
					allMatches = append(allMatches, HadithMatch{
						Book:   col.Name,
						Number: h.HadithNumber,
						Text:   h.Text,
						URL:    fmt.Sprintf("https://sunnah.com", col.Slug, h.HadithNumber),
						Score:  score,
					})
				}
			}
		}
		resp.Body.Close()
	}

	for i := 0; i < len(allMatches); i++ {
		for j := i + 1; j < len(allMatches); j++ {
			if allMatches[j].Score > allMatches[i].Score {
				allMatches[i], allMatches[j] = allMatches[j], allMatches[i]
			}
		}
	}

	limit := 2
	if len(allMatches) < limit {
		limit = len(allMatches)
	}

	topResults := allMatches[:limit]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SearchResponse{Results: topResults})
}

