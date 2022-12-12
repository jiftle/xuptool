package ui

import (
	"encoding/hex"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gitee.com/yctxkj/xcrypto/xgm"
)

type sm4_gui struct {
	lblMode         *widget.Label // 模式， ECB, CBC
	rdoGroupMode    *widget.RadioGroup
	lblPadding      *widget.Label
	rdoGroupPadding *widget.RadioGroup // 补齐规则

	lblPlain  *widget.Label
	txtPlain  *widget.Entry
	txtKey    *widget.Entry
	txtResult *widget.Entry
	lblIv     *widget.Label
	txtIv     *widget.Entry

	mode    string // 算法模式 ECB CBC
	padding string // 补齐规则
}

func NewGUI_sm4_gui() *sm4_gui {
	return &sm4_gui{}
}

func (g *sm4_gui) MakeUI(app fyne.App, w fyne.Window) fyne.CanvasObject {
	g.lblPlain = widget.NewLabel("明文")
	g.txtResult = &widget.Entry{Text: "", PlaceHolder: "this is result ! "}
	g.txtPlain = &widget.Entry{Text: "00000000000000000A01020304050680", PlaceHolder: "please intput plain ..."}
	g.txtKey = &widget.Entry{Text: "60EB9BF035B849CC2EE26BBEC22C20B1", PlaceHolder: "please input key ..."}
	g.lblIv = widget.NewLabel("向量")
	g.txtIv = &widget.Entry{Text: "00000000000000000000000000000000"}
	g.lblMode = widget.NewLabel("模式")
	g.rdoGroupMode = widget.NewRadioGroup([]string{"ECB", "CBC"}, func(s string) {
		g.mode = s
	})
	g.rdoGroupMode.SetSelected("ECB")
	g.rdoGroupMode.Horizontal = true

	g.txtPlain.OnChanged = func(s string) {
		g.lblPlain.SetText(fmt.Sprintf("明文[%v]", len(s)))
	}
	g.txtPlain.Validator = func(s string) error {
		if len(s) == 0 {
			return fmt.Errorf("请输入数据")
		}
		if len(s)%16 != 0 {
			return fmt.Errorf("长度错误")
		}
		_, err := hex.DecodeString(s)
		if err != nil {
			return fmt.Errorf("格式错误")
		}
		return nil
	}

	cobj := container.NewVBox(
		g.lblMode, g.rdoGroupMode,
		g.lblPlain,
		g.txtPlain,
		widget.NewLabel("密钥"),
		g.txtKey,
		g.lblIv, g.txtIv,
		widget.NewLabel("密文"),
		g.txtResult,
		widget.NewButtonWithIcon("加密", theme.ConfirmIcon(), func() {
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sPlain := g.txtPlain.Text
			sKey := g.txtKey.Text
			sIv := g.txtIv.Text
			sCipher := ""
			var err error
			if g.mode == "ECB" {
				sCipher, err = xgm.Encrypt_ECB(sPlain, sKey)
			} else {
				sCipher, err = xgm.Encrypt_CBC(sPlain, sKey, sIv)
			}
			if err != nil {
				g.txtResult.SetText(fmt.Sprintf("%v", err))
				return
			}
			g.txtResult.SetText(sCipher)

		}),
		widget.NewButtonWithIcon("解密", theme.CancelIcon(), func() {
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sPlain := g.txtPlain.Text
			sKey := g.txtKey.Text
			sIv := g.txtIv.Text
			sCipher := ""
			var err error
			if g.mode == "ECB" {
				sCipher, err = xgm.Decrypt_ECB(sPlain, sKey)
			} else {
				sCipher, err = xgm.Decrypt_CBC(sPlain, sKey, sIv)
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
	return cobj
}
