package main

import (
	"os"
	"runtime"
	"strings"

	"github.com/flopp/go-findfont"
)

func init() {
	switch runtime.GOOS {
	case "windows":
		setFont_Win()
	case "linux":
		setFont_linux()
	}

}

func setFont_linux() {
	// linux 指定支持中文的字体
	os.Setenv("FYNE_FONT", "/usr/share/fonts/truetype/droid/DroidSansFallbackFull.ttf")
	os.Setenv("FYNE_THEME", "light")
}

func setFont_Win() {
	os.Setenv("FYNE_THEME", "dark")

	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}
