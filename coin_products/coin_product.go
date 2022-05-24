package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	utils "github.com/shawnyu5/penguin-utils"
)

type Product struct {
	Title              string
	Description        string
	OriginalPrice      float64
	DiscountPrice      float64
	DiscountPercentage float64
	Rating             int64
}

// getProductInfo get the product currently on penguin open box
func getProductInfo(c *colly.Collector, product *Product) {
	utils.GetTitle(c, &product.Title)
	utils.GetDescription(c, &product.Description)
	utils.GetPrice(c, &product.OriginalPrice)
	utils.GetDiscountedPrice(c, &product.DiscountPrice)
	utils.GetDiscountPercentage(c, &product.DiscountPercentage)
	utils.GetStarRating(c, &product.Rating)
}

// isCoinProduct check if the product is a coin product
func isCoinProduct(product *Product) bool {
	// check if product description contains "coin"
	if strings.Contains(strings.ToLower(product.Description), "coin ") ||
		strings.Contains(strings.ToLower(product.Title), "coin ") {
		return true
	}
	return false
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com/openbox", "www.penguinmagic.com/p/3901"),
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 6}) // limit the number of parallel requests

	product := Product{}

	getProductInfo(c, &product)

	c.Visit("https://www.penguinmagic.com/openbox")
	// c.Visit("https://www.penguinmagic.com/p/3901")

	// if product is emppty, then openbox is down right now
	if product == (Product{}) {
		fmt.Println("There are no open box products currently")
		os.Exit(1)
	} else if !isCoinProduct(&product) {
		fmt.Println("Product is not a coin product")
		os.Exit(1)
	}

	parsed, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(parsed))
}
