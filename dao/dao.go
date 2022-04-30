package dao

import "database/sql"

var dB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:Zy08294641.@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	dB = db
}
