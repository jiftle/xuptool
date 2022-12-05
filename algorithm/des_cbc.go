package algorithm

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"errors"
)

func DES_CBC_Encrypt(key, plain string) (cipher string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytPlain, err := hex.DecodeString(plain)
	if err != nil {
		return
	}
	bytCipher, err := DesEncrypt(bytPlain, bytKey)
	if err != nil {
		return
	}
	cipher = hex.EncodeToString(bytCipher)
	return
}

func DES_CBC_Decrypt(key, plain string) (cipher string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytPlain, err := hex.DecodeString(plain)
	if err != nil {
		return
	}
	bytCipher, err := DesDecrypt(bytPlain, bytKey)
	if err != nil {
		return
	}
	cipher = hex.EncodeToString(bytCipher)
	return
}

func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(origData)%8 != 0 {
		origData = PKCS5Padding(origData, block.BlockSize())
	}
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

// 3DES加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func TripleDES_CBC(key, data, iv []byte, flag int) (out []byte, err error) {
	nkey := make([]byte, 24)
	if len(key) == 16 {
		copy(nkey, key)
		copy(nkey[16:], key[:8])
	} else if len(key) == 24 {
		copy(nkey, key)
	}
	block, err := des.NewTripleDESCipher(nkey)
	if err != nil {
		return
	}
	var blockMode cipher.BlockMode
	if flag == 0 {
		blockMode = cipher.NewCBCEncrypter(block, iv)
	} else {
		blockMode = cipher.NewCBCDecrypter(block, iv)
	}
	out = make([]byte, len(data))
	blockMode.CryptBlocks(out, data)
	return
}

func TripleDES_CBC_Encrypt(key, plain, iv string) (cipher string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytPlain, err := hex.DecodeString(plain)
	if err != nil {
		return
	}
	bytIv, err := hex.DecodeString(iv)
	if err != nil {
		return
	}
	if len(key) != 32 {
		err = errors.New("key length must be 16")
		return
	}
	if len(iv) != 16 {
		err = errors.New("iv length must be 8")
		return
	}
	bytCipher, err := TripleDES_CBC(bytKey, bytPlain, bytIv, 1)
	if err != nil {
		return
	}
	cipher = hex.EncodeToString(bytCipher)
	return
}

func TripleDES_CBC_Decrypt(key, plain, iv string) (cipher string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytPlain, err := hex.DecodeString(plain)
	if err != nil {
		return
	}
	bytIv, err := hex.DecodeString(iv)
	if err != nil {
		return
	}
	if len(key) != 32 {
		err = errors.New("key length must be 16")
		return
	}
	if len(iv) != 16 {
		err = errors.New("iv length must be 8")
		return
	}
	bytCipher, err := TripleDES_CBC(bytKey, bytPlain, bytIv, 0)
	if err != nil {
		return
	}
	cipher = hex.EncodeToString(bytCipher)
	return
}
