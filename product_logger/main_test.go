package main

import (
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func TestUpdateProduct(t *testing.T) {
	// create a db product
	dbProduct := DbProduct{Title: "test", Appearances: 1, Average_discount: 1.0, Average_price: 1.0}
	// create penguinProduct
	penguinProduct := PenguinProduct{Title: "test", Discount_price: 1.0, Discount_percentage: 1.0}
	updateProduct(&dbProduct, penguinProduct)
	if dbProduct.Appearances != 2 {
		t.Errorf("Expected Appearances to be 2, got %d", dbProduct.Appearances)
	}

	if dbProduct.Average_discount != 1.0 {
		t.Errorf("Expected Average_discount to be 1.0, got %f", dbProduct.Average_discount)
	}

	if dbProduct.Average_price != 1.0 {
		t.Errorf("Expected Average_price to be 1.0, got %f", dbProduct.Average_price)
	}

	// If title does not match
	dbProduct = DbProduct{Title: "apple", Appearances: 1, Average_discount: 1.0, Average_price: 1.0}
	penguinProduct = PenguinProduct{Title: "test", Discount_price: 1.0, Discount_percentage: 1.0}

	err := updateProduct(&dbProduct, penguinProduct)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestHasProductChanged(t *testing.T) {
	c = cache.New(cache.NoExpiration, 10*time.Minute)
	product := PenguinProduct{Title: "test", Discount_price: 1.0, Discount_percentage: 1.0}
	cacheProduct(product)
	changed := hasProductChanged(product)
	if changed {
		t.Errorf("Expected changed to be false, got true. Product should not have changed")
	}

	// pass in a different product than cached previously
	changed = hasProductChanged(PenguinProduct{Title: "apple", Discount_price: 2.0, Discount_percentage: 1.0})
	if !changed {
		t.Errorf("Expected changed to be true, got false. Product should have changed")
	}

	// empty cache should assume product has not changed
	c = cache.New(cache.NoExpiration, 10*time.Minute)
	changed = hasProductChanged(product)
	if changed {
		t.Errorf("empty cache: expected changed to be false, got true. Product should not have changed")
	}
}
