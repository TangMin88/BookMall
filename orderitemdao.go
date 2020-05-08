package dao

import (
	"page/modal"
	"fmt"
)

//AddOrderItem添加订单项
func AddOrderItem(orderitem *modal.OrderItem) error {
	sql :="insert into orderitem(count,amount,title,author,price,imgpath,orderid,bookid) values(?,?,?,?,?,?,?,?)"
	_,err :=Db.Exec(sql, orderitem.Count, orderitem.Amount, orderitem.Title, orderitem.Author, orderitem.Price, orderitem.ImgPath, orderitem.OrderID, orderitem.Bookid)
	if err != nil {
		fmt.Println("OrderItem",err)
		return err
	}
	return nil
}

//根据订单id删除订单项
func DeleteOrderItem(OrderID string) error{
	sql := "delete from orderitem where orderid = ?"
	_,err := Db.Exec(sql,OrderID)
	if err != nil {
		fmt.Println("DeteleOrder",err)
		return err
	}	
	return nil
}

//根据订单号查询订单项
func QueryOrderItem(OrderID string)([]*modal.OrderItem ,error){
	sql := "select count,amount,title,author,price,imgpath,orderid,bookid from orderitem where orderid=?"
	rows,err := Db.Query(sql,OrderID)
	if err != nil{
		fmt.Println("QueryOrderItem：err",err)
		return nil,err
	}
	var orderitems []*modal.OrderItem
	for rows.Next() {
		orderitem := &modal.OrderItem{}
		err1 := rows.Scan(&orderitem.Count, &orderitem.Amount, &orderitem.Title, &orderitem.Author, &orderitem.Price, &orderitem.ImgPath, &orderitem.OrderID, &orderitem.Bookid)
		if err1 != nil{
			fmt.Println("QueryOrderItem：err1",err1)
			return nil,err1
		}		
		orderitems = append(orderitems,orderitem)
	}
	return orderitems,nil
}