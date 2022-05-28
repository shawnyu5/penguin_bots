package main

import (
	"check_coin_product"
	"fmt"
	"log"
	"net/http"
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
	fmt.Println("LISTENING ON PORT 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
