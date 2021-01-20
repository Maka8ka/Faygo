package main

//Client struct
type Client struct {
	Id         int    //clientid
	Note       string //notename
	OsType     string //windows,linux,drwin
	MacAddress string //remotemac
	Status     int    //0,1,2,3,4
	Lasttask   int    //latest taskid
}

type Task struct {
	Id            int
	ClientId      int    //-> Client.Id
	CommandLine   string // -> cmd task
	CommandResult string // -> command result
	CommandStatus int    //0,1
}

type FileTransfer struct {
	Id             int    //FileID
	ClientId       int    //->Client.Id
	ReomteLocation string // "/tmp/test.file"
	FileHex        string // hexed file
	Status         int    // 0,1 TransferStatus
}


func main{

}