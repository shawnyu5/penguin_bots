package utils

import (
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
)

const webAddress = "https://www.penguinmagic.com/p/3901"

// GetPrice get the price of an product
func GetPrice() float64 {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)
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

	c.Visit(webAddress)
	return price
}

// GetDiscountPercentage returns the discount percentage of the product passed in
func GetDiscountPercentage() float64 {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)
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

	c.Visit(webAddress)
	return discountPercentage
}

// GetDiscountPrice return the discount price of the product
func GetDiscountPrice() float64 {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)

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

	c.Visit(webAddress)
	return discountPrice
}

// GetStarRating return the number of starts this product has
func GetStarRating() int64 {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)
	var rating int64
	c.OnHTML("#review_summary", func(e *colly.HTMLElement) {
		ratingLink := e.ChildAttr("img", "src")
		lastSlash := strings.LastIndex(ratingLink, "/")
		stringRating, err := strconv.ParseInt(ratingLink[lastSlash+1:lastSlash+2], 0, 0)
		if err != nil {
			log.Fatalf("Unable to convert rating %v to int", ratingLink[lastSlash+1:])
		}
		rating = stringRating
	})

	c.Visit(webAddress)
	return rating
}
