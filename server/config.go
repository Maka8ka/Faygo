package main
const (
	dsn          = "root:password@tcp(10.10.0.2:3306)/faygo?charset=utf8mb4&parseTime=True&loc=Local"
	AesKey       = "xxxxxxxxxxxxxx"
	defaultsleep = 10
	redirecturl  = "https:
)
type Client struct {
	Id         int    `gorm:"column:id" gorm:"default:NULL" json:"id" form:"id"" json:"id" form:"id"`
	Cname      string `gorm:"column:cname" gorm:"default:NULL" json:"cname" form:"cname"`
	Lan        string `gorm:"column:lan" gorm:"default:NULL" json:"lan" form:"lan"`
	Wan        string `gorm:"column:wan" gorm:"default:NULL" json:"wan" form:"wan"`
	Os         string `gorm:"column:os" gorm:"default:NULL" json:"os" form:"os"`
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
	Lan string `gorm:"column:lan" gorm:"default:NULL" json:"lan" form:"lan"`
	Os  string `gorm:"column:os" gorm:"default:NULL" json:"os" form:"os"`
	Mac string `gorm:"column:mac" gorm:"default:NULL" json:"mac" form:"mac"`
	
	Status  int    `gorm:"column:status" gorm:"default:NULL" json:"status" form:"status"`
	Sleep   int    `gorm:"column:sleep" gorm:"default:10" json:"sleep" form:"sleep"`
	Task    string `gorm:"column:task" gorm:"default:NULL" json:"task" form:"task"`
	Result  string `gorm:"column:result" gorm:"default:NULL" json:"result" form:"result"`
	FileHex string `gorm:"column:filehex" gorm:"default:NULL" json:"filehex" form:"filehex"`
	FileLo  string `gorm:"column:filelo" gorm:"default:NULL" json:"filelo" form:"filelo"`
}type TaskTime struct {
	Mac        string `gorm:"column:mac" gorm:"default:NULL" json:"mac" form:"mac"`
	LatestTime string `gorm:"column:latest_time" gorm:"default:NULL" json:"latest_time" form:"latest_time"`
}
