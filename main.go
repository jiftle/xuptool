//go:generate fyne bundle -o data.go Icon.png
package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	a.SetIcon(resourceIconPng)
	MainUI(a)
}
