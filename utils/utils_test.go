package utils

import (
	// "fmt"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/gocolly/colly"
)

var c *colly.Collector

func TestMain(m *testing.M) {
	setUp()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func setUp() {
	fmt.Println("set up")

	c = colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)

}

func TestGetPrice(t *testing.T) {

	output := GetPrice(c)

	c.OnError(func(r *colly.Response, err error) { // Set error handler
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	if output != 4.949999809265137 {
		t.Errorf("Price %v incorrect...", output)
	}
}

func TestGetDiscount(t *testing.T) {
	output := GetDiscountPercentage(c)
	if output != 50 {
		t.Errorf("Discount precentage %f is incorrect", output)
	}
}
