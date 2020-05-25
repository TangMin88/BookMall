package processorfunction

import (
	"fmt"
	"page/dao"
	"net/http"
	"strconv"
	"text/template"
	"page/modal"
	//"io"
	"os"
	"io/ioutil"
	
)

//AsTheOwner 成为店主
func AsTheOwner(w http.ResponseWriter, r *http.Request) {
	_,f :=LadingCookie(r)
	if f {
		t := template.Must(template.ParseFiles("../htm/owner.html"))
		t.Execute(w, "")
	}else{
		Login(w,r)
	}
}
//PostAsTheOwner 提交成为店主表单
func PostAsTheOwner(w http.ResponseWriter, r *http.Request) {
	session,_ :=LadingCookie(r)
	shopName := r.PostFormValue("shopName")
	//查询用户是否有店铺，避免重复
	shop:= &modal.Shop{}
	_,e :=dao.QueryShop(session.UserID)
	if e != nil {
		shop.ShopName = shopName
		shop.UserID = session.UserID
		dao.AddShop(shop)
	}
	MyBookShop(w,r)
}

//MyBookShop 我的店铺
func MyBookShop(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	session,f :=LadingCookie(r)
	if f {
		//判断是否有店铺
		shop,b := ShopCookie(session.UserID)
		if b {
			//获取当前页页码
			pageNo := r.FormValue("pageNo")
			if pageNo == ""{
				pageNo = "1"
			}
			//将页码转为int64
			ipageNo,_ := strconv.ParseInt(pageNo,10,64)
			//获取当前页图书
			page,_ := dao.ShopTotalBook(ipageNo,shop.ID)
			page.Shop = shop

			t := template.Must(template.ParseFiles("../htm/myshop.html"))
			t.Execute(w, page)						
		}else{
 			AsTheOwner(w,r)
		}
	}else{
		Login(w,r)
	}
	
}	

//MyInvoice 我的货单
func MyInvoice(w http.ResponseWriter, r *http.Request) {
	session,f :=LadingCookie(r)
	if f {
		//判断是否有店铺
		shop,b := ShopCookie(session.UserID)
		if b {
			state := r.FormValue("state")
			var orders []*modal.Order
			if state == "" {
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
		}else{
 			AsTheOwner(w,r)
		}
	}else{
		Login(w,r)
	}
	
}		

//TheDelivery 发货
func TheDelivery(w http.ResponseWriter, r *http.Request) {
	session,f :=LadingCookie(r)
	if f {
		//判断是否有店铺
		shop,b := ShopCookie(session.UserID)
		if b {
			orderid := r.FormValue("orderid")
			state := r.FormValue("state")
			if orderid !=""&& state =="1"{
				//查询订单id是否属于登录的店铺
				order,_ :=dao.QueryOrderByID(orderid)
				if order != nil && order.Shop.ID ==shop.ID {
					istate,_ := strconv.ParseInt(state,10,64)
					//更新订单状态
					dao.UpdateTheOrderState(orderid,istate)
				}
			}		
			//返回我的订单页面
			MyInvoice(w,r)
		}else{
 			AsTheOwner(w,r)
		}
	}else{
		Login(w,r)
	}
	
}

//CheckTheDetails 查看详情
func CheckTheDetails(w http.ResponseWriter, r *http.Request) {
	_,f :=LadingCookie(r)
	if f {
		orderid := r.FormValue("orderid")
		shop := r.FormValue("s")

		if orderid != ""{
			//查询订单的订单项
			orderitem,_ :=dao.QueryOrderItem(orderid)
			if orderitem != nil{
				//查询用户信息
				user := dao.QueryUserOrderID(orderid)
				check := &modal.Check{
					OrderItem:orderitem,
					User:user,	
				}
				if shop == "shop"{
					check.Judge=true
				}
				t := template.Must(template.ParseFiles("../htm/checkThedetails.html"))
				t.Execute(w,check)
			}else{
				//返回我的订单页面
				MyInvoice(w,r)
			}		
		}else{
			//返回我的订单页面
			MyInvoice(w,r)
		}
	}else{
		Login(w,r)
	}
}

//DleteShopBook 删除书籍
func DleteShopBook(w http.ResponseWriter, r *http.Request) {
	session,f :=LadingCookie(r)
	if f {
		//判断是否有店铺
		shop,b := ShopCookie(session.UserID)
		if b {
			id := r.FormValue("id")
			iid,_ := strconv.ParseInt(id,10,64)
			//查询要删除的书籍是否属于登录的书店
			book,_:=dao.QueryBook(iid)
			if book != nil && book.Shop.ID == shop.ID{
				dao.DeleteBookid(iid)
			}
			MyBookShop(w,r)
			}else{
				AsTheOwner(w,r)
		   }
	}else{
		Login(w,r)
	}
}

//添加/修改书籍
func AddShopBook(w http.ResponseWriter, r *http.Request) {
	session,f :=LadingCookie(r)
	if f {
		//判断是否有店铺
		shop,b := ShopCookie(session.UserID)
		if b {
			var j bool
			bookid := r.FormValue("bookid")
			if bookid != ""{
				//修改
				ibookid,_ := strconv.ParseInt(bookid,10,64)
				book,_ := dao.QueryBook(ibookid)
				if book != nil && book.Shop.ID == shop.ID {
					j = true
					t := template.Must(template.ParseFiles("../htm/updatebook.html"))
					t.Execute(w,map[string]interface{}{
						"j":j,
						"book":book,
					})
				}else{
					AsTheOwner(w,r)
				}
			}else{
				//添加
				t := template.Must(template.ParseFiles("../htm/updatebook.html"))
				t.Execute(w,map[string]interface{}{
					"j":j,
					"shopid":shop.ID,

				})
			}
			
		}else{
			AsTheOwner(w,r)
		}
	}else{
		Login(w,r)
	}
}
//添加书籍表单
func PostAddShopBook(w http.ResponseWriter, r *http.Request) {
	file,fileheader,err := r.FormFile("f")
	var filename string
	if err != nil{
		filename ="默认图片.jpeg"
	}else{
		filename =fileheader.Filename
		fmt.Println(filename)
		defer file.Close()//文件流句柄
		
		osfile, err2 := os.Create("../jingtai/书籍图片/"+filename)
		fmt.Println(osfile.Name())
		if err2 != nil{
			fmt.Println("err2",err2)
		}
		defer osfile.Close()//关闭临时文件句柄
		//从file读取文件，返回[]byte
		fileBytes, err3 := ioutil.ReadAll(file)
		if err3 != nil{
			fmt.Println("err3",err3)
		}
		//
		//io.Copy(osfile, file)
		osfile.Write(fileBytes)
	}
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	iprice,_ := strconv.ParseFloat(price,64)
	stock := r.PostFormValue("stock")
	istock,_ := strconv.ParseInt(stock,10,64)
	shopid := r.FormValue("shopid")
	ishopid,_ := strconv.ParseInt(shopid,10,64)
	book := &modal.Book{
		Title:title,
		Author:author,
		Price:iprice,
		Stock:istock,
		ImgPath:"jingtai/书籍图片/"+filename, 
		Sales:0,
	}
	book.Shop=&modal.Shop{
		ID:ishopid,
	}
	dao.AddBook(book)
	MyBookShop(w,r)
}

//修改
func PostUpdateShopBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	iprice,_ := strconv.ParseFloat(price,64)
	stock := r.PostFormValue("stock")
	istock,_ := strconv.ParseInt(stock,10,64)
	bookid := r.FormValue("bookid")
	ibookid,_ := strconv.ParseInt(bookid,10,64)
	book := &modal.Book{
		ID :ibookid,
		Title:title,
		Author:author,
		Price:iprice,
		Stock:istock,
		Sales:0,
	}
	
	file,fileheader,err := r.FormFile("f")
	var filename string
	if err == nil{
		filename =fileheader.Filename
		fmt.Println(filename)
		defer file.Close()//文件流句柄
		osfile, err2 := os.Create("../jingtai/书籍图片/"+filename)
		if err2 != nil{
			fmt.Println("err2",err2)
		}
		defer osfile.Close()//关闭临时文件句柄
		//从file读取文件，返回[]byte
		fileBytes, err3 := ioutil.ReadAll(file)
		if err3 != nil{
	 	  fmt.Println("err3",err3)
		}
		//io.Copy(osfile, file)
		osfile.Write(fileBytes)
		book.ImgPath= "jingtai/书籍图片/"+filename
	}else{
		f1:= r.FormValue("f1")
		book.ImgPath = f1
		
	}
	
	dao.UpdateBook(book)
	MyBookShop(w,r)
}
