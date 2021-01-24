package main

import (
	"encoding/base64"
	"fmt"

	cipher "Faygo/client/module/cipher"
)

const aeskey = "whoisyourdaddy77"

func main() {
	plaintext := []byte("whois yourdaddy")
	fmt.Println("plaintext：", string(plaintext))

	
	cryptText, err := cipher.AesCbcEncrypt(plaintext, []byte(aeskey))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("encrypt_text:", base64.StdEncoding.EncodeToString(cryptText))

	
	newplaintext, err := cipher.AesCbcDecrypt(cryptText, []byte(aeskey))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("decrypt_text", string(newplaintext))
}
