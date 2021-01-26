# Faygo

* A RAT Tools on major platforms.High scalability
* The traffic data be encryped by aes.use aeskey and aesiv
* The net moudle now use http,it can be replace to other protocol.
## Now support
* Linux
* Drwin
* Windows

## Example
click the gif for full expample video
[![Watch the video](https://raw.githubusercontent.com/Maka8ka/Faygo/main/faygo.gif)](https://github.com/Maka8ka/Faygo/blob/main/faygo.m4v)
## Detection
[Watch the video](https://raw.githubusercontent.com/Maka8ka/Faygo/main/detection.jpg)

## Client

### module

- package:cipher

	- aescbc.go

		- func AecCbcEncrypt(text string,[]byte(aeskey))(text string,err){}
		- func AecCbcDecrypt(text string,[]byte(aeskey))(text string,err){}

	- structjson.go

- package:net

	- http.go

- package:command

	- command.go

- package:getmac

### main.go

### cronjob.go

### struct.go

- type Client struct {
	Id         int    `json:"id"` //clientid
	Cname      string `json:"came"`
	Lan        string `json:"lan"`
	Wan        string `json:"wan"`
	OsType     string `json:"os"`
	Mac        string `json:"mac"`
	LatestTime string `json:"latest_time"`
	Status     int    `json:"status"`
	Task       string `json:"task"`
	Result     string `json:"result"`
	FileHex    string `json:"filehex"`
	FileLo     string `json:"filelo"`
	Notes      string `json:"notes"`
}

### func.go

### const.go(const settings)

## Server

### module

- package:cipher

	- aescbc.go

		- func AecCbcEncrypt(text string,[]byte(aeskey))(text string,err){}
		- func AecCbcDecrypt(text string,[]byte(aeskey))(text string,err){}

	- structjson.go

- package:database

	- mysql-initialize.sql

- package:httpserver

### main.go

### struct.go

- type Client struct {
	Id         int    `json:"id"` //clientid
	Cname      string `json:"came"`
	Lan        string `json:"lan"`
	Wan        string `json:"wan"`
	OsType     string `json:"os"`
	Mac        string `json:"mac"`
	LatestTime string `json:"latest_time"`
	Status     int    `json:"status"`
	Task       string `json:"task"`
	Result     string `json:"result"`
	FileHex    string `json:"filehex"`
	FileLo     string `json:"filelo"`
	Notes      string `json:"notes"`
}

### const.go(const-settings)

