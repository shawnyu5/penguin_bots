package check_coin_product

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gocolly/colly"
	"github.com/patrickmn/go-cache"
)

var c *colly.Collector

// TestIsCoinProduct tests if it is able to detect if the product is a coin product
func TestIsCoinProduct(t *testing.T) {
	// A valid coin product
	coinProduct := Product{Title: "Coin fjdsljf",
		Description:   "jfldsjf coin",
		OriginalPrice: 0.00}
	shouldBeCoinProduct := isCoinProduct(&coinProduct)
	if !shouldBeCoinProduct {
		t.Errorf("Coin product not detected")
	}

	// create a non coin product
	notCoinProduct := Product{Title: "jfsdjfsljf",
		Description:   "jfldsfjf jjfdsl",
		OriginalPrice: 0.00}

	invalidCoinProduct := isCoinProduct(&notCoinProduct)

	if invalidCoinProduct == true {
		t.Errorf("false Coin product detected")
	}
}

// TestHasProductChanged tests if it is able to detect if the product has changed
func TestHasProductChanged(t *testing.T) {
	// create a new product
	newProduct := Product{Title: "new product"}
	storage = cache.New(cache.NoExpiration, 30*time.Minute)
	// store product in cache
	cacheProduct(newProduct)
	// product should not have changed
	changed := hasProductChanged(&newProduct)
	if changed {
		t.Errorf("Product should not have changed")
	}

	// passing in a different product than product stored in cache
	changed = hasProductChanged(&Product{Title: "another product"})
	if !changed {
		t.Errorf("Product should have changed")
	}
}

// TestCheck tests if it is able to check if the product is a coin product
func TestCheck(t *testing.T) {
	product := Check("https://www.penguinmagic.com/p/1806")
	var productStruct Product
	json.Unmarshal([]byte(product), &productStruct)
	if !productStruct.IsValid {
		t.Errorf("Product should be valid, instead got %s", productStruct.Reason)
	}
}
