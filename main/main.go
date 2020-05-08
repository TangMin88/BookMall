package main

import (
	"net/http"
	"page/processorfunction"
	//"text/template"
	"log"
)

func main() {
	//处理静态资源，如css和js文件
	http.Handle("/jingtai/", http.StripPrefix("/jingtai/", http.FileServer(http.Dir("../jingtai"))))
	http.Handle("/htm/", http.StripPrefix("/htm/", http.FileServer(http.Dir("../htm"))))
	//http.Handle("/regist/", http.StripPrefix("/regist/", http.FileServer(http.Dir("../html/regist/regist"))))

	//将去主页的函数转换为http处理器函数
	http.HandleFunc("/page", processorfunction.Page)
	//去注册
	http.HandleFunc("/regist", processorfunction.HomeRegist)
	//去登录
	http.HandleFunc("/Login", processorfunction.HomeLogin)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/queryusername", processorfunction.QueryName)
	//登出
	http.HandleFunc("/deletecookie", processorfunction.DeleteCookie)
	//获取验证码
	http.HandleFunc("/acquirenumber", processorfunction.AcquireNumber)
	//获取手机号
	http.HandleFunc("/shouj", processorfunction.Shouj)
	//更新密码
	http.HandleFunc("/passwordback", processorfunction.PassWordBack)
	//添加图书到购物车
	http.HandleFunc("/addBookCar", processorfunction.AddBookCar)
	//获取购物车信息
	http.HandleFunc("/getCar", processorfunction.GetCar)
	//清空购物车
	http.HandleFunc("/deleteIDCar", processorfunction.DeleteIDCar)
	//删除购物车中的购物项
	http.HandleFunc("/deleteIDCartItm", processorfunction.DeleteIDCartItm)
	//更新购物车中购物项的数量
	http.HandleFunc("/updateCartItmID", processorfunction.UpdateCartItmID)
	//去结账
	http.HandleFunc("/tocheckout", processorfunction.ToCheckOut)
	//确认订单页面返回购物车
	http.HandleFunc("/cancellationpayment",processorfunction.CancellationPayment)
	//我的订单，查询不同状态的订单
	http.HandleFunc("/queryorder",processorfunction.QueryOrder)
	//确认订单页面，根据是否付款，更新订单
	http.HandleFunc("/updatetheorder",processorfunction.UpdateTheOrder)
	//我的订单，更新订单状态
	http.HandleFunc("/myorderstate",processorfunction.MyOrderState)
	//我的订单，取消订单
	http.HandleFunc("/cancellationoforder",processorfunction.CancellationOfOrder)
	//我的店铺
	http.HandleFunc("/mybookshop",processorfunction.MyBookShop)
	//我的货单
	http.HandleFunc("/myinvoicep",processorfunction.MyInvoice)
	//发货
	http.HandleFunc("/thedelivery",processorfunction.TheDelivery)
	//查看详情
	http.HandleFunc("/checkthedetails",processorfunction.CheckTheDetails)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
