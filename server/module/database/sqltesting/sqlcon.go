package mainimport (
	"database/sql"
	"fmt"
	"time"	_ "github.com/go-sql-driver/mysql"
)
const (
	USERNAME = "root"
	PASSWORD = "password"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "faygo"
)func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}	DB.SetConnMaxLifetime(100 * time.Second) 
	DB.SetMaxOpenConns(100)                  
		sqlstring := "INSERT INTO client (client_mac,client_status) VALUES ('testmac3',0)"
	CreateTable(DB, sqlstring)}
func CreateTable(DB *sql.DB, sqlstring string) {
	sql := sqlstring	if _, err := DB.Exec(sql); err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table successd")
}
