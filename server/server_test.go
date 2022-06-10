package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocolly/colly"
	utils "github.com/shawnyu5/penguin-utils"
)

func TestLoggerHandler(t *testing.T) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/openbox/"),
	)
	product := Product{}

	utils.GetTitle(c, &product.Title)
	utils.GetDescription(c, &product.Description)
	utils.GetPrice(c, &product.Original_price)
	utils.GetDiscountedPrice(c, &product.Discount_price)
	utils.GetDiscountPercentage(c, &product.Discount_percentage)

	c.Visit("https://www.penguinmagic.com/openbox/")

	req := httptest.NewRequest(http.MethodGet, "/logger", nil)
	w := httptest.NewRecorder()
	loggerHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading body: %v", err)
	}
	fmt.Println(fmt.Sprintf("TestLoggerHandler data: %v", string(data))) // __AUTO_GENERATED_PRINT_VAR__
	j, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	if string(data) != string(j) {
		t.Errorf("Expected %v, got %v", string(j), string(data))
	}
}
