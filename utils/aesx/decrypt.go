package aesx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func Decrypt(key, ciphertext []byte) ([]byte, error) {
	decodeChiperText, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(decodeChiperText) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := decodeChiperText[:aes.BlockSize]
	decodeChiperText = decodeChiperText[aes.BlockSize:]

	stream := cipher.NewCBCDecrypter(block, iv)
	stream.CryptBlocks(decodeChiperText, decodeChiperText)

	return unpad(decodeChiperText)
}

// PKCS7 Unpadding
func unpad(buf []byte) ([]byte, error) {
	length := len(buf)
	if length == 0 {
		return nil, fmt.Errorf("unpad error: input data is empty")
	}

	padLen := int(buf[length-1])
	if padLen > length || padLen > aes.BlockSize {
		return nil, fmt.Errorf("unpad error: invalid padding size")
	}

	return buf[:length-padLen], nil
}
