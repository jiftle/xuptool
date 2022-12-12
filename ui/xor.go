package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gitee.com/yctxkj/xcrypto/xbit"
)

type XorGUI struct {
	lblPlain  *widget.Label
	txtPlain  *widget.Entry
	lblKey    *widget.Label
	txtKey    *widget.Entry
	lblResutl *widget.Label
	txtResult *widget.Entry
}

func NewGUI_XorGUI() *XorGUI {
	return &XorGUI{}
}

func (g *XorGUI) MakeUI(app fyne.App, w fyne.Window) fyne.CanvasObject {
	g.lblPlain = widget.NewLabel("数据1")
	g.txtPlain = &widget.Entry{Text: "00000A0100000A0102010203040506", PlaceHolder: "please intput plain ..."}
	g.lblKey = widget.NewLabel("数据2")
	g.txtKey = &widget.Entry{Text: "35403C9555ABB906F2D80A7FF41ED718", PlaceHolder: "please input key ..."}
	g.lblResutl = widget.NewLabel("结果")
	g.txtResult = &widget.Entry{Text: "", PlaceHolder: "this is result ! "}

	g.txtPlain.OnChanged = func(s string) {
		g.lblPlain.SetText(fmt.Sprintf("数据1[%v]", len(s)))
	}
	g.txtKey.OnChanged = func(s string) {
		g.lblPlain.SetText(fmt.Sprintf("数据2[%v]", len(s)))
	}

	cobj := container.NewVBox(
		g.lblPlain, g.txtPlain,
		g.lblKey, g.txtKey,
		g.lblResutl, g.txtResult,
		widget.NewButtonWithIcon("异或", theme.ConfirmIcon(), func() {
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sPlain := g.txtPlain.Text
			sKey := g.txtKey.Text
			sCipher := ""
			var err error
			sCipher, err = xbit.XOR(sPlain, sKey)
			if err != nil {
				g.txtResult.SetText(fmt.Sprintf("%v", err))
				return
			}
			g.txtResult.SetText(sCipher)

		}),
		widget.NewButtonWithIcon("清空", theme.DeleteIcon(), func() {
			g.txtResult.SetText("")
		}),
		widget.NewButtonWithIcon("退出", theme.CancelIcon(), func() {
			app.Quit()
		}),
	)
	return cobj
}
