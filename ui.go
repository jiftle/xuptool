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
		container.NewTabItem("国密算法-SM4", ui.NewGUI_sm4_gui().MakeUI(a, w)),
		// 国密密钥分散算法
		container.NewTabItem("国密-密钥分散", ui.NewGUI_diversify_sm4_gui().MakeUI(a, w)),
		// 国密MAC算法
		container.NewTabItem("国密-MAC算法", ui.NewGUI_gm_mac_gui().MakeUI(a, w)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	// if runtime.GOOS == "linux" {
	// w.Resize(fyne.NewSize(600, 400))
	// } else {
	w.Resize(fyne.NewSize(800, 600))
	// }

	// a.Settings().SetTheme(&xtheme.MyTheme{})

	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.ShowAndRun()
}
