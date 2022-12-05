package ui

import (
	"encoding/hex"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gitee.com/yctxkj/xcrypto/xaes"
)

type aes_gui struct {
	lblPlain  *widget.Label
	txtPlain  *widget.Entry
	txtKey    *widget.Entry
	txtResult *widget.Entry
}

func NewGUI_AES() *aes_gui {
	return &aes_gui{}
}

func (g *aes_gui) MakeUI(app fyne.App, w fyne.Window) fyne.CanvasObject {
	g.lblPlain = widget.NewLabel("明文")
	g.txtResult = &widget.Entry{Text: "", PlaceHolder: "this is result ! "}
	g.txtPlain = &widget.Entry{Text: "11223344556677881122334455667788", PlaceHolder: "please intput plain ..."}
	g.txtKey = &widget.Entry{Text: "11223344556677881122334455667788", PlaceHolder: "please input key ..."}

	g.txtPlain.OnChanged = func(s string) {
		g.lblPlain.SetText(fmt.Sprintf("明文[%v]", len(s)))
	}
	g.txtPlain.Validator = func(s string) error {
		_, err := hex.DecodeString(s)
		if err != nil {
			return fmt.Errorf("")
		}
		if len(s) == 0 {
			return fmt.Errorf("")
		}
		if len(s)%16 != 0 {
			return fmt.Errorf("len error")
		}
		return nil
	}
	return container.NewVBox(
		g.lblPlain,
		g.txtPlain,
		widget.NewLabel("密钥"),
		g.txtKey,
		widget.NewLabel("密文"),
		g.txtResult,
		widget.NewButtonWithIcon("加密", theme.ConfirmIcon(), func() {
			// 加密
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sPlain := g.txtPlain.Text
			sKey := g.txtKey.Text
			sOut, err := xaes.AES_Encrypt_ECB(sPlain, sKey)
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}
			g.txtResult.SetText(sOut)

		}),
		widget.NewButtonWithIcon("解密", theme.CancelIcon(), func() {
			// 解密
			g.txtResult.SetText("")
			sPlain := g.txtPlain.Text
			sKey := g.txtKey.Text
			sOut, err := xaes.AES_Decrypt_ECB(sPlain, sKey)
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}
			g.txtResult.SetText(sOut)
		}),
		widget.NewButtonWithIcon("清空", theme.DeleteIcon(), func() {
			g.txtResult.SetText("")
		}),
		widget.NewButtonWithIcon("退出", theme.CancelIcon(), func() {
			app.Quit()
		}),
	)
}
