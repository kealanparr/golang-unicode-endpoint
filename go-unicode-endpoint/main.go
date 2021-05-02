package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type unicode struct {
	Count int `json:"count"`
}

// Everything after the / we will count the Unicode value of your char's and return it
// in the HTTP json body with a http header
func main() {
	http.HandleFunc("/", goHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createJSON(requestStr string) (int, error) {
	count := 0
	for _, rune := range requestStr {
		if rune == 47 { // Don't want to count / in the url
			continue
		}
		count += int(rune)
	}
	return count, nil
}

func goHandler(w http.ResponseWriter, r *http.Request) {
	count, err := createJSON(r.URL.Path[len("/"):])
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := unicode{Count: count}
	json.NewEncoder(w).Encode(response)
	r.Header.Set("Host", "200")
}
