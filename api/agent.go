package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type AgentRequest struct {
	Message string `json:"message"`
}

type AgentResponse struct {
	Reply string `json:"reply"`
}

type PollinationsMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type PollinationsPayload struct {
	Model    string                `json:"model"`
	Messages []PollinationsMessage `json:"messages"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req AgentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || strings.TrimSpace(req.Message) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userMsg := strings.TrimSpace(req.Message)

	payload := PollinationsPayload{
		Model: "openai",
		Messages: []PollinationsMessage{
			{
				Role:    "system",
				Content: "You are SahihAgent, a respectful, helpful Sunni Muslim AI assistant for SahihFinder. Keep answers clear, concise, and grounded in mainstream Sunni sources like Quran, Sahih al-Bukhari, and Sahih Muslim. Always remind users to verify exact Hadith wording on Sunnah.com or with a trusted local scholar.",
			},
			{
				Role:    "user",
				Content: userMsg,
			},
		},
	}

	jsonBytes, err := json.Marshal(payload)
	reply := ""

	if err == nil {
		client := &http.Client{Timeout: 8 * time.Second}
		httpReq, reqErr := http.NewRequest("POST", "https://pollinations.ai", bytes.NewBuffer(jsonBytes))
		if reqErr == nil {
			httpReq.Header.Set("Content-Type", "application/json")
			resp, respErr := client.Do(httpReq)
			if respErr == nil && resp.StatusCode == 200 {
				bodyBytes, readErr := io.ReadAll(resp.Body)
				resp.Body.Close()
				if readErr == nil {
					text := strings.TrimSpace(string(bodyBytes))
					if text != "" && !strings.HasPrefix(text, "<") {
						reply = text
					}
				}
			}
		}
	}

	if reply == "" {
		lower := strings.ToLower(userMsg)
		if strings.Contains(lower, "salam") || strings.Contains(lower, "hello") || strings.Contains(lower, "hi") {
			reply = "Wa Alaykum Assalam! How can I help you verify Islamic information or search trusted Sunni references today?"
		} else if strings.Contains(lower, "hadith") || strings.Contains(lower, "bukhari") || strings.Contains(lower, "muslim") {
			reply = "When verifying any Hadith, always check its grading and exact narration on Sunnah.com (such as Sahih al-Bukhari or Sahih Muslim) rather than relying on unverified social media posts."
		} else {
			reply = "I am here to assist you with mainstream Sunni Islamic knowledge and fact-checking. Please consult Sunnah.com, Quran.com, or a qualified local scholar to confirm detailed rulings."
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AgentResponse{Reply: reply})
}
