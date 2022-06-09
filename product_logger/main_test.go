package main

import "testing"

func TestUpdateProduct(t *testing.T) {
	dbProduct := DbProduct{Title: "test", Appearances: 1, Average_discount: 1.0, Average_price: 1.0}
	penguinProduct := PenguinProduct{Title: "test", DiscountPrice: 1.0, DiscountPercentage: 1.0}
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
	penguinProduct = PenguinProduct{Title: "test", DiscountPrice: 1.0, DiscountPercentage: 1.0}

	err := updateProduct(&dbProduct, penguinProduct)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
