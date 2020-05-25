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
	//带价格
	http.HandleFunc("/pricepage", processorfunction.PricePage)
	//去注册页面
	http.HandleFunc("/regist", processorfunction.Regist)
	//去登录页面
	http.HandleFunc("/login", processorfunction.Login)
	//登出
	http.HandleFunc("/deletecookie", processorfunction.DeleteCookie)
	//注册表单提交
	http.HandleFunc("/postregist", processorfunction.PostRegist)
	//登录表单提交
	http.HandleFunc("/postlogin", processorfunction.PostLogin)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/queryusername", processorfunction.PostQueryName)
	//获取验证码
	http.HandleFunc("/acquirenumber", processorfunction.PostAcquireNumber)
	//获取手机号
	http.HandleFunc("/shouj", processorfunction.PostShouj)
	//更新密码
	http.HandleFunc("/passwordback", processorfunction.PostPassWordBack)
	
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
	//添加图书到购物车
	http.HandleFunc("/postaddBookCar", processorfunction.PostAddBookCar)
	
	//我的订单，查询不同状态的订单
	http.HandleFunc("/queryorder",processorfunction.QueryOrder)
	//确认订单页面，根据是否付款，更新订单
	http.HandleFunc("/updatetheorder",processorfunction.UpdateTheOrder)
	//我的订单，更新订单状态
	http.HandleFunc("/myorderstate",processorfunction.MyOrderState)
	//我的订单，取消订单
	http.HandleFunc("/cancellationoforder",processorfunction.CancellationOfOrder)
	//查看详情
	http.HandleFunc("/checkthedetails",processorfunction.CheckTheDetails)
	
	//成为店主
	http.HandleFunc("/owner",processorfunction.AsTheOwner)
	//我的店铺
	http.HandleFunc("/mybookshop",processorfunction.MyBookShop)
	//我的货单
	http.HandleFunc("/myinvoicep",processorfunction.MyInvoice)
	//发货
	http.HandleFunc("/thedelivery",processorfunction.TheDelivery)
	//提交成为店主的表单
	http.HandleFunc("/owner/postowner",processorfunction.PostAsTheOwner)
	//删除书籍
	http.HandleFunc("/dleteshopBook",processorfunction.DleteShopBook)
	//添加书籍
	http.HandleFunc("/addshopbook",processorfunction.AddShopBook)
	//提交添加书籍表单
	http.HandleFunc("/postaddshopbook",processorfunction.PostAddShopBook)
	//提交修改书籍表单
	http.HandleFunc("/postupdateShopBook",processorfunction.PostUpdateShopBook)
	
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
