package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func FetchLatest() Jam {
	jam := Jam{}
	const url string = "https://minijamofficial.com/api/fetchMiniJams?n=0"

	// returns response from url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	jsonerr := json.NewDecoder(resp.Body).Decode(&jam)
	if jsonerr != nil {
		fmt.Println("Json err:")
		log.Fatalln(jsonerr)
	}

	defer resp.Body.Close()

	return jam
}

// return array of jams of requested index
func GetJams(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	jam := Jam{}

	index := r.URL.Query().Get("n")
	url := "https://minijamofficial.com/api/fetchMiniJams?n=" + index

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	jsonerr := json.NewDecoder(resp.Body).Decode(&jam)
	if jsonerr != nil {
		fmt.Println("Json err:")
		log.Fatalln(jsonerr)
	}
	json.NewEncoder(w).Encode(jam)
}

// fetch and return the latest mini jam entry.
func GetLatest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	jam := FetchLatest()
	json.NewEncoder(w).Encode(jam)
}
