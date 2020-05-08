package dao

import (
	"page/modal"
	"fmt"
)


//AddBook 添加图书
func AddBook(book *modal.Book) error {
	sql := "insert into book(title,author,price,sales,stock,imgpath,shopid) values(?,?,?,?,?,?,?)"
	_,err := Db.Exec(sql,book.Title,book.Author,book.Price,book.Sales,book.Stock,book.ImgPath,book.Shop.ID)
	
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(sql)
	return nil
}

//QueryBook 根据书的id查询一本书
func QueryBook(bookid int64)(*modal.Book,error) {
	sql:= "select id,title,author,price,sales,stock,imgpath,shopid from book where id = ?"
	row := Db.QueryRow(sql, bookid)
	book := &modal.Book{}
	var shopid int64
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &shopid)
	if err != nil {
		return nil,err
	}
	shop,_ := QueryShopID(shopid)
	book.Shop = shop
	return book,nil
}

//QueryBookShopID 根据店铺的id查询店铺中的所有图书
func QueryBookShopID(shopID int64) ([]*modal.Book, error) {
	sql := "select id,title,author,price,sales,stock,imgpath from book where shopid = ?"
	rows,err := Db.Query(sql,shopID)
	if err != nil{
		return nil,err
	}
	var books []*modal.Book
	for rows.Next()	{
		book := &modal.Book{}
		err1 := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err1 != nil{
			return nil,err1
		}
		//将book放入books切片中
		books = append(books,book)
	}
	return books,nil
}

//TotalBook 获取带分页的图书
func TotalBook(pageNo int64) (*modal.Page,error){
	
	//获取图书的总记录数
	sql := "select count(*) from book"
	//接收总记录数
	var totalRecord int64
	row := Db.QueryRow(sql)
	row.Scan(&totalRecord)
	//设置每页显示的条数
	var pageSize int64 = 4
	//总页数
	var totalPage int64
	if totalRecord % pageSize == 0{
		totalPage = totalRecord / pageSize
	}else{
		totalPage = totalRecord / pageSize + 1
	}
	//获取当前页的图书
	sql2 := "select id,title,author,price,sales,stock,imgpath,shopid from book limit ?,?"
	rows,err := Db.Query(sql2, (pageNo - 1)*pageSize, pageSize)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}	
	var books []*modal.Book
	for rows.Next() {
		book := &modal.Book{}
		var shopid int64
		err1 := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &shopid)
		if err1 != nil{
			fmt.Println(err1)
			return nil,err1
		}
		fmt.Println(book.ImgPath)
		shop,_ := QueryShopID(shopid)
		book.Shop = shop
		//将book放入books切片中
		books = append(books,book)
	}
	//创建page
	page:=  &modal.Page{	
		BooK:		books, 	
		PageNo:		pageNo,
		PageSize:	pageSize,
		TotalPage:	totalPage,
		TotalRecord:totalRecord,
	}
	return page,nil 
}

//根据图书的id更新图书
func UpdateBook(book *modal.Book) error {
	sql := "update book set title=?, author=?, price=?, sales=?, stock=?, imgpath=? where id =?"
	_,err := Db.Exec(sql, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath, book.ID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBookid 根据图书的id删除对应的图书
func DeleteBookid(bookid int64) error {
	sql := "delete from book where id = ?"
	_,err := Db.Exec(sql, bookid)
	if err != nil{
		return err
	}
	return nil
}

//DeleteBookShopid 根据店铺的id删除所有属于店铺的图书
func DeleteBookShopid(shopid int64) error {
	sql := "delete from book where shopid = ?"
	_,err := Db.Exec(sql, shopid)
	if err != nil{
		return err
	}
	return nil
}

//属于店铺的所有图书
func ShopTotalBook(pageNo int64,shopid int64) (*modal.Page,error){
	//获取店铺图书的总记录数
	sql := "select count(*) from book where shopid=?"
	//接收总记录数
	var totalRecord int64
	row := Db.QueryRow(sql,shopid)
	row.Scan(&totalRecord)
	//设置每页显示的条数
	var pageSize int64 = 4
	//总页数
	var totalPage int64
	if totalRecord % pageSize == 0{
		totalPage = totalRecord / pageSize
	}else{
		totalPage = totalRecord / pageSize + 1
	}
	//获取当前页的图书
	sql2 := "select id,title,author,price,sales,stock,imgpath from book where shopid=? limit ?,?"
	rows,err := Db.Query(sql2,shopid, (pageNo - 1)*pageSize, pageSize)
	if err != nil{
		return nil,err
	}	
	var books []*modal.Book
	for rows.Next() {
		book := &modal.Book{}
		err1 := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err1 != nil{
			return nil,err1
		}
		//将book放入books切片中
		books = append(books,book)
	}
	shop,_ := QueryShopID(shopid)
	//创建page
	page:=  &modal.Page{	
		BooK:		books, 	
		PageNo:		pageNo,
		PageSize:	pageSize,
		TotalPage:	totalPage,
		TotalRecord:totalRecord,
		Shop:		shop,
	}
	return page,nil 
}
func Updatek(img string) error {
	sql := "update book set imgpath=?"
	_,err := Db.Exec(sql,img)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}