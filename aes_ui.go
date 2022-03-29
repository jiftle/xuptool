package main

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

func (g *aes_gui) MakeUI() fyne.CanvasObject {
	g.lblPlain = widget.NewLabel("plain")
	g.txtResult = &widget.Entry{Text: "", PlaceHolder: "this is result ! "}
	g.txtPlain = &widget.Entry{Text: "11223344556677881122334455667788", PlaceHolder: "please intput plain ..."}
	g.txtKey = &widget.Entry{Text: "11223344556677881122334455667788", PlaceHolder: "please input key ..."}

	g.txtPlain.OnChanged = func(s string) {
		g.lblPlain.SetText(fmt.Sprintf("plain[%v]", len(s)))
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
		widget.NewLabel("key"),
		g.txtKey,
		widget.NewLabel("cipher"),
		g.txtResult,
		widget.NewButtonWithIcon("encrypt", theme.ConfirmIcon(), func() {
			// 加密
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sPlain := g.txtPlain.Text
			//dialog.NewInformation("tip", fmt.Sprintf("plain: %v", sPlain), w).Show()
			bytPlain, err := hex.DecodeString(sPlain)
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}
			bytKey, err := hex.DecodeString(g.txtKey.Text)
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}
			bytCipher, err := xaes.Encrypt_ECB(bytPlain, bytKey)
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}
			sCipher := hex.EncodeToString(bytCipher)
			g.txtResult.SetText(sCipher)

		}),
		widget.NewButtonWithIcon("decrypt", theme.CancelIcon(), func() {
			// 解密
			g.txtResult.SetText("")
			sPlain := g.txtPlain.Text
			//dialog.NewInformation("tip", fmt.Sprintf("plain: %v", sPlain), w).Show()
			bytPlain, err := hex.DecodeString(sPlain)
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}
			bytKey, err := hex.DecodeString(g.txtKey.Text)
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}
			bytCipher, err := xaes.Decrypt_ECB(bytPlain, bytKey)
			if err != nil {
				dialog.NewError(err, w).Show()
				return
			}
			sCipher := hex.EncodeToString(bytCipher)
			g.txtResult.SetText(sCipher)
		}),
		widget.NewButtonWithIcon("reset", theme.DeleteIcon(), func() {
			g.txtResult.SetText("")
		}),
		widget.NewButtonWithIcon("exit", theme.DeleteIcon(), func() {
			w.Close()
		}),
	)
}
