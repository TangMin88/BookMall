package dao

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//Db 是mysql的句柄
	Db  *sql.DB
	err error
)

//init 获得链接数据库的句柄
func init() {
	Db, err = sql.Open("mysql", "用户名:密码@tcp(网络地址)/数据库?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}
}
