package mainimport (
	"fmt"	"encoding/json"	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type Client struct {
	Id         int    `gorm:"column:id" gorm:"default:NULL" json:"id" form:"id"" json:"id" form:"id"`
	Cname      string `gorm:"column:cname" gorm:"default:NULL" json:"cname" form:"cname"`
	Lan        string `gorm:"column:lan" gorm:"default:NULL" json:"lan" form:"lan"`
	Wan        string `gorm:"column:wan" gorm:"default:NULL" json:"wan" form:"wan"`
	OsType     string `gorm:"column:os" gorm:"default:NULL" json:"wan" form:"wan"`
	Mac        string `gorm:"column:mac" gorm:"default:NULL" json:"mac" form:"mac"`
	LatestTime string `gorm:"column:latest_time" gorm:"default:NULL" json:"latest_time" form:"latest_time"`
	Status     int    `gorm:"column:status" gorm:"default:0" json:"status" form:"status"`
	Sleep      int    `gorm:"column:sleep" gorm:"default:10" json:"sleep" form:"sleep"`
	Task       string `gorm:"column:task" gorm:"default:NULL" json:"task" form:"task"`
	Result     string `gorm:"column:result" gorm:"default:NULL" json:"result" form:"result"`
	FileHex    string `gorm:"column:filehex" gorm:"default:NULL" json:"filehex" form:"filehex"`
	FileLo     string `gorm:"column:filelo" gorm:"default:NULL" json:"filelo" form:"filelo"`
	Notes      string `gorm:"column:notes" gorm:"default:NULL" json:"notes" form:"notes"`
}
type Task struct {
	Mac        string `gorm:"column:mac" gorm:"default:NULL" json:"mac" form:"mac"`
	LatestTime string `gorm:"column:latest_time" gorm:"default:NULL" json:"latest_time" form:"latest_time"`
	Status     int    `gorm:"column:status" gorm:"default:NULL" json:"status" form:"status"`
	Task       string `gorm:"column:task" gorm:"default:NULL" json:"task" form:"task"`
	Sleep      int    `gorm:"column:sleep" gorm:"default:10" json:"sleep" form:"sleep"`
	Result     string `gorm:"column:result" gorm:"default:NULL" json:"result" form:"result"`
	FileHex    string `gorm:"column:filehex" gorm:"default:NULL" json:"filehex" form:"filehex"`
	FileLo     string `gorm:"column:filelo" gorm:"default:NULL" json:"filelo" form:"filelo"`
}func main() {
	
	dsn := "root:password@tcp(10.10.0.2:3306)/faygo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
		
	
		
	var clienttmp Client
	db.Table("client").Where("mac = ?", "aaa").Find(&clienttmp)
	json, _ := json.Marshal(clienttmp)
	fmt.Println(string(json))	
	db.Table("client").Model(&Task{}).Where("mac = ?", clienttmp.Mac).Update("result", "")
	}
