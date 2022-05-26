package check_coin_product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
func getProductInfo(c *colly.Collector, product *Product, url string) {
	utils.GetTitle(c, &product.Title)
	utils.GetDescription(c, &product.Description)
	utils.GetPrice(c, &product.OriginalPrice)
	utils.GetDiscountedPrice(c, &product.DiscountPrice)
	utils.GetDiscountPercentage(c, &product.DiscountPercentage)
	utils.GetStarRating(c, &product.Rating)
	c.Visit(url)
}

// isCoinProduct check if the product is a coin product. Return true if it is a coin product. False other wise
func isCoinProduct(product *Product) bool {
	// read from file
	f, err := ioutil.ReadFile("product_info.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileProduct := string(f)

	// if current product is the same as product in file, then exit
	if product.Title == fileProduct {
		fmt.Println(fmt.Sprintf("Product %s has not changed", product.Title))
		os.Exit(1)
	}

	// check if product description contains "coin"
	if strings.Contains(strings.ToLower(product.Description), "coin ") ||
		strings.Contains(strings.ToLower(product.Title), "coin ") {
		return true
	}
	return false
}

// saveProductToFile save the product to file
func saveProductToFile(product *Product) {
	// write to file
	err := os.WriteFile("product_info.txt", []byte(product.Title), 0644)
	if err != nil {
		panic(err)
	}
}

// Check gets the current product from open box, and checks if it's a coin product.
// Returns a json object with the product info
func Check() string {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com",
			"www.penguinmagic.com/p/17235", "www.penguinmagic.com/openbox",
			"www.penguinmagic.com/p/3901"),
	)

	product := Product{}

	getProductInfo(c, &product, "https://www.penguinmagic.com/openbox")

	saveProductToFile(&product)
	// if product is empty, then openbox is down right now
	if product == (Product{}) {
		fmt.Println("There are no open box products currently")
		os.Exit(1)
	} else if !isCoinProduct(&product) {
		fmt.Println(fmt.Sprintf("Product %s is not a coin product", product.Title))
		os.Exit(1)
	}

	// parse into json
	parsedJson, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(parsedJson)
}
