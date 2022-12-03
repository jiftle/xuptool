package algorithm

import (
	"crypto/des"
	"encoding/hex"
	"errors"
	"strconv"
)

const (
	// DECRYPT 解密模式
	DECRYPT = 0
	// ENCRYPT 加密模式
	ENCRYPT = 1
)

func DES_ECB_Decrypt(key, plain string) (cipher string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytPlain, err := hex.DecodeString(plain)
	if err != nil {
		return
	}
	bytCipher, err := DESECB(bytKey, bytPlain, 0)
	if err != nil {
		return
	}
	cipher = hex.EncodeToString(bytCipher)
	return
}

func DES_ECB_Encrypt(key, plain string) (cipher string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytPlain, err := hex.DecodeString(plain)
	if err != nil {
		return
	}
	bytCipher, err := DESECB(bytKey, bytPlain, 1)
	if err != nil {
		return
	}
	cipher = hex.EncodeToString(bytCipher)
	return
}

// DESECB single des for ecb mode
// flag	0：解密，1：加密
func DESECB(keyValue []byte, txtValue []byte, flag int) ([]byte, error) {
	key := make([]byte, 8)
	if len(keyValue) >= 8 {
		copy(key, keyValue[:8])
	} else {
		err := errors.New("key length must be greater than 8 bytes")
		return nil, err
	}
	if len(txtValue)%8 != 0 {
		err := errors.New("input length not multiple of 8 bytes")
		return nil, err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	data := txtValue
	out := make([]byte, len(txtValue))
	dst := out
	for len(data) > 0 {
		if flag == ENCRYPT {
			block.Encrypt(dst, data[:blockSize])
		} else {
			block.Decrypt(dst, data[:blockSize])
		}
		data = data[blockSize:]
		dst = dst[blockSize:]
	}

	return out, nil
}

// TripleDESECB triple des for ecb mode
// flag	0：解密，1：加密
func TripleDESECB(keyValue []byte, txtValue []byte, flag int) ([]byte, error) {
	key := make([]byte, 24)
	if len(keyValue) > 24 {
		copy(key, keyValue[:24])
	} else if len(keyValue) >= 16 {
		copy(key, keyValue[:16])
		copy(key[16:], keyValue[:8])
	} else {
		err := errors.New("key length must be greater than 16")
		return nil, err
	}
	if len(txtValue)%8 != 0 {
		err := errors.New("input length not multiple of 8 bytes")
		return nil, err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	data := txtValue
	out := make([]byte, len(txtValue))
	dst := out
	for len(data) > 0 {
		if flag == ENCRYPT {
			block.Encrypt(dst, data[:blockSize])
		} else {
			block.Decrypt(dst, data[:blockSize])
		}
		data = data[blockSize:]
		dst = dst[blockSize:]
	}

	return out, nil
}

// GetSubKeyDES 密钥分散
func GetSubKeyDES(ckKey []byte, dvsData []byte) ([]byte, error) {
	subData := make([]byte, 16)
	copy(subData[:8], dvsData[:8])
	for i := 0; i < 8; i++ {
		subData[8+i] = ^dvsData[i]
	}
	out, err := TripleDESECB(ckKey, subData, ENCRYPT)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcPBOCMacDES pboc mac 3des
func CalcPBOCMacDES(key []byte, iv []byte, data []byte, macLen int) ([]byte, error) {
	if macLen > 8 || macLen < 0 {
		return nil, errors.New("invalid mac length " + strconv.Itoa(macLen))
	}
	blockcount := len(data) / 8
	macDataSize := blockcount * 8
	macData := make([]byte, macDataSize+8)
	copy(macData, data)
	if len(data)%8 != 0 {
		blockcount++
		macDataSize += 8
		macData[len(data)] = 0x80
	}

	blockData := make([]byte, 8)
	ivData := make([]byte, 8)
	if iv != nil {
		copy(ivData, iv[:8])
	}

	tmpMacData := macData
	var err error
	for i := 0; i < blockcount; i++ {
		for j := 0; j < 8; j++ {
			blockData[j] = tmpMacData[j] ^ ivData[j]
		}
		ivData, err = DESECB(key[:8], blockData, ENCRYPT)
		if err != nil {
			return nil, err
		}
		tmpMacData = tmpMacData[8:]
	}
	blockData, err = DESECB(key[8:], ivData, DECRYPT)
	if err != nil {
		return nil, err
	}
	ivData, err = DESECB(key[:8], blockData, ENCRYPT)
	if err != nil {
		return nil, err
	}
	out := make([]byte, macLen)
	copy(out, ivData[:macLen])
	return out, nil
}
