package main

import (
	net "Faygo/client/module/net"
	"fmt"
)

func test() {
	//命令执行
	c1 := Client{
		Task: "ls -ll /",
		//Result: "",
	}
	c1.Runcommand()
	//fmt.Println(c1.Result)

	//测试Aes加解密方法
	testAes()

	//测试Json Struct方法
	//testStruct()

	//测试 获取ip mac
	fmt.Println(GetIP())
	fmt.Println(GetMac())

	//判断mac长度 从而判断系统版本
	fmt.Println(len(GetMac()))

	//测试 http get
	testHttpGet()

	//测试 http post
	testHttpPost()
}

// //Examples

// func testStruct() {
// 	//测试 func Struct2Json()
// 	account1 := Client{
// 		Id:    2,
// 		Cname: "testclient",
// 	}
// 	fmt.Printf("%v\n", account1)
// 	account1json := Struct2Json(account1)
// 	fmt.Println(account1json)

// 	//测试 Json2Struct ,此处用到结构体方法和接受者相关内容
// 	var account2 Client
// 	account2.Json2Struct(account1json)
// 	fmt.Printf("%T%v\n", account2, account2)
// }

func testAes() {
	str := "test"

	enstr := AesEnCode(str)

	fmt.Println(enstr)

	nextstr := AesDeCode(enstr)

	fmt.Println(nextstr)

}

// 使用封装的http.Get
func testHttpGet() {
	url := "http://10.10.0.2:5000/"
	response := net.HttpGet(url)
	fmt.Println(response)
}

func testHttpPost() {
	url := "http://10.10.0.2:5000/login"
	body := "{\"user\":\"user\",\"password\":\"password\"}"
	response := net.HttpPost(url, body)
	fmt.Println(response)
}
