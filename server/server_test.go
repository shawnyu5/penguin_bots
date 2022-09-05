package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"server/utils"

	"github.com/gocolly/colly"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	Context("Logger Handler", func() {
		It("Should return the same product as the one on penguin magic currently", func() {
			c := colly.NewCollector(
				colly.AllowedDomains("www.penguinmagic.com", "www.penguinmagic.com/openbox/"),
			)

			// the product we are expecting to get back
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
				fmt.Println(fmt.Errorf("Error reading body: %v", err))
			}

			var loggerProduct LoggerProduct
			err = json.Unmarshal(data, &loggerProduct) // parse json data into struct
			if err != nil {
				fmt.Println(fmt.Errorf("Error unmarshalling json: %v", err))
			}

			// the product our logger handler returns should be the same as the product we got from penguinmagic
			Expect(loggerProduct).To(Equal(penguinProduct))
		})

		Context("Search Handler", func() {
			It("Should return a list of products", func() {
				// a successful search should return a list of products
				req := httptest.NewRequest(http.MethodPost, "/search", strings.NewReader("hello"))
				w := httptest.NewRecorder()
				searchHandler(w, req)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					panic(err)
				}

				// if we didnt get anything back, something has gone wrong
				Expect(len(data)).ToNot(Equal(0), "Expected data to not be empty")

			})

			It("Should return 404 given an invalid request method", func() {
				// a successful search should return a list of products
				req := httptest.NewRequest(http.MethodGet, "/search", nil)
				w := httptest.NewRecorder()
				searchHandler(w, req)
				res := w.Result()
				defer res.Body.Close()
				Expect(res.StatusCode).To(Equal(404), "Expected status code 404, got %v", res.StatusCode)
			})
		})
	})
})
