package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gitee.com/yctxkj/xcrypto/xgm"
)

type diversify_sm4_gui struct {
	lblKey     *widget.Label
	txtKey     *widget.Entry
	lblDiverfy *widget.Label
	txtDiverfy *widget.Entry
	lblResult  *widget.Label
	txtResult  *widget.Entry
}

func NewGUI_diversify_sm4_gui() *diversify_sm4_gui {
	return &diversify_sm4_gui{}
}

func (g *diversify_sm4_gui) MakeUI(app fyne.App, w fyne.Window) fyne.CanvasObject {
	g.lblKey = widget.NewLabel("根密钥")
	g.txtKey = &widget.Entry{Text: "11111111111111111111111111111111", PlaceHolder: "please input key ..."}
	g.lblDiverfy = widget.NewLabel("分散因子")
	g.txtDiverfy = &widget.Entry{Text: "11111111111111111111111111111111", PlaceHolder: "please input key ..."}
	g.lblResult = widget.NewLabel("子密钥")
	g.txtResult = &widget.Entry{Text: "", PlaceHolder: "this is result ! "}

	cobj := container.NewVBox(
		g.lblKey, g.txtKey,
		g.lblDiverfy, g.txtDiverfy,
		g.lblResult, g.txtResult,
		widget.NewButtonWithIcon("密钥分散", theme.ConfirmIcon(), func() {
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sKey := g.txtKey.Text
			sDiverfy := g.txtDiverfy.Text
			sOut := ""
			var err error
			sOut, err = xgm.DiversifyKey(sKey, sDiverfy)
			if err != nil {
				g.txtResult.SetText(fmt.Sprintf("%v", err))
				return
			}
			g.txtResult.SetText(sOut)

		}),
		widget.NewButtonWithIcon("退出", theme.CancelIcon(), func() {
			app.Quit()
		}),
	)
	return cobj
}
