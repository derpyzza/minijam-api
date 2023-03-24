package main

import (
	"fmt"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"minijamapi.com/packages/core"
	"minijamapi.com/packages/util"
)

func main() {
	fmt.Println("hello there;")
	db, err := sql.Open("sqlite3", "./test.db")
	util.CheckErr(err)

	// do stuff
	core.Serve()

	db.Close()
}

// worry about this later.
// return a specific jam
// func getJam(w http.ResponseWriter, r *http.Request) {

// 	// struct to hold jam data in.
// 	jam := Jam{}
// 	w.Header().Set("Content-Type", "application/json")

// 	// the jam id we're after.
// 	// eg /jam?id=30 = the 30th mini jam.
// 	jamid, err := strconv.ParseFloat(r.URL.Query().Get("id"), 64)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// get total number of jams
// 	totalJams := float64(fetchLatest().ID[0] - 1)

// 	// get page and offset of the jam
// 	placement := (totalJams / 9) - (jamid / 9)
// 	// get page of the jam
// 	page := math.Floor(placement)
// 	// get offset of the jam
// 	offset := math.Floor(placement * 10)

// 	url := "https://minijamofficial.com/api/fetchMiniJams?n=" + strconv.FormatFloat(page)

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()

// 	jsonerr := json.NewDecoder(resp.Body).Decode(&jam)
// 	if jsonerr != nil {
// 		fmt.Println("Json err:")
// 		log.Fatalln(jsonerr)
// 	}
// 	json.NewEncoder(w).Encode(jam)
// }
