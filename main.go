package main

import (
	"fmt"
	"utils"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/openbox", "www.penguinmagic.com/p/3901"),
		// colly.Async(true), // TODO: figure out async stuff
	)

	type Product struct {
		title       string
		description string
		price       float64
		discount    float64
		starRating  int64
	}
	product := Product{price: 0, discount: 0}

	utils.GetTitle(c, &product.title)
	utils.GetDescription(c, &product.description)

	c.Visit("https://www.penguinmagic.com/p/3901")
	fmt.Println(fmt.Sprintf("main product.description: %v", product.description)) // __AUTO_GENERATED_PRINT_VAR__

}
