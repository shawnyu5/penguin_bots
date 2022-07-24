package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"server/utils"

	"github.com/gocolly/colly"
)

func TestLoggerHandler(t *testing.T) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/openbox/"),
	)

	penguinProduct := LoggerProduct{} // product on penguinmagic

	utils.GetTitle(c, &penguinProduct.Title)
	utils.GetPrice(c, &penguinProduct.Original_price)
	utils.GetDiscountedPrice(c, &penguinProduct.Discount_price)
	utils.GetDiscountPercentage(c, &penguinProduct.Discount_percentage)

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

	var loggerProduct LoggerProduct
	err = json.Unmarshal(data, &loggerProduct) // parse json data into struct
	if err != nil {
		t.Errorf("Error unmarshalling json: %v", err)
	}

	if loggerProduct != penguinProduct {
		t.Errorf("Expected %v, got %v", loggerProduct, data)
	}
}

// func TestCoinProductHandler(t *testing.T) {
// req := httptest.NewRequest(http.MethodGet, "/coinProduct", nil)
// w := httptest.NewRecorder()
// coinProductHandler(w, req)
// res := w.Result()
// defer res.Body.Close()

// data, err := ioutil.ReadAll(res.Body)
// if err != nil {
// t.Errorf("Error reading body: %v", err)
// }

// if data == nil {
// t.Errorf("Expected data, got nil")
// }
// }
