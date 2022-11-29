package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/savioxavier/termlink"
)

type Jam struct {
	ID    []int64  `json:"jamId"`
	Name  []string `json:"jamName"`
	URL   []string `json:"jamLink"`
	Image []string `json:"jamImage"`
}

func Serve() {

	router := mux.NewRouter()

	// GET list of jams.
	router.HandleFunc("/api/jams", GetJams).Methods("GET")
	// GET latest jam.
	router.HandleFunc("/api/latest", GetLatest).Methods("GET", "OPTIONS")

	ScrapeItch()

	fmt.Println(termlink.ColorLink("App served on Port :8080", "http://localhost:8080/api/latest", "italic blue"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
