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
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)

	product := Product{}

	getProductInfo(c, &product)

	c.Visit("https://www.penguinmagic.com/p/3901")

	if !isCoinProduct(&product) {
		fmt.Println("Product is not a coin product")
		os.Exit(1)
	}

	parsed, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(parsed))
}
