package main

import (
	"encoding/hex"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gitee.com/yctxkj/xuptool/algorithm"
)

type des_gui struct {
	app fyne.App

	rdoGroup  *widget.RadioGroup
	lblPlain  *widget.Label
	txtPlain  *widget.Entry
	txtKey    *widget.Entry
	txtResult *widget.Entry

	mode string // 算法模式 ECB CBC
}

func NewGUI_DES() *des_gui {
	return &des_gui{}
}

func (g *des_gui) MakeUI(app fyne.App) fyne.CanvasObject {
	g.lblPlain = widget.NewLabel("明文")
	g.txtResult = &widget.Entry{Text: "", PlaceHolder: "this is result ! "}
	g.txtPlain = &widget.Entry{Text: "1111111111111111", PlaceHolder: "please intput plain ..."}
	g.txtKey = &widget.Entry{Text: "1111111111111111", PlaceHolder: "please input key ..."}

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
	g.rdoGroup = widget.NewRadioGroup([]string{"ECB", "CBC"}, func(s string) {
		g.mode = s
	})
	cobj := container.NewVBox(
		widget.NewLabel("算法模式"),
		g.rdoGroup,
		g.lblPlain,
		g.txtPlain,
		widget.NewLabel("密钥"),
		g.txtKey,
		widget.NewLabel("结果"),
		g.txtResult,
		widget.NewButtonWithIcon("加密", theme.ConfirmIcon(), func() {
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sPlain := g.txtPlain.Text
			sKey := g.txtKey.Text
			sCipher := ""
			var err error
			if g.mode == "ECB" {
				sCipher, err = algorithm.DES_ECB_Encrypt(sKey, sPlain)
			} else {
				sCipher, err = algorithm.DES_CBC_Encrypt(sKey, sPlain)
			}
			if err != nil {
				g.txtResult.SetText(fmt.Sprintf("%v", err))
				return
			}
			g.txtResult.SetText(sCipher)
		}),
		widget.NewButtonWithIcon("解密", theme.CancelIcon(), func() {
			sPlain := g.txtPlain.Text
			sKey := g.txtKey.Text
			sCipher := ""
			var err error
			if g.mode == "ECB" {
				sCipher, err = algorithm.DES_ECB_Decrypt(sKey, sPlain)
			} else {
				sCipher, err = algorithm.DES_CBC_Decrypt(sKey, sPlain)
			}
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
	g.rdoGroup.SetSelected("ECB")
	return cobj
}
