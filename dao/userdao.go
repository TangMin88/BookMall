package dao

import (
	"page/modal"
	"fmt"
)

//AddUser 向数据库中添加用户
func AddUser(username string, password string, email string, number string,address string) error {
	sql := "insert into users(username, password, email, number,address) values(?,?,?,?,?)"
	_, err := Db.Exec(sql,username, password, email, number,address)
	fmt.Println(sql)
	if err != nil {
		return err
	}
	return nil
}

//QueryUser 根据用户名向数据库中查询一条记录
func QueryUser(username string) *modal.User {
	sql := "select id, username, password, email, number, address from users where username = ?"
	row := Db.QueryRow(sql, username)
	user := &modal.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Number,&user.Address)
	return user
}

//UpdateUser 更新密码
func UpdateUser(username string, password string) error {
	sql := "update users set password = ? where username = ?"
	_, err := Db.Exec(sql, password, username)
	if err != nil {
		return err
	}
	return nil
}

//QueryUserOrderID 根据订单id查询用户信息
func QueryUserOrderID(orderid string) *modal.User{
	sql := "select username,number,address from users where id = (select userid from orders where id = ?)"
	row := Db.QueryRow(sql,orderid)
	user := &modal.User{}
	row.Scan(&user.Username, &user.Number,&user.Address)
	return user
}
