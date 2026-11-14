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
	Title   string `json:"title"`
	Snippet string `json:"snippet"`
	URL     string `json:"url"`
	Score   int    `json:"-"`
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

	rawQuery := strings.TrimSpace(req.Query)
	lowerQuery := strings.ToLower(rawQuery)
	words := strings.Fields(lowerQuery)

	stopWords := map[string]bool{
		"the": true, "and": true, "for": true, "that": true, "this": true,
		"with": true, "you": true, "are": true, "from": true, "have": true,
		"hadith": true, "about": true, "what": true, "said": true, "prophet": true,
	}

	var keywords []string
	for _, word := range words {
		clean := strings.Trim(word, ".,!?\"'():;")
		if len(clean) > 2 && !stopWords[clean] {
			keywords = append(keywords, clean)
		}
	}

	if len(keywords) == 0 {
		for _, word := range words {
			clean := strings.Trim(word, ".,!?\"'():;")
			if len(clean) > 1 {
				keywords = append(keywords, clean)
			}
		}
	}

	liveCollections := []struct {
		Label  string
		Slug   string
		APIURL string
	}{
		{"40 Hadith Nawawi", "nawawi40", "https://jsdelivr.net"},
		{"40 Hadith Qudsi", "qudsi40", "https://jsdelivr.net"},
	}

	var matches []HadithMatch
	client := &http.Client{Timeout: 4 * time.Second}

	for _, col := range liveCollections {
		resp, err := client.Get(col.APIURL)
		if err != nil || resp.StatusCode != 200 {
			continue
		}

		var apiData HadithAPIResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiData); err == nil {
			for _, h := range apiData.Hadiths {
				if h.Text == "" {
					continue
				}
				lowerText := strings.ToLower(h.Text)
				score := 0

				if strings.Contains(lowerText, lowerQuery) {
					score += 10
				}

				for _, kw := range keywords {
					if strings.Contains(lowerText, kw) {
						score += 2
					}
				}

				if score > 0 {
					snippet := h.Text
					if len(snippet) > 360 {
						snippet = snippet[:360] + "..."
					}
					matches = append(matches, HadithMatch{
						Title:   fmt.Sprintf("%s #%d", col.Label, h.HadithNumber),
						Snippet: snippet,
						URL:     fmt.Sprintf("https://sunnah.com", col.Slug, h.HadithNumber),
						Score:   score,
					})
				}
			}
		}
		resp.Body.Close()
	}

	for i := 0; i < len(matches); i++ {
		for j := i + 1; j < len(matches); j++ {
			if matches[j].Score > matches[i].Score {
				matches[i], matches[j] = matches[j], matches[i]
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SearchResponse{Results: matches})
}
