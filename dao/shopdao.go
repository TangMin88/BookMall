package dao

import (
	"page/modal"
	
)

//AddShop 添加店铺
func AddShop(shop *modal.Shop) error {
	sql := "insert into shop(shopname,userid) values(?,?)"
	_,err := Db.Exec(sql, shop.ShopName, shop.UserID)
	if err != nil{
		return err
	}
	return nil
}

//DeleteShop 删除店铺
func DeleteShop(shopid int64) error {
	//删除店铺时先删除店铺中的所有图书
	DeleteBookShopid(shopid)
	sql := "delete from shop where id = ?"
	_,err := Db.Exec(sql, shopid)
	if err != nil{
		return err
	}
	return nil
}

//QueryShop 查询店铺
func QueryShop(userid int64) (*modal.Shop,error){
	sql := "select id, shopname,userid from shop where userid = ?"
	row := Db.QueryRow(sql,userid)
	shop := &modal.Shop{}
	err := row.Scan(&shop.ID, &shop.ShopName, &shop.UserID)
	if err != nil{	
		return nil,err
	}
	return shop,nil
}

//UpdateShopName 更新店名
func UpdateShopName(shop *modal.Shop)error {
	sql := "update shop set shopname = ? where userid = ?"
	_,err := Db.Exec(sql, shop.ShopName, shop.UserID)
	if err != nil {
		return err
	}
	return nil
}

//QueryShopID 根据店铺id查询店铺
func QueryShopID(shopid int64) (*modal.Shop,error){
	sql := "select id, shopname,userid from shop where id = ?"
	row := Db.QueryRow(sql,shopid)
	shop := &modal.Shop{}
	err := row.Scan(&shop.ID, &shop.ShopName, &shop.UserID)
	if err != nil{
		return nil,err
	}
	return shop,nil
}