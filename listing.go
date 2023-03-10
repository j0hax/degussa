// Package degussa implements a basic API to fetch Degussa Goldhandel prices.
package degussa

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Material allows us to differentiate between product materials
type Material int

// Degussa allows users to sort by these metals
const (
	Gold Material = iota
	Silver
	Platinum
	Palladium
	Diverse
)

// An Item contains general information of Degussa's items for sale
type Item struct {
	ItemNo    string `json:"itemNo"`
	ImageURL  string `json:"imageURL"`
	Name      string `json:"name"`
	Material  Material
	BuyPrice  int64 `json:"buyPrice"`
	SellPrice int64 `json:"sellPrice"`
}

// extractNumbers strips all non-numeral characters in a string
// and returns the natural number as an int64
func extractNumbers(s string) int64 {
	re := regexp.MustCompile(`\d+`)

	numbers := re.FindAllString(s, -1)

	// Some items have no buy and/or price, as such we just return 0
	if len(numbers) == 0 {
		return 0
	}

	digits := strings.Join(numbers, "")

	nat, err := strconv.Atoi(digits)
	if err != nil {
		log.Panic(err)
	}

	return int64(nat)
}

// getMaterial determines the material type of a Degussa Product
func getMaterial(name string) Material {
	l := strings.ToLower(name)

	if strings.Contains(l, "gold") {
		return Gold
	}

	if strings.Contains(l, "silber") {
		return Silver
	}

	if strings.Contains(l, "platin") {
		return Platinum
	}

	if strings.Contains(l, "palladium") {
		return Platinum
	}

	return Diverse
}

// parseRow parses a scraped Degussa price table row (tr) and returns item information
func parseRow(el *colly.HTMLElement) (*Item, error) {
	if el.Name != "tr" {
		return nil, errors.New("passed element is not of type <td>, but " + el.Name)
	}

	itemNo := el.ChildText("td:nth-child(1)")
	img := el.ChildAttr("img", "src")
	name := el.ChildText("td:nth-child(3)")
	buy := extractNumbers(el.ChildText("td:nth-child(4)"))
	sell := extractNumbers(el.ChildText("td:nth-child(6)"))

	return &Item{
		ItemNo:    itemNo,
		ImageURL:  img,
		Name:      name,
		Material:  getMaterial(name),
		BuyPrice:  buy,
		SellPrice: sell,
	}, nil
}

func (l Item) String() string {
	b := float64(l.BuyPrice) / 100
	s := float64(l.SellPrice) / 100
	return fmt.Sprintf("%s: %s (%.2f/%.2f)", l.ItemNo, l.Name, b, s)
}
