package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Seller struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}

type Offer struct {
	Type          string `json:"@type"`
	Seller        Seller `json:"seller"`
	URL           string `json:"url"`
	PriceCurrency string `json:"priceCurrency"`
	Price         string `json:"price"`
	Condition     string `json:"itemCondition"`
	Availability  string `json:"availability"`
	ValidUntil    string `json:"priceValidUntil"`
}

type Listing struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	SKU         string `json:"sku"`
	Offers      Offer  `json:"offers"`
}

type DegussaItem struct {
	Name  string `json:"name"`
	Price int64  `json:"price`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("shop.degussa-goldhandel.de"),
	)

	c.OnHTML("script", func(e *colly.HTMLElement) {
		if e.Attr("type") == "application/ld+json" {
			data := Listing{}
			err := json.Unmarshal([]byte(e.Text), &data)
			if err != nil {
				log.Panic(err)
			}

			fmt.Printf("Information for %s\nPrice: %s %s\n", data.Name, data.Offers.Price, data.Offers.PriceCurrency)
		}

	})

	c.Visit("https://shop.degussa-goldhandel.de/gold/anlagemuenzen/krugerrand/1-oz-krugerrand-goldmunze-sudafrika-2022")
	//c.Visit("https://shop.degussa-goldhandel.de/1-oz-britannia-goldmunze-100-pfund-grossbritannien-2023")
}
