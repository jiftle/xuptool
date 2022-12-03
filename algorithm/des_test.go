package algorithm

import (
	"fmt"
	"testing"
)

func Test_DES_CBC_Encrypt(t *testing.T) {
	cipher, err := DES_CBC_Encrypt("1111111111111111", "1111111111111111")
	fmt.Println(cipher, err)
}

func Test_DES_ECB_Encrypt(t *testing.T) {
	sKey := "1111111111111111"
	sPlain := "1111111111111111"
	sCipherExpect := "f40379ab9e0ec533"

	sCipher, _ := DES_ECB_Encrypt(sKey, sPlain)
	if sCipher != sCipherExpect {
		t.Fatalf("DES ECB 失败")
	}
}
