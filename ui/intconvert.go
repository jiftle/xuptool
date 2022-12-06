package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type IntcovertGUI struct {
	lblPlain  *widget.Label
	txtPlain  *widget.Entry
	lblResutl *widget.Label
	txtResult *widget.Entry
}

func NewGUI_IntcovertGUI() *IntcovertGUI {
	return &IntcovertGUI{}
}

func (g *IntcovertGUI) MakeUI(app fyne.App, w fyne.Window) fyne.CanvasObject {
	g.lblPlain = widget.NewLabel("数据")
	g.txtPlain = &widget.Entry{Text: "03E8", PlaceHolder: "please intput plain ..."}
	g.lblResutl = widget.NewLabel("结果")
	g.txtResult = &widget.Entry{Text: "", PlaceHolder: "this is result ! "}

	cobj := container.NewVBox(
		g.lblPlain, g.txtPlain,
		g.lblResutl, g.txtResult,
		widget.NewButtonWithIcon("16进制转10进制", theme.ConfirmIcon(), func() {
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sPlain := g.txtPlain.Text
			var err error

			n, err := strconv.ParseUint(sPlain, 16, 64)
			if err != nil {
				g.txtResult.SetText(err.Error())
				return
			}
			g.txtResult.SetText(fmt.Sprintf("%v", n))

		}),
		widget.NewButtonWithIcon("10进制转16进制", theme.ConfirmIcon(), func() {
			g.txtResult.SetText("")
			g.txtResult.Refresh()
			sNum := g.txtPlain.Text
			nNum, err := strconv.Atoi(sNum)
			if err != nil {
				g.txtResult.SetText(err.Error())
				return
			}
			g.txtResult.SetText(fmt.Sprintf("%X", nNum))

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
