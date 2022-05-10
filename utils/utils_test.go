package utils

import (
	"log"
	"testing"

	"github.com/gocolly/colly"
)

var c *colly.Collector

// beforeEach call before each test case
func beforeEach() {
	c = colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)
}

func TestGetDiscountPercentage(t *testing.T) {
	beforeEach()
	var output float64
	GetDiscountPercentage(c, &output)

	c.Visit("https://www.penguinmagic.com/p/3901")

	if output != 50 {
		t.Errorf("Expected %v, got %v", 50, output)
	}
}

func TestGetStarRating(t *testing.T) {
	beforeEach()
	var output int64
	GetStarRating(c, &output)
	if output != 5 {
		t.Fatalf("Expected %d, got %d", 5, output)
	}
}

// func TestGetPrice(t *testing.T) {
// beforeEach()
// var output float64
// GetPrice(c, &output)
// handleError()

// if output != 4.949999809265137 {
// t.Errorf("Price %v incorrect...", output)
// }
// }

// func TestGetDiscountPrice(t *testing.T) {
// beforeEach()
// var output float64
// GetDiscountedPrice(c, &output)
// expected := 5.050000190734863
// handleError()

// if output != expected {
// t.Errorf("Got %f. Expected %f", output, expected)
// }
// }
