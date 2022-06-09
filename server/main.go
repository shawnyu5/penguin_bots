package main

import (
	"check_coin_product"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

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

func main() {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	routes["/"] = homeHandler(routes)
	routes["/coinProduct"] = coinProductHandler
	routes["/favicon.ico"] = doNothing
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
