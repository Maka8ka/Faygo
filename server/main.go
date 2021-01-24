package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	StartHttpServer()
}


func DealRequests(body string) string {
	
	fmt.Println("requests body", AesDeCode(body))

	
	posttask := Task{}
	posttask.Json2Struct(AesDeCode(body))
	posttask.UpdataTime()

	switch posttask.Status {
	case 0:
		
		ok, _ := GetStatus(posttask) 
		if !ok {
			poststruct := Client{}
			poststruct.Json2Struct(AesDeCode(body))
			if InitClient(poststruct) {
				fmt.Println("init success")
				return "init success"
			} else {
				fmt.Println("init error")
				return "int error"
			}

		} else {
			fmt.Println("mac already init,please sleep")
			
			
			

			return ReturnTask(posttask)
		}
	case 1:
		
		_, status := GetStatus(posttask)
		switch status {
		case 1:
			transstruct := Task{Mac: posttask.Mac}
			transstruct.GetTask()
			
			fmt.Println("case 1-1")
			return AesEnCode(Struct2Json(transstruct))
		default:
			
			fmt.Println("case 1-d")
			return "case1 default"
		}

	default:
		
		fmt.Println("get default")
		return ReturnTask(posttask)
	}

}

func ReturnTask(task Task) string {
	transstruct := Task{Mac: task.Mac}
	transstruct.GetTask()
	return AesEnCode(Struct2Json(transstruct))
}

func GetStatus(poststruct Task) (bool, int) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Table("client").Where("mac = ?", poststruct.Mac).Find(poststruct)
	if poststruct.Status == 0 {
		return false, 0
	} else {
		return true, poststruct.Status
	}
}
