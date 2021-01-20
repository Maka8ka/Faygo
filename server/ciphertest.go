package main

import (
	"encoding/base64"
	"fmt"

	cipher "Faygo/server/cipher"
)

const aeskey = "whoisyourdaddy77"

func main() {
	plaintext := []byte("whois yourdaddy")
	fmt.Println("plaintext：", string(plaintext))

	// 传入明文和自己定义的密钥，密钥为16字节 可以自己传入初始化向量,如果不传就使用默认的初始化向量,16字节
	cryptText, err := cipher.AesCbcEncrypt(plaintext, []byte(aeskey))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("encrypt_text:", base64.StdEncoding.EncodeToString(cryptText))

	// 传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错 可以自己传入初始化向量,如果不传就使用默认的初始化向量,16字节
	newplaintext, err := cipher.AesCbcDecrypt(cryptText, []byte(aeskey))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("decrypt_text", string(newplaintext))
}
