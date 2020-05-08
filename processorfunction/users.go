package processorfunction

import (
	"fmt"
	"net/http"
	"page/dao"
	"page/modal"
	"text/template"
	"strconv"
)

//Page 主页
func Page(w http.ResponseWriter, r *http.Request) {
	//获取当前页页码
	pageNo := r.FormValue("pageNo")
	if pageNo == ""{
		pageNo = "1"
	}
	//fmt.Println("Page", pageNo)
	//将页码转为int64
	ipageNo,_ := strconv.ParseInt(pageNo,10,64)
	//获取当前页图书
	page,_ := dao.TotalBook(ipageNo)
	//fmt.Println("Page", page.PageNo)
	//判断是否登录
	session, flag := LadingCookie(r)
	//fmt.Println("Page", session, flag)
	if flag {
		page.Judge = true
		page.Session = session
		//判断是否有店铺
		shop,_ := dao.QueryShop(session.UserID)
		if shop != nil {
			page.JudgeShop =true
		}
	}	
	t := template.Must(template.ParseFiles("../page.html"))
	t.Execute(w, page)
}

//HomeLogin 去登录
func HomeLogin(w http.ResponseWriter, r *http.Request) {
	//判断cookie是否存在，避免重复建立cookie
	_, flag := LadingCookie(r)
	if flag {
		Page(w, r)
	} else {
		username := r.PostFormValue("userName")
		password := r.PostFormValue("password")
		user := dao.QueryUser(username)
		if user.ID > 0 && user.Password == password {
			//fmt.Println(username, password)
			//创建一个session后与cookie相关联
			cookie := AddSession(user)
			//fmt.Println("cookie2", cookie)
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("../htm/daozhuan.html"))
			t.Execute(w, username)
		} else {
			t := template.Must(template.ParseFiles("../htm/login.html"))
			t.Execute(w, "用户名或密码错误")
		}
	}
}

//HomeRegist 去注册
func HomeRegist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("userName")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	number := r.PostFormValue("number")
	address := r.PostFormValue("address")
	fmt.Println(password)
	dao.AddUser(username, password, email, number,address)
	t := template.Must(template.ParseFiles("../htm/login.html"))
	t.Execute(w, "注册成功，请重新登录")

}

//QueryName 通过Ajax请求验证用户名是否可用
func QueryName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("userName")
	user := dao.QueryUser(username)
	if user.ID > 0 {

		w.Write([]byte("用户名已存在"))
	} else {

		w.Write([]byte("<font styles='color:green'>用户名可用</font>"))
	}

}

//AcquireNumber 获取验证码
func AcquireNumber(w http.ResponseWriter, r *http.Request) {
	number := r.PostFormValue("number")
	//fmt.Println(number)
	code := modal.Verification(number)
	//fmt.Println(code)
	w.Write([]byte(code))
}

//Shouj 获取手机号
func Shouj(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("userName")
	user := dao.QueryUser(username)
	if user.ID > 0 {

		w.Write([]byte(user.Number))
	} else {

		w.Write([]byte("用户不存在"))
	}

}

//PassWordBack 更新密码
func PassWordBack(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("userName")
	password := r.PostFormValue("password")
	//更新用户
	dao.UpdateUser(username, password)
	//返回登录，
	t := template.Must(template.ParseFiles("../htm/login.html"))
	t.Execute(w, "密码更新成功，请重新登录")
}
