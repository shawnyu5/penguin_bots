package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	a := app.New()
	mainWindow := a.NewWindow("Penguin")
	titleEntry := widget.NewEntry()

	products := widget.NewLabel("")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Product title", Widget: titleEntry},
		},
		OnSubmit: func() {
			log.Println("Submitted:", titleEntry.Text)
			go func() {
				getProducts(titleEntry.Text, products)
				fmt.Println(fmt.Sprintf("main products: %v", products)) // __AUTO_GENERATED_PRINT_VAR__
			}()
		},
	}

	// vBox := layout.NewVBoxLayout()
	content := container.NewVScroll(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			form,
			products,
		))

	content.Resize(fyne.NewSize(300, 300))
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()

}

// getProducts returns a list of products from the database matching the search query
func getProducts(searchQuery string, w *widget.Label) {

	type product struct {
		Title              string  `json:"title"`
		Appearances        int64   `json:"appearances"`
		AverageDiscount    float64 `json:"average_discount"`
		AveragePrice       float64 `json:"average_price"`
		DiscountPercentage int64   `json:"discount_percentage"`
	}

	var products []product

	res, err := http.Post(os.Getenv("API_URL")+"/search", "plain/text", strings.NewReader(searchQuery))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &products)

	productWidget := widget.NewLabel("")
	for i := 0; i < 20; i++ {
		curr := products[i]
		productWidget.Text += fmt.Sprintf("%+v\n\n", curr)
	}
	w.SetText(productWidget.Text)
	// return productWidget
}
