package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Jam struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	URL    string `json:"url"`
	Image  string `json:"image"`
}

func getJams(w http.ResponseWriter, r *http.Request) {
	jam := Jam{
		Name:   "Noir",
		Number: "115",
		URL:    "https://itch.io/jam/mini-jam-115-noir",
		Image:  "https://minijamofficial.com/images/MJ115.gif"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jam)
}

func main() {
	fmt.Println("hello there;")

	router := mux.NewRouter()

	router.HandleFunc("/jam", getJams).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
