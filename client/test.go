package main

import (
	net "Faygo/client/module/net"
	"fmt"
)

func test() {
	
	c1 := Client{
		Task: "ls -ll /",
		
	}
	c1.Runcommand()
	

	
	testAes()

	
	

	
	fmt.Println(GetIP())
	fmt.Println(GetMac())

	
	fmt.Println(len(GetMac()))

	
	testHttpGet()

	
	testHttpPost()
}



















func testAes() {
	str := "test"

	enstr := AesEnCode(str)

	fmt.Println(enstr)

	nextstr := AesDeCode(enstr)

	fmt.Println(nextstr)

}


func testHttpGet() {
	url := "http:
	response := net.HttpGet(url)
	fmt.Println(response)
}

func testHttpPost() {
	url := "http:
	body := "{\"user\":\"user\",\"password\":\"password\"}"
	response := net.HttpPost(url, body)
	fmt.Println(response)
}
