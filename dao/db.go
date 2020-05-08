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
	Db, err = sql.Open("mysql", "db_write:tm19941015@tcp(47.56.89.150:3306)/test?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}
}
