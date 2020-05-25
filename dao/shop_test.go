package dao

import (
	"testing"
	"fmt"
	"page/modal"
)

func TestShop(t *testing.T) {
	//fmt.Println("shop中的函数")
	//t.Run("添加书店", testAddShop)
	//t.Run("查询书店", testQueryShop)
}

func testAddShop(t *testing.T) {
	shop := &modal.Shop{
		ShopName:"小小书店",
		UserID:11,
	}
	AddShop(shop)
}

func testQueryShop(t *testing.T) {
	shop,_ := QueryShop(11)
	fmt.Println(shop)
}



