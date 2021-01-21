# Faygo

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

### const.go(常量配置)

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

### const.go(常量配置)

*XMind - Trial Version*
