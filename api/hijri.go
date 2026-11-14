package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HijriRequest struct {
	Date string `json:"date"`
}

type HijriResponse struct {
	Hijri     string `json:"hijri"`
	Gregorian string `json:"gregorian"`
}

func gregorianToHijri(year, month, day int) (int, int, int) {
	y := year
	m := month
	d := day

	if m < 3 {
		y--
		m += 12
	}

	a := y / 100
	b := 2 - a + (a / 4)
	jd := int(365.25*float64(y+4716)) + int(30.6001*float64(m+1)) + d + b - 1524

	l := jd - 1948440 + 10632
	n := (l - 1) / 10631
	l = l - 10631*n + 354
	j := ((10985-l)/5316)*((50*l)/17719) + (l/5670)*((43*l)/15238)
	l = l - ((30-j)/15)*((17719*j)/50) - (j/16)*((15238*j)/43) + 29

	hMonth := (24 * l) / 709
	hDay := l - (709*hMonth)/24
	hYear := 30*n + j - 30

	return hYear, hMonth, hDay
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req HijriRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Date == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	months := []string{
		"",
		"Muharram",
		"Safar",
		"Rabi' al-Awwal",
		"Rabi' al-Thani",
		"Jumada al-Awwal",
		"Jumada al-Thani",
		"Rajab",
		"Sha'ban",
		"Ramadan",
		"Shawwal",
		"Dhu al-Qi'dah",
		"Dhu al-Hijjah",
	}

	hYear, hMonth, hDay := gregorianToHijri(parsedDate.Year(), int(parsedDate.Month()), parsedDate.Day())
	monthName := months[hMonth]

	res := HijriResponse{
		Hijri:     fmt.Sprintf("%d %s %d AH", hDay, monthName, hYear),
		Gregorian: parsedDate.Format("January 2, 2006"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
