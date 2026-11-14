package handler

import (
	"encoding/json"
	"fmt"
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

type DorarResponse struct {
	Ahadith struct {
		Result string `json:"result"`
	} `json:"ahadith"`
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

	query := strings.TrimSpace(req.Query)
	var matches []HadithMatch

	localDatabase := []struct {
		Keywords []string
		Title    string
		Snippet  string
		URL      string
	}{
		{
			Keywords: []string{"intentions", "intention", "actions", "deeds", "niyyah"},
			Title:    "Sahih al-Bukhari 1 - Actions are by Intentions",
			Snippet:  "I heard Allah's Messenger saying, 'The reward of deeds depends upon the intentions and every person will get the reward according to what he has intended.'",
			URL:      "https://sunnah.com",
		},
		{
			Keywords: []string{"intentions", "actions", "deeds", "emigration", "hijrah"},
			Title:    "Sahih Muslim 1907 - Deeds are Judged by Motives",
			Snippet:  "Actions are (judged) by motives (niyyah), so each man will have what he intended.",
			URL:      "https://sunnah.com",
		},
		{
			Keywords: []string{"pillars", "islam", "five", "shahada", "prayer", "zakat", "ramadan", "hajj"},
			Title:    "Sahih al-Bukhari 8 - The Five Pillars of Islam",
			Snippet:  "Islam is based on (the following) five (principles): To testify that none has the right to be worshipped but Allah and Muhammad is Allah's Messenger, to offer the compulsory prayers, to pay Zakat, to perform Hajj, and to observe fast during Ramadan.",
			URL:      "https://sunnah.com",
		},
		{
			Keywords: []string{"manners", "character", "best", "akhlaq"},
			Title:    "Sahih al-Bukhari 3559 - Best in Manners and Character",
			Snippet:  "The best among you are those who have the best manners and character.",
			URL:      "https://sunnah.com",
		},
		{
			Keywords: []string{"anger", "angry", "strong", "wrestle", "control"},
			Title:    "Sahih al-Bukhari 6114 - Controlling Anger",
			Snippet:  "The strong is not the one who overcomes the people by his strength, but the strong is the one who controls himself while in anger.",
			URL:      "https://sunnah.com",
		},
		{
			Keywords: []string{"speak", "good", "silent", "silence", "tongue", "neighbor", "guest"},
			Title:    "Sahih al-Bukhari 6018 - Speak Good or Remain Silent",
			Snippet:  "Whoever believes in Allah and the Last Day should talk what is good or keep quiet, and whoever believes in Allah and the Last Day should not hurt his neighbor, and whoever believes in Allah and the Last Day should entertain his guest generously.",
			URL:      "https://sunnah.com",
		},
		{
			Keywords: []string{"muslim", "tongue", "hands", "safe", "peace"},
			Title:    "Sahih al-Bukhari 10 - Safety from Tongue and Hands",
			Snippet:  "A Muslim is the one who avoids harming Muslims with his tongue and hands.",
			URL:      "https://sunnah.com0",
		},
		{
			Keywords: []string{"cleanliness", "purity", "faith", "iman", "wudu"},
			Title:    "Sahih Muslim 223 - Cleanliness is Half of Faith",
			Snippet:  "Cleanliness is half of faith and Alhamdulillah (Praise be to Allah) fills the scale.",
			URL:      "https://sunnah.com",
		},
		{
			Keywords: []string{"brother", "love", "loves", "himself", "wish"},
			Title:    "Sahih al-Bukhari 13 - Loving for Your Brother",
			Snippet:  "None of you will have faith till he wishes for his (Muslim) brother what he likes for himself.",
			URL:      "https://sunnah.com3",
		},
		{
			Keywords: []string{"knowledge", "learn", "quran", "teach", "best"},
			Title:    "Sahih al-Bukhari 5027 - Learning and Teaching the Quran",
			Snippet:  "The best among you (Muslims) are those who learn the Qur'an and teach it.",
			URL:      "https://sunnah.com",
		},
		{
			Keywords: []string{"smile", "charity", "sadaqah", "smiling"},
			Title:    "Jami at-Tirmidhi 1956 - Smiling is Charity",
			Snippet:  "Your smiling in the face of your brother is charity, commanding good and forbidding evil is charity.",
			URL:      "https://sunnah.com",
		},
		{
			Keywords: []string{"parents", "mother", "father", "paradise", "jannah", "companionship"},
			Title:    "Sahih al-Bukhari 5971 - Good Treatment of Parents",
			Snippet:  "A man asked, 'O Allah's Messenger! Who is more entitled to be treated with the best companionship by me?' The Prophet said, 'Your mother.' (Three times, then 'Your father.')",
			URL:      "https://sunnah.com",
		},
	}

	lowerQuery := strings.ToLower(query)
	queryWords := strings.Fields(lowerQuery)

	for _, entry := range localDatabase {
		matched := false
		for _, kw := range entry.Keywords {
			if strings.Contains(lowerQuery, kw) {
				matched = true
				break
			}
			for _, qw := range queryWords {
				cleanQw := strings.Trim(qw, ".,!?\"'")
				if len(cleanQw) > 2 && strings.Contains(kw, cleanQw) {
					matched = true
					break
				}
			}
			if matched {
				break
			}
		}

		if matched {
			matches = append(matches, HadithMatch{
				Title:   entry.Title,
				Snippet: entry.Snippet,
				URL:     entry.URL,
			})
			if len(matches) >= 2 {
				break
			}
		}
	}

	if len(matches) < 2 {
		dorarURL := "https://dorar.net" + url.QueryEscape(query)
		client := &http.Client{Timeout: 4 * time.Second}
		resp, err := client.Get(dorarURL)
		if err == nil && resp.StatusCode == 200 {
			var dResp DorarResponse
			if err := json.NewDecoder(resp.Body).Decode(&dResp); err == nil && dResp.Ahadith.Result != "" {
				tagRe := regexp.MustCompile(`<[^>]*>`)
				parts := strings.Split(dResp.Ahadith.Result, "--------------")
				for _, part := range parts {
					clean := strings.TrimSpace(tagRe.ReplaceAllString(part, " "))
					clean = strings.Join(strings.Fields(clean), " ")
					if len(clean) > 20 {
						if len(clean) > 240 {
							clean = clean[:240] + "..."
						}
						matches = append(matches, HadithMatch{
							Title:   "Dorar Hadith Verification Index",
							Snippet: clean,
							URL:     fmt.Sprintf("https://sunnah.com", url.QueryEscape(query)),
						})
						if len(matches) >= 2 {
							break
						}
					}
				}
			}
			resp.Body.Close()
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SearchResponse{Results: matches})
}

