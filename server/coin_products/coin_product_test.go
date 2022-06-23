package check_coin_product

import (
	"testing"

	"github.com/gocolly/colly"
)

var c *colly.Collector

// TestIsCoinProduct tests if it is able to detect if the product is a coin product
func TestIsCoinProduct(t *testing.T) {
	// A valid coin product
	coinProduct := CoinProduct{Title: "Coin fjdsljf",
		Description:   "jfldsjf coin",
		OriginalPrice: 0.00}
	shouldBeCoinProduct := isCoinProduct(&coinProduct)
	if !shouldBeCoinProduct {
		t.Errorf("Coin product not detected")
	}

	// create a non coin product
	notCoinProduct := CoinProduct{Title: "jfsdjfsljf",
		Description:   "jfldsfjf jjfdsl",
		OriginalPrice: 0.00}

	invalidCoinProduct := isCoinProduct(&notCoinProduct)

	if invalidCoinProduct == true {
		t.Errorf("false Coin product detected")
	}
}

// TestCheck tests if it is able to check if the product is a coin product
func TestCheck(t *testing.T) {
	product := Check("https://www.penguinmagic.com/p/1806")
	if !product.IsValid {
		t.Errorf("Product should be valid, instead got %s", product.Reason)
	}
}
