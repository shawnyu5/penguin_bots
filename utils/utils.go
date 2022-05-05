package utils

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// GetPrice get the price of an product
func GetPrice(c *colly.Collector) float64 {
	var price float64
	c.OnHTML(".ourprice", func(e *colly.HTMLElement) {
		stringPrice := e.Text[2:]

		value, err := strconv.ParseFloat(stringPrice, 32)
		if err != nil {
			log.Fatalf("Can not convert price %s to interger", e.Text)
		}
		price = value

	})

	c.OnError(func(r *colly.Response, err error) { // Set error handler
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://www.penguinmagic.com/p/3901")
	return price
}

// GetDiscountPercentage returns the discount percentage of the product passed in
func GetDiscountPercentage(c *colly.Collector) float64 {
	var discountPercentage float64

	c.OnHTML(".yousave", func(e *colly.HTMLElement) {
		discountPercentageString := e.Text

		index := strings.Index(discountPercentageString, "(")
		discountPercentageString = discountPercentageString[index+1:]
		discountPercentageString = discountPercentageString[:len(discountPercentageString)-2]

		value, err := strconv.ParseFloat(discountPercentageString, 32)
		if err != nil {
			log.Fatalf("Can not convert price %s to interger", e.Text)
		}

		discountPercentage = value
	})

	c.OnError(func(r *colly.Response, err error) { // Set error handler
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://www.penguinmagic.com/p/3901")
	return discountPercentage
}
