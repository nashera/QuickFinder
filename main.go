package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {fmt.Println("Menu New")}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Cut", func(){fmt.Println("Menu Cut")}),
			fyne.NewMenuItem("Copy", func(){fmt.Println("Menu Copy")}),
			fyne.NewMenuItem("Paste", func(){fmt.Println("Menu Paste")}),
		),
		fyne.NewMenu("Tool",
			fyne.NewMenuItem("setting", func() {fmt.Println("get setting")}),
		),		
	))

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	w.ShowAndRun()
}
