package main

import (
	"check_coin_product"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 404)
		return
	}
	productInfo := check_coin_product.Check("https://www.penguinmagic.com/openbox/")
	fmt.Fprintf(w, productInfo)
}
func main() {
	http.HandleFunc("/", homeHandler)
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
