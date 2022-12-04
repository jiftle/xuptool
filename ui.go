package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gitee.com/yctxkj/xuptool/ui"
)

var w fyne.Window

func MainUI(a fyne.App) {
	w = a.NewWindow("算法工具 --by jiftle 2022")

	// 多标签程序
	tabs := container.NewAppTabs(
		// 对称算法-通用算法
		container.NewTabItem("对称算法-DES", ui.NewGUI_DES().MakeUI(a, w)),
		container.NewTabItem("对称算法-3DES", ui.New_tripledes_gui().MakeUI(a, w)),
		container.NewTabItem("对称算法-AES", ui.NewGUI_AES().MakeUI(a, w)),
		// 对称算法-国密
		container.NewTabItem("国密算法-SM4", ui.NewGUI_SM4().MakeUI(a, w)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	w.Resize(fyne.NewSize(800, 600))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.ShowAndRun()
}
