package core

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/savioxavier/termlink"
	"github.com/charmbracelet/log"
)

type Jam struct {
	ID    []int64  `json:"jamId"`
	Name  []string `json:"jamName"`
	URL   []string `json:"jamLink"`
	Image []string `json:"jamImage"`
}

func Serve() {

	router := mux.NewRouter()

	// {{ API ENDPOINTS --

	// list of jams.
	router.HandleFunc("/api/jams", GetJams).Methods("GET")
	// latest jam.
	router.HandleFunc("/api/latest", GetLatest).Methods("GET", "OPTIONS")
	// winners.
	router.HandleFunc("/api/winners", GetWinners).Methods("GET", "OPTIONS")

	// -- API ENDPOINTS }}

	ScrapeItch()

	log.Info("App served at")
	fmt.Println(termlink.ColorLink("\t~> Port :8080", "http://localhost:8080/api/latest", "italic blue"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
