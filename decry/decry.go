package decry

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func PKCSUnpadding(input []byte) []byte {
	length := len(input)
	if length == 0 {
		return input
	}
	unpad := int(input[length-1])
	return input[:len(input)-unpad]
}

func AesDecryptCBC(input, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("Err:%v\n", err)
		return []byte{}
	}

	blockSize := block.BlockSize()
	iv := key[:blockSize]
	if len(input)%blockSize != 0 {
		return []byte{}
	}
	//iv := bytes.Repeat([]byte{byte(0)}, blockSize)
	blockMode := cipher.NewCBCDecrypter(block, iv)

	output := make([]byte, len(input))
	blockMode.CryptBlocks(output, input)

	unpadText := PKCSUnpadding(output)

	fmt.Printf("Decry Information:\n")
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	fmt.Printf("key(len:%d):%s\niv(len:%d):%s\ninput(len:%d):%v\noutput(len:%d):%v\ntext(len:%d):%s\n",
		len(key), key, len(iv), iv, len(input), input, len(output), output, len(unpadText), unpadText)
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	return unpadText
}

func AesDecrypt(input []byte, key []byte, mode string) []byte {
	if mode == "CBC" {
		return AesDecryptCBC(input, key)
	}

	return []byte{}
}
