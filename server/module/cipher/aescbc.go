package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"log"
	"runtime"
)

const (
	ivaes  = "Kasy65xGUhjbzg5f" 
	aeskey = "qqqqqqqqqqqqqqqq" 
)

func PKCS5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func PKCS5UnPadding(plainText []byte) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number > length {
		return nil, nil 
	}
	return plainText[:length-number], nil
}

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}


func AesCbcEncrypt(plainText, key []byte, ivAes ...byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	paddingText := PKCS5Padding(plainText, block.BlockSize())

	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != 16 {
			return nil, nil 
		} else {
			iv = ivAes
		}
	} else {
		iv = []byte(ivaes)
	} 
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}


func AesCbcDecrypt(cipherText, key []byte, ivAes ...byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Key error")
			default:
				log.Println("error:", err)
			}
		}
	}()
	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != 16 {
			return nil, nil
		} else {
			iv = ivAes
		}
	} else {
		iv = []byte(ivaes)
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	paddingText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(paddingText, cipherText)

	plainText, err := PKCS5UnPadding(paddingText)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
