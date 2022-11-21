package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Jam struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

// return array of jams of requested size
func getJams(w http.ResponseWriter, r *http.Request) {
	jam := Jam{
		ID:    "115",
		Name:  "Noir",
		URL:   "https://itch.io/jam/mini-jam-115-noir",
		Image: "https://minijamofficial.com/images/MJ115.gif"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jam)
}

// fetch and return the latest mini jam entry.
func getLatest(w http.ResponseWriter, r *http.Request) {
	// todo
}

func main() {
	fmt.Println("hello there;")

	// TEMPORARY URL TO FETCH LATEST JAM
	url := "https://minijamofficial.com/api/fetchMiniJams?n=0"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)

	router := mux.NewRouter()

	// get list of jams
	router.HandleFunc("/api/jam", getJams).Methods("GET")
	// get latest jam
	router.HandleFunc("/api/latest", getLatest).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
