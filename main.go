package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	//key, _ := hex.DecodeString(sKey)
	//plain, _ := hex.DecodeString(sPlain)

	//cipher, err := gm.EncryptECB(plain, key)
	//if err != nil {
	//log.Printf("加密失败, %v\n", err)
	//log.Fatal()
	//}

	//log.Printf("密文: %v\n", hex.EncodeToString(cipher))

	a := app.New()

	MainUI(a)
}
