package main

import (
	"check_coin_product"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	utils "github.com/shawnyu5/penguin-utils"
)

func main() {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	routes["/"] = homeHandler(routes)
	routes["/coinProduct"] = coinProductHandler
	routes["/favicon.ico"] = doNothing
	routes["/logger"] = loggerHandler
	for k, v := range routes {
		http.HandleFunc(k, v)
	}

	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// get from env
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	fmt.Println("LISTENING ON PORT " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// coinProductHandler is the handler for the /coinProduct endpoint
func coinProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 404)
		return
	}
	productInfo := check_coin_product.Check("https://www.penguinmagic.com/openbox/")
	fmt.Fprintf(w, productInfo)
}

// doNothing is a do nothing function
func doNothing(w http.ResponseWriter, r *http.Request) {}

func homeHandler(routes map[string]func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var list string
		for k := range routes {
			list += k + "\n"
		}
		fmt.Fprintln(w, list)
	}

}

func loggerHandler(w http.ResponseWriter, r *http.Request) {
	type Product struct {
		Title               string  `json:"title"`
		Description         string  `json:"description"`
		Original_price      float64 `json:"original_price"`
		Discount_price      float64 `json:"discount_price"`
		Discount_percentage float64 `json:"discount_percentage"`
	}
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/openbox/"),
	)
	product := Product{}

	utils.GetTitle(c, &product.Title)

	fmt.Println(fmt.Sprintf("loggerHandler product.Title: %+v", product)) // __AUTO_GENERATED_PRINT_VAR__
	c.Visit("https://www.penguinmagic.com/openbox/")
}
