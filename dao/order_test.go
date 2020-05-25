package dao

import (
	"testing"
	//"fmt"
	"page/modal"
	"time"
)

func TestOrder(t *testing.T) {
	//fmt.Println("Order中的函数")
	//t.Run("添加", testAddOrder)
}

func testAddOrder(t *testing.T) {
	shop := &modal.Shop{
		ID :1,
	}
	orderid := modal.UniqueID()
	order := &modal.Order{
	OrderID:orderid,
	CreateTime : time.Now(),
	TotalCount:	2,
	TotalAmout:50,
	State:0,
	UserID:3,
	Shop:shop,
	}
	//fmt.Println(order.CreateTime)
	AddOrder(order)
	orderitem := &modal.OrderItem{
	
	Count:	2,
	Amount:50,
	Title :	"三国演义",
	Author :"罗贯中",
	Price:25,
	ImgPath :"jingtai/regis/",
	OrderID	:orderid,
	}
	AddOrder(order)
	AddOrderItem(orderitem)
}
