package dao

import (
	"fmt"
	"page/modal"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("测试book中的函数")
	//t.Run("TotalBook", testTotalBook)
	//t.Run("tianjia",testAddBook)
	t.Run("Updatek",testUpdatek)
}
func testTotalBook(*testing.T) {
	page,_ := TotalBook(2)
	for _,v :=range page.BooK{
		fmt.Println(v)
	}
}	
func testAddBook(*testing.T) {
	shop := &modal.Shop{
		ID :2,
	}
	book :=&modal.Book{
	Title: "寄小读者",
	Author:"冰心",
	Price:25.3,
	Sales :145,
	Stock :100,
	ImgPath :"jingtai/regist/img/icon2.png",
	Shop:shop,
	}
	AddBook(book)
}

func testUpdatek(*testing.T) {
	img := "jingtai/书籍图片/默认图片.jpeg"
	Updatek(img)
}