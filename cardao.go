package dao

import (
	"page/modal"
)

//AddCar 添加购物车
func AddCar(car *modal.Car) error{
	sql:= "insert into car(carid,totalcount,totalamount,userid) values(?,?,?,?)"
	_,err := Db.Exec(sql, car.CarID, car.GetTotalCount(), car.GetTotalAmount(), car.UserID)
	if err != nil {
		return err
	}
	//遍历购物车中的购物项
	for _,cartItm := range car.CartItms{
		for _,v := range cartItm {
			//将遍历的购物项一个个加入数据库中
			AddCartitm(v)
		}	
	}
	return nil
}

//GetCarUserID 根据用户的id查对应的购物车
func GetCarUserID(userid int64) (*modal.Car,error) {
	sql := "select carid, totalcount, totalamount,userid from car where userid = ?"
	row := Db.QueryRow(sql, userid)
	car := &modal.Car{}
	err := row.Scan(&car.CarID, &car.TotalCount, &car.TotalAmount, &car.UserID)
	if err != nil {
		return nil,err
	}
	cartItm := make(map[string]modal.CartItmSlice,8)
	car.CartItms = cartItm
	//获取购物车对应的购物项
	cartItms,_:= GetCartitmCarID(car.CarID)
	for _,v :=range cartItms{
		car.GetMap(v)
	}
	//car.CartItms = cartItms
	return car,nil
}
//UpdateCar 更新购物车
func UpdateCar(car *modal.Car) error {
	sql := "update car set totalcount = ?, totalamount = ? where carid =?"
	_,err := Db.Exec(sql, car.GetTotalCount(), car.GetTotalAmount(), car.CarID)
	if err != nil {
		return err
	}
	return nil 
}

//DeleteCar 根据购物车id删除购物车
func DeleteCar(carID string) error{
	//删除购物车前先删除对应的购物项
	err :=DeleteCartItm(carID)
	if err != nil{
		return err		
	}
	sql := "delete from car where carid = ?"
	_,err2 :=Db.Exec(sql,carID)
	if err2 != nil{
		return err2		
	}
	return nil
}
