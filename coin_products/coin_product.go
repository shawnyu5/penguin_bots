package check_coin_product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	utils "github.com/shawnyu5/penguin-utils"
)

type Product struct {
	Title              string
	Description        string
	OriginalPrice      float64
	DiscountPrice      float64
	DiscountPercentage float64
	Rating             int64
	IsValid            bool
	Reason             string
}

var PRODUCT_INFO_FILE string
var NOT_INTERESTED_FILE string

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
	f, err := ioutil.ReadFile(PRODUCT_INFO_FILE)
	if err != nil {
		log.Fatal(err)
	}
	fileProduct := string(f)

	// if current product is the same as product in file, then exit
	if product.Title == fileProduct {
		fmt.Println(fmt.Sprintf("Product %s has not changed", product.Title))
		return false
	}

	// check if product description contains "coin"
	if strings.Contains(strings.ToLower(product.Description), "coin ") ||
		strings.Contains(strings.ToLower(product.Title), "coin ") {
		return true
	}
	return false
}

// saveProductToFile save the product to product_info.txt
func saveProductToFile(product *Product) {
	// write to file
	err := os.WriteFile(PRODUCT_INFO_FILE, []byte(product.Title), 0644)
	if err != nil {
		panic(err)
	}
}

// Check gets the product from the url passed in, and checks if it's a coin product.
// Returns a json object with the product info
func Check(url string) string {
	godotenv.Load()
	PRODUCT_INFO_FILE = os.Getenv("PRODUCT_INFO_FILE")
	NOT_INTERESTED_FILE = os.Getenv("NOT_INTERESTED_FILE")
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com",
			"www.penguinmagic.com/p/17235", "www.penguinmagic.com/openbox",
			"www.penguinmagic.com/p/3901", url),
	)

	product := Product{}

	getProductInfo(c, &product, url)

	saveProductToFile(&product)

	// check if we are interested in the product
	if !utils.IfInterested(product.Title) {
		fmt.Println(fmt.Sprintf("Product %s is not interested", product.Title))
		product.IsValid = false
		product.Reason = fmt.Sprintf("Product %s is not interested", product.Title)
	}

	// if product is empty, then openbox is down right now
	if product == (Product{}) {
		fmt.Println("There are no open box products currently")
		product.IsValid = false
		product.Reason = "There are no open box products currently"
		// os.Exit(1)
	} else if !isCoinProduct(&product) {
		fmt.Println(fmt.Sprintf("Product %s is not a coin product", product.Title))
		product.IsValid = false
		product.Reason = fmt.Sprintf("Product *%s* is not a coin product", product.Title)
	}

	// parse into json
	parsedJson, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(parsedJson)
}
