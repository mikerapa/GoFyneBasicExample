package main

import (
	"fmt"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	count := 0

	countLabel := widget.NewLabel("Count")
	countButton := widget.NewButton("Add 1", func() {
		count = count + 1
		countLabel.SetText(fmt.Sprintf("count: %d", count))
	})
	w := a.NewWindow("Counter...")
	w.SetContent(widget.NewVBox(
		countLabel,
		countButton,
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	w.ShowAndRun()
}
