package modal

import (
	"time"
)

//Order 订单
type Order struct {
	OrderID 	string 		//订单号
	CreateTime 	time.Time	//生成订单的时间
	TotalCount 	int64		//一份订单中图书的总数量
	TotalAmout 	float64		//一份订单中的总金额
	State 		int64		//订单的状态 0 未发货，1 已发货， 2交易完成 3未付款 4 未确定
	UserID 		int64		//订单所属的用户
	OrderItem	[]*OrderItem
	Shop		*Shop		//订单所属商店
}

type Orders struct{
	Order	[]*Order
	Judge	bool//是否查到订单
}

//获取订单总数量
func (order *Order)GetOrderTotalCount() int64{	
	var totalCount int64
	//遍历购物车中的购物项
	for _,v :=range order.OrderItem{	
		totalCount = totalCount + v.Count			
	}
	return totalCount
}

//获取订单总金额
func (order *Order)GetOrderTotalAmount() float64{
	var TotalAmount float64
	for _,v :=range order.OrderItem{		
		TotalAmount = TotalAmount + v.Amount				
	}
	return TotalAmount 
}

//未发货
func (order *Order)NoSend()bool{
	return order.State == 0
}
//已发货
func (order *Order)SendComplate()bool{
	return order.State == 1
}
//未付款
func (order *Order)NotPaying()bool{
	return order.State == 3
}
//交易完成
func (order *Order)TheDeal()bool{
	return order.State == 2
}