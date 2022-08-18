package check_coin_product

import (
	"fmt"
	"testing"

	"github.com/gocolly/colly"
)

var c *colly.Collector

// setup creates and returns an instance of the CoinProductServiceImpl struct
func setup() CoinProductService {
	var p CoinProductService
	p = CoinProductServiceImpl{}
	return p
}

// TestMakeDB tests the makeDB function is able to create a local database
func TestMakeDB(t *testing.T) {
	db := makeDB()
	defer db.Close()
	if db == nil {
		t.Error("makeDB returned nil")
	}
}

// TestIsCoinProduct tests isCoinProduct() is able to detect a coin product
func TestIsCoinProduct(t *testing.T) {
	// valid product
	product := CoinProduct{Title: "Coin fjdsljf"}
	valid := isCoinProduct(&product)
	if !valid {
		t.Errorf("Product should be valid, instead got %s", product.Reason)
	}

	// invalid product
	product = CoinProduct{Title: "jfsdjfsljf"}
	valid = isCoinProduct(&product)
	if valid {
		t.Errorf("Product should be invalid, instead got %s", product.Reason)
	}
}

// TestCheckInvalidProduct tests if Check() is able to detect an invalid product
func TestCheckInvalidProduct(t *testing.T) {
	p := setup()
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/8474"),
	)

	product := CoinProduct{}
	getProductInfo(c, &product, "https://www.penguinmagic.com/p/8474")
	p.Check(&product)
	if product.IsValid {
		fmt.Println(fmt.Sprintf("TestCheckInvalidProduct product.IsValid: %v", product.IsValid)) // __AUTO_GENERATED_PRINT_VAR__
		t.Errorf("Product should be invalid, instead got %s", product.Reason)
	}
}

// TestCheckInvalidProduct tests if it is able to detect a valid product
func TestCheckValidProduct(t *testing.T) {
	p := setup()
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/1797"),
	)

	product := CoinProduct{}
	getProductInfo(c, &product, "https://www.penguinmagic.com/p/1797")
	p.Check(&product)
	if !product.IsValid {
		t.Errorf("Product should be valid, instead got %s", product.Reason)
	}
}
