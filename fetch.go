package degussa

import (
	"log"

	"github.com/gocolly/colly"
)

// All returns all listed precious metal bars and coins
func All() []Item {
	var items []Item

	c := colly.NewCollector(
		colly.AllowedDomains("www.degussa-goldhandel.de"),
	)

	c.OnHTML("tbody", func(t *colly.HTMLElement) {
		t.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			l, err := parseRow(el)
			if err != nil {
				log.Panic(err)
			}

			items = append(items, *l)
		})
	})

	c.Visit("https://www.degussa-goldhandel.de/preise/preisliste/")

	return items
}
