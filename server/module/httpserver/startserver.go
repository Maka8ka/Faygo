package httpserver

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	router := gin.Default()
	router.POST("/post", PostJson)
	router.POST("/login", LoginParms)
	router.GET("/", HttpDefault)
	router.POST("/", HttpPost)
	router.Run(":5000")
}

func PostJson(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"json": string(data),
		})
	}
}


func HttpDefault(c *gin.Context) {
	
	c.Redirect(302, "https:
}


func HttpPost(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.String(http.StatusOK, string(data))
		fmt.Println(string(data))
	}
}



type LoginForm struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginParms(c *gin.Context) {
	form := &LoginForm{}

	
	
	if c.BindJSON(&form) == nil {
		fmt.Println(form.User, form.Password)
		if form.User == "user" && form.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}
