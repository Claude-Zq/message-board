package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dB *sql.DB

//连接上数据库
func InitDB() {

	db, err := sql.Open("mysql", "root:Zhouqing123456//@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//检查数据库是否可用可访问
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	dB = db
}
