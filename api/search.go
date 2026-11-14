package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"
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
}

type SearchResponse struct {
	Results []HadithMatch `json:"results"`
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

	searchQuery := "site:sunnah.com " + req.Query
	ddgURL := "https://duckduckgo.com" + url.QueryEscape(searchQuery)

	client := &http.Client{Timeout: 6 * time.Second}
	httpReq, _ := http.NewRequest("GET", ddgURL, nil)
	httpReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := client.Do(httpReq)
	var matches []HadithMatch

	if err == nil && resp.StatusCode == 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		htmlStr := string(bodyBytes)

		linkRe := regexp.MustCompile(`class="result__a" href="([^"]+)"[^>]*>(.*?)</a>`)
		snipRe := regexp.MustCompile(`class="result__snippet"[^>]*>(.*?)</a>`)
		tagRe := regexp.MustCompile(`<[^>]*>`)

		links := linkRe.FindAllStringSubmatch(htmlStr, 5)
		snips := snipRe.FindAllStringSubmatch(htmlStr, 5)

		for i := 0; i < len(links) && len(matches) < 2; i++ {
			rawURL := links[i][1]
			if strings.Contains(rawURL, "uddg=") {
				if parsed, err := url.Parse(rawURL); err == nil {
					if realURL := parsed.Query().Get("uddg"); realURL != "" {
						rawURL = realURL
					}
				}
			}

			if !strings.Contains(rawURL, "sunnah.com") {
				continue
			}

			cleanTitle := strings.TrimSpace(tagRe.ReplaceAllString(links[i][2], ""))
			cleanSnip := ""
			if i < len(snips) {
				cleanSnip = strings.TrimSpace(tagRe.ReplaceAllString(snips[i][1], ""))
			}

			matches = append(matches, HadithMatch{
				Title:   cleanTitle,
				Snippet: cleanSnip,
				URL:     rawURL,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SearchResponse{Results: matches})
}
