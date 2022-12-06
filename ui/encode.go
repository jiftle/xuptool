package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gitee.com/yctxkj/xcrypto/xencoding"
)

type EncodeGUI struct {
	lblPlain  *widget.Label
	txtPlain  *widget.Entry
	lblResutl *widget.Label
	txtResult *widget.Entry
}

func NewGUI_EncodeGUI() *EncodeGUI {
	return &EncodeGUI{}
}

func (g *EncodeGUI) MakeUI(app fyne.App, w fyne.Window) fyne.CanvasObject {
	g.lblPlain = widget.NewLabel("数据")
	g.txtPlain = &widget.Entry{Text: "68656C6C6F", PlaceHolder: "please intput plain ..."}
	g.lblResutl = widget.NewLabel("结果")
	g.txtResult = &widget.Entry{Text: "", PlaceHolder: "this is result ! "}

	g.txtPlain.OnChanged = func(s string) {
		g.lblPlain.SetText(fmt.Sprintf("数据1[%v]", len(s)))
	}

	cobj := container.NewVBox(
		g.lblPlain, g.txtPlain,
		g.lblResutl, g.txtResult,
		widget.NewButtonWithIcon("解码", theme.ConfirmIcon(), func() {
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sPlain := g.txtPlain.Text
			sCipher := ""
			var err error
			sCipher, err = xencoding.HexStr2Utf8Str(sPlain)
			if err != nil {
				g.txtPlain.SetText(err.Error())
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
