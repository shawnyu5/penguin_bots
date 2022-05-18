package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// handleError provided a generic implenation of colly.OnError
func handleError(c *colly.Collector) {
	c.OnError(func(r *colly.Response, err error) { // Set error handler
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
}

// GetPrice get the price of an product
func GetPrice(c *colly.Collector, price *float64) {
	c.OnHTML(".ourprice", func(e *colly.HTMLElement) {
		stringPrice := e.Text[2:]

		value, err := strconv.ParseFloat(stringPrice, 32)
		if err != nil {
			log.Fatalf("Can not convert price %s to interger", e.Text)
		}
		*price = value
	})

	handleError(c)
}

// GetDiscountPercentage returns the discount percentage of the product passed in
func GetDiscountPercentage(c *colly.Collector, discountPercentage *float64) {
	c.OnHTML(".yousave", func(e *colly.HTMLElement) {
		discountPercentageString := strings.TrimSpace(e.Text)

		openBracket := strings.Index(discountPercentageString, "(")
		percentSign := strings.Index(discountPercentageString, "%")
		discountPercentageString = discountPercentageString[openBracket+1 : percentSign]

		value, err := strconv.ParseFloat(discountPercentageString, 32)
		if err != nil {
			log.Fatalf("Can not convert price %s to interger", e.Text)
		}

		*discountPercentage = value
	})

	handleError(c)
}

// GetDiscountedPrice return the discount price of the product
func GetDiscountedPrice(c *colly.Collector, price *float64) {

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

		*price = value
	})

	handleError(c)
}

// GetStarRating return the number of starts this product has
func GetStarRating(c *colly.Collector, rating *int64) {

	c.OnHTML("#review_summary", func(e *colly.HTMLElement) {
		ratingLink := e.ChildAttr("img", "src")
		lastSlash := strings.LastIndex(ratingLink, "/")
		stringRating, err := strconv.ParseInt(ratingLink[lastSlash+1:lastSlash+2], 0, 0)
		if err != nil {
			log.Fatalf("Unable to convert rating %v to int", ratingLink[lastSlash+1:])
		}
		*rating = stringRating
	})
	handleError(c)
}

// GetTitle gets the title of a product
func GetTitle(c *colly.Collector, title *string) {
	c.OnHTML("#product_name", func(e *colly.HTMLElement) {
		*title = e.ChildText("h1")
	})
	handleError(c)
}

// GetDescription get the description of a product
func GetDescription(c *colly.Collector, description *string) {
	c.OnHTML(".product_subsection", func(e *colly.HTMLElement) {
		des := e.ChildText("p")
		if des != "" {
			*description = des
		}
	})
	handleError(c)
}

// addNotInterested add a product to the not interested list. Return true if product was successfully written to file. False other wise
func AddNotInterested(productTitle string) bool {
	// open a csv file
	file, err := os.OpenFile("not_interested.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("Unable to open file: %v", err)
		return false
	}
	defer file.Close()

	// ask for user confirmation before writing to file
	fmt.Printf("Are you sure you want to add %s to the not interested list? (y/n)", productTitle)
	var answer string
	fmt.Scanln(&answer)

	if answer == "y" {
		// write the product title to the file if it is not already in the file
		if _, err := file.WriteString(productTitle + "\n"); err != nil {
			fmt.Printf("Unable to write to file: %v", err)
		}
		return true
	}
	return false
}

// IfInterested check if the product passed in is interesting.
// Return true if product is interesting. False otherwise
func IfInterested(title string) bool {
	// open a csv file
	file, err := os.OpenFile("../not_interested_products.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("Unable to open file: %v", err)
		return false
	}
	defer file.Close()

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// if the product title is in the file, product is not intersting
		if title == scanner.Text() {
			return false
		}
	}
	return true
}
