package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//数据库连接信息
const (
	USERNAME = "root"
	PASSWORD = "1qaz!QAZ"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "faygo"
)

func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超时的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	//CreateTable(DB)

	sqlstring := "INSERT INTO client (client_mac,client_status) VALUES ('testmac3',0)"
	CreateTable(DB, sqlstring)

}

// CreateTable 向表中插入数据 通用方法 插入方法
func CreateTable(DB *sql.DB, sqlstring string) {
	sql := sqlstring

	if _, err := DB.Exec(sql); err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table successd")
}
