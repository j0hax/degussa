package degussa

import (
	"log"

	"github.com/gocolly/colly"
)

func init() {
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
}
