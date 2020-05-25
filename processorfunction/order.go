package processorfunction

import (
	//"fmt"
	"page/dao"
	"page/modal"
	"net/http"
	"time"
	"text/template"
	"strconv"
)
//ToCheckOut 去结账
func ToCheckOut(w http.ResponseWriter, r *http.Request) {
	session, f := LadingCookie(r)
	if f {
		//获取用户名
		userid := session.UserID
		//查询用户信息
		user := dao.QueryUser(session.UserName)
		//查询购物车
		car,_ := dao.GetCarUserID(userid)
		var orders []*modal.Order
		//购物车中一家一个订单编号
		for _,cartitm := range car.CartItms{
			//订单编号
			orderid := modal.UniqueID()
			var orderItems []*modal.OrderItem
			//保存订单项，遍历购物车中的map中的购物项
			shop := &modal.Shop{}
			for _,v := range cartitm{
				//fmt.Println(v)
				//创建订单项，将订单项添加的数据库中
				orderitem := &modal.OrderItem{
					Count:v.Count,
					Amount:v.Amout,
					Title:v.Book.Title,
					Author:v.Book.Author,
					Price:v.Book.Price,
					ImgPath:v.Book.ImgPath,
					OrderID:orderid,
					Bookid:v.Book.ID,
				}
				
				//更新当前图书的销量与库存
				book := v.Book
				shop = book.Shop
				book.Sales = book.Sales + v.Count
				book.Stock = book.Stock - v.Count
				dao.UpdateBook(book)
				orderItems= append(orderItems,orderitem)
				//创建一个订单项就删除对应的购物项
				//dao.DeleteCartItmByID(v.CartItmID)
			}	
			//创建订单
			order := &modal.Order{
				OrderID:orderid,
				CreateTime:time.Now(),
				OrderItem: orderItems,
				State:4,
				UserID:userid,
				Shop:shop,
			} 							
			order.TotalCount = order.GetOrderTotalCount()
			order.TotalAmout = order.GetOrderTotalAmount()
			//fmt.Println(order)
			dao.AddOrder(order)		
			orders = append(orders,order)
		}		
							
		confrim := &modal.Confrim{
			OrderSlice:orders,
			TotalAmounts:car.TotalAmount,
			User:user,
		}
		t := template.Must(template.ParseFiles("../htm/confirm.html"))
		t.Execute(w,confrim)
	}else{
		Login(w,r)
	}
	
}	

//CancellationOfOrder 确认订单页面返回购物车
func CancellationPayment(w http.ResponseWriter, r *http.Request) {
	session, f := LadingCookie(r)
	if f {
		//获取用户名
		userid := session.UserID
		//查询状态为4的订单
		orders,_ := dao.QueryOrderStatus(userid,4)
		if orders != nil{
			for _,v := range orders{
				orderID:=v.OrderID 
				//将图书的数量与销量更改回，删除未确定的订单在返回购物车
				ChangeBook(orderID)
				dao.DeteleOrder(orderID)
			}
		}
		GetCar(w,r)
	}else{
		Login(w,r)
	}
}

//变更图书
func ChangeBook(orderID string) {
	//订单取消后将图书的数量与销量更改回
	orderitem,_ :=dao.QueryOrderItem(orderID)
	for _,v := range orderitem{
		book,_ :=dao.QueryBook(v.Bookid)
		book.Sales = book.Sales - v.Count
		book.Stock = book.Stock + v.Count
		dao.UpdateBook(book)
	}
}

//QueryOrderStatus 我的订单，查询不同状态的订单
func QueryOrder(w http.ResponseWriter, r *http.Request) {
	session, f := LadingCookie(r)
	if f {
		//获取用户名
		userid := session.UserID
		//获取状态
		state := r.FormValue("state")
		var orders []*modal.Order
		if state == ""{
			//查询全部订单
			orders,_ = dao.QueryOrderUser(userid)
		}else{
			//查询不同状态的订单
			istate,_ := strconv.ParseInt(state,10,64)
			orders,_ = dao. QueryOrderStatus(userid,istate)
		}	
		orderss :=&modal.Orders{}
		if orders != nil{
			orderss.Order= orders
			orderss.Judge = true
		}
		t := template.Must(template.ParseFiles("../htm/myorder.html"))
		t.Execute(w,orderss)
	}else{
		Login(w,r)
	}
	
}

//UpdateTheOrder 确认订单页面，根据是否付款，更新订单
func UpdateTheOrder(w http.ResponseWriter, r *http.Request) {
	session, f := LadingCookie(r)
	if f {
		//获取用户名
		userid := session.UserID
		//获取状态
		state := r.FormValue("state")
		if state == "0" || state == "3"{	
			istate,_ := strconv.ParseInt(state,10,64)
			//获取数据库中未确定状态的订单
			orders,_ := dao.QueryOrderStatus(userid,4)
			//遍历订单切片，更新订单的状态
			for _,v := range orders{
				orderID:=v.OrderID 
				dao.UpdateTheOrderState(orderID,istate)
			}
			//找到购物车
			car,_ :=dao.GetCarUserID(userid)
			//情空购物车
			dao.DeleteCar(car.CarID)
		}
		//我的订单，查询不同状态的订单
		QueryOrder(w,r)
	}else{
		Login(w,r)
	}
	
}

//MyOrderState 我的订单，更新订单状态
func MyOrderState(w http.ResponseWriter, r *http.Request) {
	_, f := LadingCookie(r)
	if f {
		orderid := r.FormValue("orderid")
		state := r.FormValue("state")
		if orderid != "" && (state == "2" || state == "3"){
			istate,_ := strconv.ParseInt(state,10,64)
			//更新订单状态
			dao.UpdateTheOrderState(orderid,istate)
		}
		//返回我的订单页面
		QueryOrder(w,r)
	}else{
		Login(w,r)
	}	
}

//我的订单，取消订单
func CancellationOfOrder(w http.ResponseWriter, r *http.Request) {
	_, f := LadingCookie(r)
	if f {
		orderid := r.FormValue("orderid")
		//变更图书销量与数量
		ChangeBook(orderid)
		//删除订单
		dao.DeteleOrder(orderid)
		//返回我的订单页面
		QueryOrder(w,r)
	}else{
		Login(w,r)
	}	
}


	
