package dao

import (
	"page/modal"
	"fmt"
	//"time"
)

//AddOrder添加订单
func AddOrder(order *modal.Order) error{
	sql :=  "insert into orders(id,totalcount,totalamount,state,userid,shopid,createtime) values(?,?,?,?,?,?,?)"
	_,err := Db.Exec(sql,order.OrderID, order.TotalCount, order.TotalAmout, order.State, order.UserID, order.Shop.ID, order.CreateTime)
	if err != nil {
		fmt.Println("AddOrder",err)
		return err
	}
	for _,v := range order.OrderItem{
		AddOrderItem(v)
	}	
	return nil
}

//DeteleOrder 根据订单id删除订单
func DeteleOrder(orderid string) error {
	DeleteOrderItem(orderid)
	sql := "delete from orders where id = ?"
	_,err := Db.Exec(sql,orderid)
	if err != nil {
		fmt.Println("DeteleOrder",err)
		return err
	}	
	return nil
}

//QueryOrder 根据用户id和状态查询订单
func QueryOrderStatus(userid int64,state int64)([]*modal.Order, error) {
	sql := "select id,totalcount,totalamount,state,userid,shopid,createtime from orders where userid=? and state=? order by createtime"
	rows,err := Db.Query(sql,userid,state)
	if err != nil{
		fmt.Println("QueryOrderStatus：err",err)
		return nil,err
	}
	var orders []*modal.Order
	for rows.Next() {
		var shopid int64
		order := &modal.Order{}
		err1 := rows.Scan(&order.OrderID, &order.TotalCount, &order.TotalAmout, &order.State, &order.UserID, &shopid, &order.CreateTime)
		if err1 != nil{
			fmt.Println("QueryOrderStatus：err1",err1)
			return nil,err1
		}
		shop,_ := QueryShopID(shopid)
		order.Shop = shop
		orders = append(orders,order)
	}
	return orders,nil
}

//QueryOrder 根据用户id查询全部订单
func QueryOrderUser(userid int64)([]*modal.Order, error) {
	sql := "select id,totalcount,totalamount,state,userid,shopid,createtime from orders where userid=? order by createtime"
	rows,err := Db.Query(sql,userid)
	if err != nil{
		fmt.Println("QueryOrderUser：err",err)
		return nil,err
	}
	var orders []*modal.Order
	for rows.Next() {
		var shopid int64
		order := &modal.Order{}
		err1 := rows.Scan(&order.OrderID, &order.TotalCount, &order.TotalAmout, &order.State, &order.UserID, &shopid, &order.CreateTime)
		if err1 != nil{
			fmt.Println("QueryOrderUser：err1",err1)
			return nil,err1
		}
		shop,_ := QueryShopID(shopid)
		order.Shop = shop
		orders = append(orders,order)
	}
	return orders,nil
}

//UpdateTheOrderState 根据订单id更新订单状态
func UpdateTheOrderState(orderid string,state int64) error {
	sql := "update orders set state = ? where id = ?"
	_,err := Db.Exec(sql,state,orderid)
	if err != nil {
		fmt.Println("UpdateTheOrderState",err)
		return err
	}	
	return nil
}

//QueryOrder 根据店铺id查询全部订单
func QueryOrderShopID(shopid int64)([]*modal.Order, error) {
	sql := "select id,totalcount,totalamount,state,userid,createtime from orders where shopid=? order by createtime"
	rows,err := Db.Query(sql,shopid)
	if err != nil{
		fmt.Println("QueryOrderShopID：err",err)
		return nil,err
	}
	var orders []*modal.Order
	for rows.Next() {
		order := &modal.Order{}
		err1 := rows.Scan(&order.OrderID, &order.TotalCount, &order.TotalAmout, &order.State, &order.UserID, &order.CreateTime)
		if err1 != nil{
			fmt.Println("QueryOrderShopID：err1",err1)
			return nil,err1
		}
		orders = append(orders,order)
	}
	return orders,nil
}

//QueryOrder 根据店铺id和状态查询订单
func QueryOrderStatusShopID(shopid int64,state int64)([]*modal.Order, error) {
	sql := "select id,totalcount,totalamount,state,userid,createtime from orders where shopid=? and state=? order by createtime"
	rows,err := Db.Query(sql,shopid,state)
	if err != nil{
		fmt.Println("QueryOrderStatusShopID：err",err)
		return nil,err
	}
	var orders []*modal.Order
	for rows.Next() {
		//var shopid int64
		order := &modal.Order{}
		err1 := rows.Scan(&order.OrderID, &order.TotalCount, &order.TotalAmout, &order.State, &order.UserID, &order.CreateTime)
		if err1 != nil{
			fmt.Println("QueryOrderStatusShopID：err1",err1)
			return nil,err1
		}
		//shop,_ := QueryShopID(shopid)
		//order.Shop = shop
		orders = append(orders,order)
	}
	return orders,nil
}

//QueryOrder 查询一份id
func QueryOrderByID(id string) (*modal.Order, error) {
	sql := "select totalcount,totalamount,state,createtime,shop.id,shopname from orders,shop where orders.userid == shop.userid and orders.id= ?"
	row := Db.QueryRow(sql, id)
	order := &modal.Order{}
	order.Shop = &modal.Shop{}
	err:= row.Scan(&order.TotalCount, &order.TotalAmout, &order.State, &order.CreateTime, &order.Shop.ID, &order.Shop.ShopName)
	if err != nil{
		return nil,err
	}
	return order,nil
}