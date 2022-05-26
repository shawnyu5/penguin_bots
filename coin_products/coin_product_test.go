package check_coin_product

import (
	"io/ioutil"
	"testing"

	"github.com/gocolly/colly"
)

var c *colly.Collector

func TestIsCoinProduct(t *testing.T) {
	coinProduct := Product{Title: "Coin fjdsljf",
		Description:   "jfldsjf coin",
		OriginalPrice: 0.00}
	shouldBeCoinProduct := isCoinProduct(&coinProduct)
	if !shouldBeCoinProduct {
		t.Errorf("Coin product not detected")
	}

	notCoinProduct := Product{Title: "jfsdjfsljf",
		Description:   "jfldsfjf jjfdsl",
		OriginalPrice: 0.00}

	invalidCoinProduct := isCoinProduct(&notCoinProduct)

	if invalidCoinProduct {
		t.Errorf("false Coin product detected")
	}
}

func TestSaveProductToFile(t *testing.T) {
	handleError := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	// original contents of file
	oldfileProduct, err := ioutil.ReadFile("product_info.txt")
	handleError(err)
	// create new product and write it to file
	newProduct := Product{Title: "new product"}
	saveProductToFile(&newProduct)
	newfileProduct, err := ioutil.ReadFile("product_info.txt")
	handleError(err)

	if "new product" != string(newfileProduct) {
		t.Errorf("Product not saved to file")
	}

	// replace new product with old product
	err = ioutil.WriteFile("product_info.txt", oldfileProduct, 0644)
	handleError(err)
}
