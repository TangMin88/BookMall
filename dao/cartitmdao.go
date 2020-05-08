package dao

import (
	"page/modal"
)

//AddCartitm 添加购物项
func AddCartitm(cartitm *modal.CartItm) error {
	sql := "insert into cartitm(count, amount, carid, bookid) values(?,?,?,?) "
	_,err := Db.Exec(sql, cartitm.Count, cartitm.GetAmout(), cartitm.CarID, cartitm.Book.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartitmBookID 根据book的id与购物车的id查询对应的购物车是否有对应的购物项
func GetCartitmBookID(bookid int64, carid string) (*modal.CartItm,error){
	sql := "select cartitmid, count, amount, carid from cartitm where bookid = ? and carid = ? "
	row := Db.QueryRow(sql, bookid, carid)
	cartItm := &modal. CartItm{}
	err := row.Scan(&cartItm.CartItmID, &cartItm.Count, &cartItm.Amout, &cartItm.CarID)
	if err != nil {
		return nil,err
	}
	//根据在数据库中的bookID查到对应的book，加入到购物项中
	book,_ := QueryBook(bookid)
	cartItm.Book = book
	return cartItm,nil
}

//GetCartitmCarID 根据car的id查询购物车对应的所有购物项
func GetCartitmCarID(carid string) ([]*modal.CartItm, error){
	sql := "select cartitmid, count, amount, carid, bookid from cartitm where carid = ?"
	rows,err :=Db.Query(sql, carid)
	if err != nil{
		return nil,err
	}
	var cartItms []*modal.CartItm
	for rows.Next() {
		//接收数据库中查到的bookID
		var bookid int64
		cartItm := &modal.CartItm{}
		err1 := rows.Scan(&cartItm.CartItmID, &cartItm.Count, &cartItm.Amout, &cartItm.CarID, &bookid)
		if err1 != nil{
			return nil,err1
		}
		//根据在数据库中的bookID查到对应的book，加入到购物项中
		book,_ := QueryBook(bookid)
		cartItm.Book = book
		//将cartItm放入cartItms切片中
		cartItms = append(cartItms,cartItm)
	}	
	return cartItms,nil
}
//UpdateCartitm 更新购物项
func UpdateCartitm(cartitm *modal.CartItm)error {
	sql := "update cartitm set count = ?, amount = ? where cartitmid = ? "
	_,err := Db.Exec(sql, cartitm.Count, cartitm.GetAmout(), cartitm.CartItmID)
	if err != nil {
		return err
	}
	return nil 
}

//DeleteCartItm 根据购物车id删除对应的所有购物项
func DeleteCartItm(carID string) error{
	sql := "delete from cartitm where carid = ?"
	_,err :=Db.Exec(sql,carID)
	if err != nil{
		return err		
	}
	return nil
}

//DeleteIDCartItm 根据购物项id删除对应的购物项
func DeleteCartItmByID(cartitmID int64)error{
	sql := "delete from cartitm where cartitmid = ?"
	_,err :=Db.Exec(sql,cartitmID)
	if err != nil{
		return err		
	}
	return nil
}