package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var w fyne.Window

func MainUI(a fyne.App) {
	w = a.NewWindow("算法工具 --by jiftle 2022")

	tabs := container.NewAppTabs(
		container.NewTabItem("SM4", NewGUI_SM4().MakeUI()),
		container.NewTabItem("AES", NewGUI_AES().MakeUI()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	w.Resize(fyne.NewSize(600, 400))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.ShowAndRun()
}
