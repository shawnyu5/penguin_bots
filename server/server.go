package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"server/middleware"
	"server/search"
	"server/utils"

	"git.mills.io/prologic/bitcask"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

type LoggerProduct struct {
	Title               string  `json:"title"`
	Original_price      float64 `json:"original_price"`
	Discount_price      float64 `json:"discount_price"`
	Discount_percentage float64 `json:"discount_percentage"`
}

type CoinProductService interface {
}

type dbTypes struct{}

func (db dbTypes) coin_product_title() string {
	return "coin_product_title"
}

func (db dbTypes) product_title() string {
	return "product_title"
}

var storage *bitcask.Bitcask

func main() {
	// initialize the cache
	b, err := bitcask.Open("./db")
	if err != nil {
		log.Fatal(err)
	}
	storage = b
	routes := make(map[string]middleware.LoggerInter)
	routes["/"] = middleware.NewLogger(homeHandler(routes))
	routes["/logger"] = middleware.NewLogger(loggerHandler)
	routes["/search"] = middleware.NewLogger(searchHandler)
	routes["/favicon.ico"] = middleware.NewLogger(doNothing)
	for route, handler := range routes {
		http.HandleFunc(route, handler.ServeHTTP)
	}

	// load .env
	err = godotenv.Load()
	if err != nil {
		log.Println("(server) Error loading .env file")
	}
	// get port from env
	port := ":" + os.Getenv("PORT")
	// set default port to 8080
	if port == ":" {
		port = ":8080"
	}
	fmt.Println("LISTENING ON PORT " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// doNothing is a do nothing function
func doNothing(w http.ResponseWriter, r *http.Request) {}

// homeHandler handles the home route
func homeHandler(routes map[string]middleware.LoggerInter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		var list string
		for k := range routes {
			list += k + "\n"
		}
		fmt.Fprintln(w, list)
	}
}

// loggerHandler handles the /logger route
func loggerHandler(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/openbox/", "https://www.penguinmagic.com/p/12449"),
	)
	product := LoggerProduct{}

	utils.GetTitle(c, &product.Title)
	utils.GetPrice(c, &product.Original_price)
	utils.GetDiscountedPrice(c, &product.Discount_price)
	utils.GetDiscountPercentage(c, &product.Discount_percentage)

	// c.Visit("https://www.penguinmagic.com/p/17318")
	// c.Visit("https://www.penguinmagic.com/p/12449")
	c.Visit("https://www.penguinmagic.com/openbox/")

	types := dbTypes{}
	if storage != nil {
		storage.Put([]byte(types.product_title()), []byte(product.Title))
		// storage.Set("product_title", product.Title, cache.DefaultExpiration)
	}
	j, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(j))
}

// searchHandler is the handler for the /search endpoint
// Returns a list of products that match the search query in json
func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method "+r.Method+" not allowed", 404)
		return
	}

	logger := log.New(os.Stdout, "", log.LUTC)

	var s search.SearchService
	s = search.SearchServiceImpl{}
	s = search.LoggingMiddleware{Logger: logger, Next: s}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	product := search.Product{Title: string(body)}
	result, err := s.SearchByRegex(&product)
	if err != nil {
		panic(err)
	}

	j, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(j))
}
