package core

import "net/http"

type WinnerData struct {
	Name     string  `json:"name"`
	Wins     float64 `json:"wins"`
	JamScore float64 `json:"jamScore"`
}

// get and return jam winner data.
func GetWinners(w http.ResponseWriter, r *http.Request) {
	/*
		fetch latest jam - 1
		query database for the latest jam the server has data on
		if the latest jam - 1 == latest jam in database, continue;
		if the latest jam - 1 != latest jam in database, recursively fetch data from itch.io till the latest jam and latest jam in database match.
		return winners table;
	*/
}
