package handler

import (
	"encoding/json"
	"net/http"
	"strings"
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

	database := []HadithMatch{
		{
			Title:   "Sahih al-Bukhari 1 - Actions are by Intentions",
			Snippet: "I heard Allah's Messenger saying, 'The reward of deeds depends upon the intentions and every person will get the reward according to what he has intended.'",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih Muslim 1907 - Deeds are Judged by Motives",
			Snippet: "Actions are judged by motives (niyyah), so each man will have what he intended.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih al-Bukhari 8 - The Five Pillars of Islam",
			Snippet: "Islam is based on five principles: To testify that none has the right to be worshipped but Allah and Muhammad is Allah's Messenger, to offer the compulsory prayers, to pay Zakat, to perform Hajj, and to observe fast during Ramadan.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih al-Bukhari 10 - Safety from Tongue and Hands",
			Snippet: "A Muslim is the one who avoids harming Muslims with his tongue and hands.",
			URL:     "https://sunnah.com0",
		},
		{
			Title:   "Sahih al-Bukhari 13 - Loving for Your Brother",
			Snippet: "None of you will have faith till he wishes for his Muslim brother what he likes for himself.",
			URL:     "https://sunnah.com3",
		},
		{
			Title:   "Sahih al-Bukhari 69 - Make Things Easy",
			Snippet: "Facilitate things to people concerning religious matters, and do not make it hard for them and give them good tidings and do not make them run away.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih al-Bukhari 528 - The Five Daily Prayers",
			Snippet: "If there was a river at the door of anyone of you and he took a bath in it five times a day, would you notice any dirt on him? That is the example of the five prayers with which Allah blots out evil deeds.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih al-Bukhari 1899 - Fasting is a Shield",
			Snippet: "Fasting is a shield with which a servant protects himself from the Fire.",
			URL:     "https://sunnah.com899",
		},
		{
			Title:   "Sahih al-Bukhari 3559 - Best in Manners and Character",
			Snippet: "The best among you are those who have the best manners and character.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih al-Bukhari 5027 - Learning and Teaching the Quran",
			Snippet: "The best among you are those who learn the Qur'an and teach it.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih al-Bukhari 5971 - Good Treatment of Parents",
			Snippet: "A man asked, 'O Allah's Messenger! Who is more entitled to be treated with the best companionship by me?' The Prophet said, 'Your mother.' (Three times, then 'Your father.')",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih al-Bukhari 6018 - Speak Good or Remain Silent",
			Snippet: "Whoever believes in Allah and the Last Day should talk what is good or keep quiet, and whoever believes in Allah and the Last Day should not hurt his neighbor, and whoever believes in Allah and the Last Day should entertain his guest generously.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih al-Bukhari 6094 - Truthfulness Leads to Righteousness",
			Snippet: "Truthfulness leads to righteousness, and righteousness leads to Paradise. And a man keeps on telling the truth until he becomes a truthful person.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih al-Bukhari 6114 - Controlling Anger",
			Snippet: "The strong is not the one who overcomes the people by his strength, but the strong is the one who controls himself while in anger.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih Muslim 223 - Cleanliness is Half of Faith",
			Snippet: "Cleanliness is half of faith and Alhamdulillah fills the scale.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih Muslim 2564 - Allah Looks at Your Hearts",
			Snippet: "Verily Allah does not look to your faces and your wealth, but He looks to your heart and to your deeds.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Sahih Muslim 2699 - Seeking Knowledge",
			Snippet: "Whoever follows a path in the pursuit of knowledge, Allah will make a path to Paradise easy for him.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "Jami at-Tirmidhi 1956 - Smiling is Charity",
			Snippet: "Your smiling in the face of your brother is charity, commanding good and forbidding evil is charity.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "40 Hadith Nawawi 15 - Good Speech and Neighborliness",
			Snippet: "Let him who believes in Allah and the Last Day either speak good or keep silent, and let him who believes in Allah and the Last Day be generous to his neighbor.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "40 Hadith Nawawi 16 - Do Not Become Angry",
			Snippet: "A man said to the Prophet, 'Counsel me.' He said, 'Do not become angry.' The man repeated his request several times, and he said, 'Do not become angry.'",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "40 Hadith Nawawi 34 - Changing Evil",
			Snippet: "Whosoever of you sees an evil, let him change it with his hand; and if he is not able to do so, then with his tongue; and if he is not able to do so, then with his heart.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "40 Hadith Qudsi 1 - Mercy Prevails Over Wrath",
			Snippet: "When Allah decreed the Creation He pledged Himself by writing in His book which is laid down with Him: My mercy prevails over My wrath.",
			URL:     "https://sunnah.com",
		},
		{
			Title:   "40 Hadith Qudsi 17 - Oppression is Forbidden",
			Snippet: "O My servants, I have forbidden oppression for Myself and have made it forbidden amongst you, so do not oppress one another.",
			URL:     "https://sunnah.com7",
		},
		{
			Title:   "40 Hadith Qudsi 25 - Reward for Good Intentions",
			Snippet: "Allah has written down the good deeds and the bad ones. Whoever intended to perform a good deed, yet did not do it, then Allah writes it down with Himself as a complete good deed.",
			URL:     "https://sunnah.com",
		},
	}

	matches := []HadithMatch{}
	for _, entry := range database {
		searchable := strings.ToLower(entry.Title + " " + entry.Snippet)
		score := 0

		if strings.Contains(searchable, lowerQuery) {
			score += 10
		}

		for _, kw := range keywords {
			if strings.Contains(searchable, kw) {
				score += 2
			}
		}

		if score > 0 {
			entry.Score = score
			matches = append(matches, entry)
		}
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
