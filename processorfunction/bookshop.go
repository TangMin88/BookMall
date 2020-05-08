package processorfunction

import (
	//"fmt"
	"page/dao"
	"net/http"
	"strconv"
	"text/template"
	"page/modal"
)

//MyBookShop 我的店铺
func MyBookShop(w http.ResponseWriter, r *http.Request) {
	session,_ :=LadingCookie(r)
	//获取店铺信息
	shop,_ := dao.QueryShop(session.UserID)
	//获取当前页页码
	pageNo := r.FormValue("pageNo")
	if pageNo == ""{
		pageNo = "1"
	}
	//将页码转为int64
	ipageNo,_ := strconv.ParseInt(pageNo,10,64)
	//获取当前页图书
	page,_ := dao.ShopTotalBook(ipageNo,shop.ID)

	t := template.Must(template.ParseFiles("../htm/myshop.html"))
	t.Execute(w, page)						
}	

//MyInvoice 我的货单
func MyInvoice(w http.ResponseWriter, r *http.Request) {
	session,_ :=LadingCookie(r)
	//获取店铺信息
	shop,_ := dao.QueryShop(session.UserID)
	state := r.FormValue("state")
	var orders []*modal.Order
	if state == "" {
		//查询店铺的所有订单
		orders,_ = dao.QueryOrderShopID(shop.ID)
	}else{
		//查询不同状态的订单
		istate,_ := strconv.ParseInt(state,10,64)
		orders,_ = dao.QueryOrderStatusShopID(shop.ID,istate)
	}	
	orderss :=&modal.Orders{}
	if orders != nil{
		orderss.Order= orders
		orderss.Judge = true
	}
	t := template.Must(template.ParseFiles("../htm/myinvoice.html"))
	t.Execute(w,orderss)
}		

//TheDelivery 发货
func TheDelivery(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderid")
	state := r.FormValue("state")
	istate,_ := strconv.ParseInt(state,10,64)
	//更新订单状态
	dao.UpdateTheOrderState(orderid,istate)
	//返回我的订单页面
	MyInvoice(w,r)
}

//CheckTheDetails 查看详情
func CheckTheDetails(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderid")
	shop := r.FormValue("s")
	orderitem,_ :=dao.QueryOrderItem(orderid)
	user := dao.QueryUserOrderID(orderid)
	check := &modal.Check{
		OrderItem:orderitem,
		User:user,	
	}
	if shop != ""{
		check.Judge=true
	}
	t := template.Must(template.ParseFiles("../htm/checkThedetails.html"))
	t.Execute(w,check)
}