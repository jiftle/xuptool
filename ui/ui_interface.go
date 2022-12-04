package ui

import "fyne.io/fyne/v2"

type XuptoolUI interface {
	MakeUI(app fyne.App, w fyne.Window) fyne.CanvasObject
}
