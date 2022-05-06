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
		discountPercentageString := strings.TrimSpace(e.Text)

		openBracket := strings.Index(discountPercentageString, "(")
		percentSign := strings.Index(discountPercentageString, "%")
		discountPercentageString = discountPercentageString[openBracket+1 : percentSign]

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

// GetDiscountPrice return the discount price of the product
func GetDiscountPrice(c *colly.Collector) float64 {

	var discountPrice float64
	c.OnHTML(".yousave", func(e *colly.HTMLElement) {
		discountPriceString := strings.TrimSpace(e.Text)

		discountPriceString = strings.Replace(discountPriceString, "$", "", 1)
		firstSpace := strings.Index(discountPriceString, " ")
		// get string up to first space
		discountPriceString = discountPriceString[:firstSpace]

		value, err := strconv.ParseFloat(discountPriceString, 32)
		if err != nil {
			log.Fatalf("Can not convert price %s to interger", e.Text)
		}

		discountPrice = value
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://www.penguinmagic.com/p/3901")
	return discountPrice
}
