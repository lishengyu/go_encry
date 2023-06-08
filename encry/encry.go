package encry

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func PKCSPadding(input []byte, blockSize int) []byte {
	padding := blockSize - len(input)%blockSize

	var padText []byte
	if padding == 0 {
		padText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		padText = bytes.Repeat([]byte{byte(padding)}, padding)
	}

	return append(input, padText...)
}

func AesEncryptCBC(input, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("Err:%v\n", err)
	}

	blockSize := block.BlockSize()
	padText := PKCSPadding(input, blockSize)

	iv := key[:blockSize]
	//iv := bytes.Repeat([]byte{byte(0)}, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	output := make([]byte, len(padText))
	blockMode.CryptBlocks(output, padText)

	fmt.Printf("Encry Information:\n")
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	fmt.Printf("key(len:%d):%s\niv(len:%d):%s\ninput(len:%d):%v\noutput(len:%d):%v\n",
		len(key), key, len(iv), iv, len(input), input, len(output), output)
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")

	return output
}

func AesEncrypt(input []byte, key []byte, mode string) []byte {
	if mode == "CBC" {
		return AesEncryptCBC(input, key)
	}

	return []byte{}
}
