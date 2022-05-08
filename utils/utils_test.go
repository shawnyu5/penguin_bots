package utils

import (
	"log"
	"os"
	"testing"

	"github.com/gocolly/colly"
)

var c *colly.Collector

func TestMain(m *testing.M) {
	// setUp()
	exitVal := m.Run()
	c.Visit("https://www.penguinmagic.com/p/3901")
	os.Exit(exitVal)
}

// beforeEach call before each test case
func beforeEach() {
	c = colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)
}

func handleError() {
	c.OnError(func(r *colly.Response, err error) { // Set error handler
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
}

func TestGetDiscountPercentage(t *testing.T) {
	beforeEach()
	output := GetDiscountPercentage()
	if output != 50 {
		t.Errorf("Discount precentage %f is incorrect", output)
	}
}

func TestGetPrice(t *testing.T) {
	beforeEach()
	output := GetPrice()
	handleError()

	if output != 4.949999809265137 {
		t.Errorf("Price %v incorrect...", output)
	}
}

func TestGetDiscountPrice(t *testing.T) {
	beforeEach()
	output := GetDiscountedPrice(c)
	expected := 5.050000190734863
	handleError()

	if output != expected {
		t.Errorf("discounted price %f is incorrect. Expected %f", output, expected)
	}
}
