package main

import (
	"fmt"
	"utils"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
		// colly.Async(true), // TODO: figure out async stuff
	)

	rating := utils.GetStarRating(c)
	percentage := utils.GetDiscountPercentage(c)
	fmt.Println(fmt.Sprintf("main percentage: %v", percentage)) //__AUTO_GENERATED_PRINT_VAR__
	fmt.Println(fmt.Sprintf("main output: %v", rating))         // __AUTO_GENERATED_PRINT_VAR__
	fmt.Println(fmt.Sprintf("main output2: %v", percentage))    // __AUTO_GENERATED_PRINT_VAR__
	c.Visit("https://www.penguinmagic.com/p/3901")
}
