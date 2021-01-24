package main

import "fmt"

func testinggettask() {
	t1 := Task{Mac: "zzz"}
	t1.GetTask()
	fmt.Println(t1)

}

func testinginit() {
	c1 := Client{Mac: "jjj"}
	result := InitClient(c1)
	fmt.Println(result)
}

func testupdate() {
	t1 := Task{Mac: "aaa", Status: 9, Result: "ok" /*, LatestTime: GetTime()*/}
	t1.UpdataTask()
	fmt.Println(t1)

}
