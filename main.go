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

	var price float64
	utils.GetDiscountedPrice(c, &price)
	c.Visit("https://www.penguinmagic.com/p/3901")
	fmt.Println(fmt.Sprintf("main price: %v", price)) // __AUTO_GENERATED_PRINT_VAR__
}
