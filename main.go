package main

import (
	"fmt"
	"utils"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/p/3901"),
	)

	output := utils.GetDiscountPercentage(c)
	fmt.Println(fmt.Sprintf("main output: %v", output)) // __AUTO_GENERATED_PRINT_VAR__
	c.Visit("https://www.penguinmagic.com/p/3901")

}
