package main

import (
	"fmt"
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
	case "darwin":
		setFont_macos()
	default:
		fmt.Printf("未定义分支, %v\n", runtime.GOOS)
	}

}

func setFont_macos() {
	os.Setenv("FYNE_THEME", "light")
	os.Setenv("FYNE_FONT", "/System/Library/Fonts/Supplemental/Songti.ttc")
}

func setFont_linux() {
	// linux 指定支持中文的字体
	os.Setenv("FYNE_FONT", "/usr/share/fonts/truetype/droid/DroidSansFallbackFull.ttf")
	os.Setenv("FYNE_THEME", "light")
	// os.Unsetenv("FYNE_FONT")
	// fontPaths := findfont.List()
	// for _, path := range fontPaths {
	// 	fmt.Println(path)
	// 	// defaultFont := "MesloLGS NF Bold Italic.ttf"
	// 	// if strings.Contains(path, defaultFont) {
	// 	if strings.Contains(path, "MesloLGS") {
	// 		os.Setenv("FYNE_FONT", path)
	// 		// break
	// 	}
	// }
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
