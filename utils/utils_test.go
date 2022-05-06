package utils

import (
	"fmt"
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
	fmt.Println("before each")
	c = colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)
}

func TestGetDiscount(t *testing.T) {
	beforeEach()
	output := GetDiscountPercentage(c)
	if output != 50 {
		t.Errorf("Discount precentage %f is incorrect", output)
	}
}

func TestGetPrice(t *testing.T) {
	beforeEach()
	output := GetPrice(c)
	c.OnError(func(r *colly.Response, err error) { // Set error handler
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	if output != 4.949999809265137 {
		t.Errorf("Price %v incorrect...", output)
	}
}
