package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gitee.com/yctxkj/xcrypto/xgm"
)

type gm_mac_gui struct {
	lblPlain  *widget.Label
	txtPlain  *widget.Entry
	lblKey    *widget.Label
	txtKey    *widget.Entry
	lblIv     *widget.Label
	txtIv     *widget.Entry
	lblResutl *widget.Label
	txtResult *widget.Entry
}

func NewGUI_gm_mac_gui() *gm_mac_gui {
	return &gm_mac_gui{}
}

func (g *gm_mac_gui) MakeUI(app fyne.App, w fyne.Window) fyne.CanvasObject {
	g.lblPlain = widget.NewLabel("MAC数据")
	g.txtPlain = &widget.Entry{Text: "00000A0100000A0102010203040506", PlaceHolder: "please intput plain ..."}
	g.lblKey = widget.NewLabel("密钥")
	g.txtKey = &widget.Entry{Text: "35403C9555ABB906F2D80A7FF41ED718", PlaceHolder: "please input key ..."}
	g.lblIv = widget.NewLabel("向量")
	g.txtIv = &widget.Entry{Text: "00000000000000000000000000000000"}
	g.lblResutl = widget.NewLabel("MAC")
	g.txtResult = &widget.Entry{Text: "", PlaceHolder: "this is result ! "}

	g.txtPlain.OnChanged = func(s string) {
		g.lblPlain.SetText(fmt.Sprintf("plain[%v]", len(s)))
	}

	cobj := container.NewVBox(
		g.lblPlain, g.txtPlain,
		g.lblKey, g.txtKey,
		g.lblIv, g.txtIv,
		g.lblResutl, g.txtResult,
		widget.NewButtonWithIcon("计算MAC", theme.ConfirmIcon(), func() {
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sPlain := g.txtPlain.Text
			sKey := g.txtKey.Text
			sIv := g.txtIv.Text
			sCipher := ""
			var err error
			sCipher, err = xgm.PbocMac(sPlain, sKey, sIv)
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
