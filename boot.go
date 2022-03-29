package main

import "os"

func init() {
	// linux 指定支持中文的字体
	os.Setenv("FYNE_FONT", "/usr/share/fonts/truetype/droid/DroidSansFallbackFull.ttf")
	os.Setenv("FYNE_THEME", "light")
}
