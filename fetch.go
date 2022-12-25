package degussa

import (
	"log"

	"github.com/gocolly/colly"
)

var c = colly.NewCollector(
		colly.AllowedDomains("www.degussa-goldhandel.de"),
	)

// FilterTable fetches the price table and returns a list of products
// which return true for f.
func FilterTable(f func(Item) bool) []Item {
	var items []Item

	c.OnHTML("tbody", func(t *colly.HTMLElement) {
		t.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			l, err := parseRow(el)
			if err != nil {
				log.Panic(err)
			}

			if f(*l) {
			items = append(items, *l)
			}
		})
	})

	c.Visit("https://www.degussa-goldhandel.de/preise/preisliste/")

	return items
}

// All returns all listed precious metal bars and coins
func All() []Item {
	return FilterTable(func(i Item) bool {
		return true
	})
}
