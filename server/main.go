package mainimport (
	"Faygo/server/module/file"
	"fmt"
	"strings"	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)func main() {
	StartHttpServer()
}
func DealRequests(body string) string { 
	
	
	acquirebodystruct := Task{}
	acquirebodystruct.Json2Struct(AesDeCode(body))
	switch acquirebodystruct.Status {
	case 0: 
		return CaseZero(acquirebodystruct)
	case 1: 
		return CaseOne(acquirebodystruct)
	case 2: 
		fmt.Println("case2")
		return CaseTwo(acquirebodystruct)
	case 3: 
		fmt.Println("case3")
		return CaseThree(acquirebodystruct)
	case 4: 
		fmt.Println("case4")
		return CaseFour(acquirebodystruct)
	case 5: 
		fmt.Println("case5")
		return CaseTwo(acquirebodystruct)
	default:
		return "error"
	}}func CaseFour(t Task) string {
	if t.FileHex != "" {
		filelo := fmt.Sprintf(strings.Replace(t.FileLo, "/", ".", -1))
		filelo = "/home/user/go/src/Faygo/server/upload/" + fmt.Sprintf(strings.Replace(filelo, ":", "-", -1))
		fmt.Println("文件保存在",filelo)
		file.HexToFile(t.FileHex, filelo)
		fmt.Println("客户端文件读取成功")
	} else {
		fmt.Println("客户端文件读取失败")
	}
	t.UpdataTask()
	return ReturnTask(t)
}func CaseThree(t Task) string {
	if t.FileHex == "文件写入成功" {
		fmt.Println("客户端文件写入成功")
	} else {
		fmt.Println("客户端文件写入失败")
	}
	t.UpdataTask()
	return ReturnTask(t)
}func CaseZero(t Task) string {
	if GetStatus(t) == 0 {
		client := Client{}
		client.Json2Struct(Struct2Json(t))
		if InitClient(client) {
			fmt.Println("init success")
			return ReturnTask(t)
		} else {
			fmt.Println("init error")
			return ReturnTask(t)
		}	} else {
		fmt.Println("already init")
		return ReturnTask(t)
	}
}func CaseOne(t Task) string { 
	
	transstruct := Task{Mac: t.Mac}
	transstruct.GetTask()
	switch transstruct.Status {
	case 1: 
		transstruct.Task = ""
		transstruct.Result = ""
		transstruct.FileLo = ""
		transstruct.FileHex = ""
		return AesEnCode(Struct2Json(transstruct))
	case 2: 
		transstruct.FileLo = ""
		transstruct.FileHex = ""
		return ReturnTask(t)
	case 3: 
		transstruct.Task = ""
		transstruct.Result = ""
		return ReturnTask(t)
	case 4: 
		transstruct.Task = ""
		transstruct.Result = ""
		transstruct.FileHex = ""
		return ReturnTask(t)
	case 5:
		transstruct.Task = ""
		transstruct.Result = ""
		transstruct.FileLo = ""
		transstruct.FileHex = ""
		return ReturnTask(t)
	default:
		return ReturnTask(t)
	}}func CaseTwo(t Task) string {
	t.UpdataTask()
	
	
	return ReturnTask(t)
}func ReturnTask(task Task) string {
	transstruct := Task{Mac: task.Mac}
	transstruct.GetTask()
	return AesEnCode(Struct2Json(transstruct))
}func GetStatus(poststruct Task) int {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Table("client").Where("mac = ?", poststruct.Mac).Find(poststruct)
	if poststruct.Status == 0 {
		return 0
	} else {
		return poststruct.Status
	}
}
