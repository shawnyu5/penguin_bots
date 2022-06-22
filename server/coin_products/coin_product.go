package check_coin_product

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/patrickmn/go-cache"
	utils "github.com/shawnyu5/penguin-utils"
)

type CoinProduct struct {
	Title              string
	Description        string
	OriginalPrice      float64
	DiscountPrice      float64
	DiscountPercentage float64
	Rating             int64
	IsValid            bool
	Reason             string
}

// var PRODUCT_INFO_FILE string
var storage *cache.Cache

// Check gets the product from the url passed in, and checks if it's a coin product, and if it is an interesting product.
// NOTE this function is not responsible for checking if the product has changed.
// Returns a CoinProduct struct
func Check(url string) CoinProduct {
	homeDir, _ := os.UserHomeDir()
	storage = cache.New(cache.NoExpiration, 30*time.Minute)
	utils.SetFilePath(homeDir + "/python/penguin_bots/not_interested_products.csv")

	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com",
			"www.penguinmagic.com/p/17235", "www.penguinmagic.com/openbox",
			"www.penguinmagic.com/p/3901", url),
	)

	product := CoinProduct{}

	getProductInfo(c, &product, url)

	// if product is empty, then openbox is down right now
	if product.Title == "" {
		product.IsValid = false
		product.Reason = "Openbox is down currently..."
		log.Println(product.Reason)
	} else if !utils.IfInterested(product.Title) {
		product.IsValid = false
		product.Reason = fmt.Sprintf("Product %s is not interested", product.Title)
		log.Println(product.Reason)
		// NOTE: it should be up to the caller to determine if the product has changed from the last call of this function
		// } else if !hasProductChanged(&product) {
		// product.IsValid = false
		// product.Reason = fmt.Sprintf("Product *%s* has not changed", product.Title)
		// log.Println(product.Reason)
	} else if !isCoinProduct(&product) {
		product.IsValid = false
		product.Reason = fmt.Sprintf("Product *%s* is not a coin product", product.Title)
		log.Println(product.Reason)
	} else {
		product.IsValid = true
		product.Reason = "Product is a coin product"
	}

	// parse into json
	// parsedJson, err := json.MarshalIndent(product, "", "  ")
	// if err != nil {
	// panic(err)
	// }
	return product
}

// getProductInfo get the product currently on penguin open box
func getProductInfo(c *colly.Collector, product *CoinProduct, url string) {
	utils.GetTitle(c, &product.Title)
	utils.GetDescription(c, &product.Description)
	utils.GetPrice(c, &product.OriginalPrice)
	utils.GetDiscountedPrice(c, &product.DiscountPrice)
	utils.GetDiscountPercentage(c, &product.DiscountPercentage)
	utils.GetStarRating(c, &product.Rating)
	c.Visit(url)
}

// hasProductChanged checks if the product has changed compared to product in cache.
// Return true if it has changed. False otherwise
func hasProductChanged(product *CoinProduct) bool {
	// read from file
	fileProduct, found := storage.Get("product_title")
	// if no product in cache, product has changed
	if !found {
		return true
	}

	// if current product is the same as product in file, then exit
	if product.Title == fileProduct {
		// log.Println(fmt.Sprintf("Product %s has not changed", product.Title))
		return false
	}
	return true
}

// isCoinProduct check if the product is a coin product.
// Return true if it is a coin product. False other wise
func isCoinProduct(product *CoinProduct) bool {
	// check if product description contains "coin"
	if strings.Contains(strings.ToLower(product.Description), "coin ") ||
		strings.Contains(strings.ToLower(product.Title), "coin ") ||
		strings.Contains(strings.ToLower(product.Description), "coins ") {
		log.Println("Coin product found")
		return true
	}
	return false
}
