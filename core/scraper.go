package core

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func ScrapeItch() WinnerData {

	c := colly.NewCollector(
	// colly.AllowedDomains("https://itch.io/jam/"),
	)
	// winner name.
	c.OnHTML(
		".first_place .game_summary > h3",
		func(e *colly.HTMLElement) {
			name := e.ChildText("a")
			fmt.Println(name)
		},
	)
	// winner's itch page
	c.OnHTML(
		".first_place .game_summary > h3",
		func(e *colly.HTMLElement) {
			url := e.ChildAttr("a", "href")
			fmt.Println(url)
		},
	)
	// winner's total score
	c.OnHTML(
		".first_place .ranking_results_table > tbody > tr:nth-child(1)",
		func(e *colly.HTMLElement) {
			score := e.ChildText("td:nth-child(3)")
			fmt.Println("score: ", score)
		},
	)
	// winning game's url
	c.OnHTML(
		".first_place .game_summary > h2",
		func(e *colly.HTMLElement) {
			url := e.ChildAttr("a", "href")
			fmt.Println(url)
		},
	)

	cerr := c.Visit("https://itch.io/jam/mini-jam-118-vampires/results")
	if cerr != nil {
		fmt.Println("VISIT ERROR: ")
		log.Fatalln(cerr)
	}

	return WinnerData{}
}
