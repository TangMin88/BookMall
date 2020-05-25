package processorfunction

import (
	//"fmt"
	"page/dao"
	"text/template"
	"page/modal"
	"net/http"
	"strconv"
)

//AddBookCar 添加图书到购物车
func PostAddBookCar(w http.ResponseWriter, r *http.Request) {
	bookid := r.FormValue("bookId")
	//将bookid转为int64
	ibookid,_ := strconv.ParseInt(bookid,10,64)
	//先判断是否是登录状态
	session, flag := LadingCookie(r)
	if !flag {
		w.Write([]byte("请先登录"))	
		return
	}		
	//根据bookID查询book
	book,_ :=dao.QueryBook(ibookid)	
	//获取用户id
	userid := session.UserID
	//判断是否有对应的购物车
	car,_ := dao.GetCarUserID(userid)
	if car == nil{	
		//创建一个购物车
		carid := modal.UniqueID()
		car = &modal.Car{
			CarID:carid,
			UserID:userid,
		}
		//创建购物项		
		cartItm :=&modal.CartItm{
			CarID:carid,
			Book:book,
			Count:1,
		}
		//将购物项加入到购物车map中
		car.GetMap(cartItm)	
		//将购物车保存到数据库
		dao.AddCar(car)
	}else {			
		//根据bookid查询是否有对应的购物项
		cartItm,_ := dao.GetCartitmBookID(ibookid, car.CarID)
		if cartItm == nil {
			//创建一个购物项
			cartItm = &modal.CartItm{
				CarID:car.CarID,
			 	Book:book,
			 	Count:1,
			}
			car.GetMap(cartItm)
			//将购物项加入到数据库中	
			dao.AddCartitm(cartItm)												
		}else {							
		breaktag:
		for _,val := range car.CartItms{
			for _,v := range val {
				if v.Book.ID == ibookid{
					v.Count = v.Count + 1
					//更新数据库
					dao.UpdateCartitm(v)
					break breaktag
				}				
				}
			}
		}	
		//不管是否有对应的购物项，都需要更新购物车
		dao.UpdateCar(car)
	}
	w.Write([]byte(book.Title))
}

//GetCar 获取购物车信息
func GetCar(w http.ResponseWriter,r *http.Request) {
	session,f := LadingCookie(r)
	if f {
		//获取用户名
		userid := session.UserID
		//通过用户名查购物车
		car,_ := dao.GetCarUserID(userid)
		//当购物项的数量大于0，或者购物车不为空
		cars := &modal.CarS{}
		if car != nil && len(car.CartItms)>0{	
			cars.Judge = true
			cars.Car = car
				
		}
		t := template.Must(template.ParseFiles("../htm/car.html"))
		t.Execute(w,cars)	
	}else{
		Login(w,r)
	}
	
}

//DeleteIDCar清空购物车
func DeleteIDCar(w http.ResponseWriter, r *http.Request){
	_,f := LadingCookie(r)
	if f {
		carid := r.FormValue("carid")
		if carid != ""{
			//删除购物车
			dao.DeleteCar(carid)
		}
		//再次查询购物车信息
		GetCar(w,r)
	}else{
		Login(w,r)
	}
}

//DeleteIDCartItm删除购物车中的购物项
func DeleteIDCartItm(w http.ResponseWriter, r *http.Request){
	session,f := LadingCookie(r)
	if f {
		//获取购物项id
		cartitmid := r.FormValue("cartitmid")
		if cartitmid != "" {
			//将字符串转成int64
			icartitmid,_ := strconv.ParseInt(cartitmid,10,64)
			//获取用户id，找到对应的购物车
			userid := session.UserID
			car,_ := dao.GetCarUserID(userid)
			//遍历购物车中的购物项，找到要删除的购物项，删除后更新购物车
			//cartItm := car.CartItms
			for _,v :=range car.CartItms{
				for k,v1 := range v{
					if v1.CartItmID == icartitmid{
						shopname := v1.Book.Shop.ShopName
						//删除购物项
						dao.DeleteCartItmByID(icartitmid)
						//如果购物车中对应的购物项只有一个，就将对应的键从map中删除，		
						if len(v) == 1{
							delete(car.CartItms,shopname)
						}else{
							//将购物项从购物车切片移除
							v = append(v[:k],v[k+1:]...)
							//将删除购物项后的切片再次赋给购物车中的切片
							car.CartItms[shopname] = v
						}																						
						
					//更新购物车
					dao.UpdateCar(car)
					}		
				}	
				//更新购物车
				//dao.UpdateCar(car)
			}
		}
		//再次查询购物车信息
		GetCar(w,r)
	}else{
		Login(w,r)
	}
}
//
//UpdateCartItmID 根据购物项的数量更新购物车
func UpdateCartItmID(w http.ResponseWriter, r *http.Request){
	session,f := LadingCookie(r)
	if f {
		//购物项id
		cartitmid := r.FormValue("cartitmid")
		//图书数量
		bookcount := r.FormValue("bookcount")
		if cartitmid != "" && bookcount != "" {
			icartitmid,_ := strconv.ParseInt(cartitmid,10,64)
			ibookcount,_ := strconv.ParseInt(bookcount,10,64)
			//获取用户id，找到对应的购物车
			userid := session.UserID
			car,_ := dao.GetCarUserID(userid)
			//遍历购物车中的购物项，找到要更新的购物项，修改数量后更新购物车
			cartItms := car.CartItms
			breaktag:
			for _,v :=range cartItms{
				for _,v1 := range v{
					if v1.CartItmID == icartitmid && v1.Book.Stock > ibookcount{
						v1.Count = ibookcount
						//更新购物项
						dao.UpdateCartitm(v1)
						break breaktag
					}				
				}
			}	
			//更新购物车
			dao.UpdateCar(car)	
		}
		//再次查询购物车信息
		GetCar(w,r)
	}else{
		Login(w,r)
	}
}