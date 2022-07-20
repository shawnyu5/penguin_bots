package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	mainWindow := a.NewWindow("Penguin")
	titleEntry := widget.NewEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Product title", Widget: titleEntry},
		},
		OnSubmit: func() {
			log.Println("Submitted:", titleEntry.Text)
		},
	}

	mainWindow.SetContent(form)
	mainWindow.ShowAndRun()

}
