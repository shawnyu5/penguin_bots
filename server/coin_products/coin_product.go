package check_coin_product

import (
	"fmt"
	"strings"

	"server/utils"

	"git.mills.io/prologic/bitcask"
	"github.com/gocolly/colly"
)

type CoinProductService interface {
	// Check gets the product from the url passed in, and checks if it's a coin product, and if it is an interesting product.
	// sets the product.IsValid field to true if it is a coin product. False otherwise. And sets the product.Reason field
	Check(product *CoinProduct)
}

type CoinProductServiceImpl struct{}

type CoinProduct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsValid     bool   `json:"is_valid"`
	Reason      string `json:"reason"`
}

// var PRODUCT_INFO_FILE string
// var storage *cache.Cache

func makeDB() *bitcask.Bitcask {
	b, err := bitcask.Open("./db")
	if err != nil {
		panic(err)
	}
	return b
}

// Check gets the product from the url passed in, and checks if it's a coin product, and if it is an interesting product.
// sets the product.IsValid field to true if it is a coin product. False otherwise. And the product.Reason field
func (CoinProductServiceImpl) Check(product *CoinProduct) {
	// if product is empty, then openbox is down right now
	if product.Title == "" {
		product.IsValid = false
		product.Reason = "Openbox is down currently..."
	} else if !utils.IfInterested(product.Title) {
		product.IsValid = false
		product.Reason = fmt.Sprintf("Product %s is not interested", product.Title)
	} else if !isCoinProduct(product) {
		product.IsValid = false
		product.Reason = fmt.Sprintf("Product *%s* is not a coin product", product.Title)
	} else {
		product.IsValid = true
		product.Reason = "Product is a coin product"
	}
}

// getProductInfo get the product currently on penguin open box. Stores the product info in product struct passed in
func getProductInfo(c *colly.Collector, product *CoinProduct, url string) {
	utils.GetTitle(c, &product.Title)
	utils.GetDescription(c, &product.Description)
	c.Visit(url)
}

// // hasProductChanged checks if the product has changed compared to product in database.
// // Return true if it has changed. False otherwise
// func hasProductChanged(product *CoinProduct, db bitcask.Bitcask) bool {
// // read from file
// // fileProduct, found := storage.Get("product_title")
// // if no product in cache, product has changed
// // if !found {
// // return true
// // }

// // // if current product is the same as product in file, then exit
// // if product.Title == fileProduct {
// // // log.Println(fmt.Sprintf("Product %s has not changed", product.Title))
// // return false
// // }
// return true
// }

// isCoinProduct check if the product is a coin product.
// Return true if it is a coin product. False otherwise
func isCoinProduct(product *CoinProduct) bool {
	// check if product description contains "coin"
	if strings.Contains(strings.ToLower(product.Description), "coin ") ||
		strings.Contains(strings.ToLower(product.Title), "coin ") ||
		strings.Contains(strings.ToLower(product.Description), "coins ") {
		return true
	}
	return false
}
