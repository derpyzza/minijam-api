package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Jam struct {
	ID    []int64  `json:"jamId"`
	Name  []string `json:"jamName"`
	URL   []string `json:"jamLink"`
	Image []string `json:"jamImage"`
}

// return array of jams of requested size
func getJams(w http.ResponseWriter, r *http.Request) {
	// jam := Jam{
	// 	ID:    115,
	// 	Name:  "Noir",
	// 	URL:   "https://itch.io/jam/mini-jam-115-noir",
	// 	Image: "https://minijamofficial.com/images/MJ115.gif"}
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(jam)
}

// fetch and return the latest mini jam entry.
func getLatest(w http.ResponseWriter, r *http.Request) {
	const url string = "https://minijamofficial.com/api/fetchMiniJams?n=0"

	// returns response from url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	jam := Jam{}

	w.Header().Set("Content-Type", "application/json")
	jsonerr := json.NewDecoder(resp.Body).Decode(&jam)
	if jsonerr != nil {
		fmt.Println("Json err:")
		log.Fatalln(jsonerr)
	}
	json.NewEncoder(w).Encode(jam)
}

func main() {
	fmt.Println("hello there;")

	router := mux.NewRouter()

	// get list of jams
	router.HandleFunc("/api/jam", getJams).Methods("GET")
	// get latest jam
	router.HandleFunc("/api/latest", getLatest).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
