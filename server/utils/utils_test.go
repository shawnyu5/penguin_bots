package utils_test

import (
	"log"
	"os/exec"
	"testing"

	"github.com/gocolly/colly"
	utils "github.com/shawnyu5/penguin-utils"
)

var c *colly.Collector

// handleError provided a generic implenation of colly.OnError
func handleError(c *colly.Collector) {
	c.OnError(func(r *colly.Response, err error) { // Set error handler
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
}

// beforeEach create a new colly collector
func beforeEach() {
	c = colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)
}

// visit visit web address to start scraping
func visit(c *colly.Collector) {
	c.Visit("https://www.penguinmagic.com/p/3901")

}

func TestGetDiscountPercentage(t *testing.T) {
	beforeEach()
	var output float64
	utils.GetDiscountPercentage(c, &output)

	visit(c)

	if output != 50 {
		t.Errorf("Expected %v, got %v", 50, output)
	}
}

func TestGetStarRating(t *testing.T) {
	beforeEach()
	var output int64
	utils.GetStarRating(c, &output)

	visit(c)

	if output != 4 {
		t.Fatalf("Expected %d, got %d", 4, output)
	}
}

func TestGetPrice(t *testing.T) {
	beforeEach()
	var output float64
	utils.GetPrice(c, &output)
	handleError(c)

	visit(c)

	if output != 4.949999809265137 {
		t.Errorf("Expected %v, got %v", 4.949999809265137, output)
	}
}

func TestGetDiscountPrice(t *testing.T) {
	beforeEach()
	var output float64
	utils.GetDiscountedPrice(c, &output)
	expected := 5.050000190734863
	handleError(c)

	visit(c)

	if output != expected {
		t.Errorf("Got %f. Expected %f", output, expected)
	}
}

func TestGetDescription(t *testing.T) {
	beforeEach()
	handleError(c)
	var output string
	expected := "Nick Diffatte has created a wonderful pocket illusion you can use ANYTIME you're going to pay for something.\n\nPull out a slip of Monopoly Money, and when the clerk informs you they don't accept that form of currency, visually change it in FULL VIEW."

	utils.GetDescription(c, &output)

	visit(c)

	if output != expected {
		t.Fatalf("Expected: %s, got %s", expected, output)
	}
}

func TestIfIntersting(t *testing.T) {

	utils.SetFilePath("~/python/penguin_bots/not_interested_products.csv")
	interesting := utils.IfInterested("jfdslfj")
	notInteresting := utils.IfInterested("Code Red by Cody Nottingham (DVD & Download)")
	if !interesting {
		t.Logf("Expected true, got %v", interesting)
	}

	if notInteresting {
		t.Logf("Expected false, got %v", notInteresting)
	}

	// delete files
	// if _, err := exec.Command("rm", "not_interested_products.csv").Output(); err != nil {
	// log.Fatal(err)
	// }

}

func TestAddNotInterestedProduct(t *testing.T) {
	utils.SetFilePath("~/python/penguin_bots/not_interested_products.csv")
	utils.AddNotInterestedProduct("jfdslfj")

	// delete test string from file again
	_, err := exec.Command("sed", "-i", "/jfdslfj/d", "../not_interested_products.csv").Output()
	if err != nil {
		panic(err)
	}
}
