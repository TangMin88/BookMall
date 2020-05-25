package processorfunction

import(
	"net/http"
	"strconv"
	"page/dao"
	"text/template"
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
	//fmt.Println("Page", session)
	if flag {
		page.Judge= true
		page.Session = session	
		shop,b := ShopCookie(session.UserID)	
		if b {
			page.SJudge= true
			page.Shop = shop	
		}
	}
	
	t := template.Must(template.ParseFiles("../page.html"))
	t.Execute(w, page)
}

//Page 带价格查询的主页
func PricePage(w http.ResponseWriter, r *http.Request) {
	//获取当前页页码
	pageNo := r.FormValue("pageNo")
	if pageNo == ""{
		pageNo = "1"
	}
	//fmt.Println("Page", pageNo)
	//将页码转为int64
	ipageNo,_ := strconv.ParseInt(pageNo,10,64)
	//获取当前页图书
	pricemin := r.FormValue("min")
	pricemax := r.FormValue("max")
	if pricemin != "" && pricemax != ""{
		ipricemin,_ :=strconv.ParseFloat(pricemin,64)
		ipricemax,_ :=strconv.ParseFloat(pricemax,64)
		page,_ := dao.TotalBookPrice(ipageNo,ipricemin,ipricemax)
		//fmt.Println("Page", page.PageNo)
		//判断是否登录
		session, flag := LadingCookie(r)
		//fmt.Println("Page", session)
		if flag {
			page.Judge= true
			page.Session = session	
			shop,b := ShopCookie(session.UserID)	
			if b {
				page.SJudge= true
				page.Shop = shop	
			}
		}
		
		t := template.Must(template.ParseFiles("../htm/pricepage.html"))
		t.Execute(w, page)
	}else{
		Page(w,r)
	}
	
}

