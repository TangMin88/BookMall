package dao

import (
	"testing"
	"fmt"
	"page/modal"
)

func TestCar(t *testing.T) {
	//fmt.Println("car中的函数")
	//t.Run("AddCar", testAddCar)
	//t.Run("GetCarUserID", testGetCarUserID)
	//t.Run("DeleteCar", testDeleteCar)
}
//添加购物车
// func testAddCar(t *testing.T) {
// 	book,_ := QueryBook(1)
// 	var cartItms []*modal.CartItm
// 	cartitm := &modal.CartItm{
// 		Book:book,
// 		Count:2,
// 		//Amout:GetAmout(),
// 		CarID:"58585858",
// 	}
// 	cartItms = append(cartItms,cartitm)
// 	car := &modal.Car{
// 		CarID:"58585858",
// 		CartItm:cartItms,
// 		UserID:3,
// 	}
// 	AddCar(car)
// }

//添加购物车2
func testAddCar(t *testing.T) {
	book,_ := QueryBook(1)	
	cartitm := &modal.CartItm{
		Book:book,
		Count:2,
		//Amout:GetAmout(),
		CarID:"58585858",
	}
	car := &modal.Car{
		CarID:"58585858",
		//CartItms:c,
		UserID:3,
	}
	car.GetMap(cartitm)
	AddCar(car)
}

//根据用户的id查对应的购物车
func testGetCarUserID(t *testing.T) {
	car,_ := GetCarUserID(3)
	fmt.Println(car.CarID)
	for _,v:= range car.CartItms{
		for _,v1 := range v{
			fmt.Println(v1)
		}		
	}
	fmt.Println(car.TotalCount)
	fmt.Println(car.TotalAmount)
	fmt.Println(car.UserID)
}
//删除购物车
func testDeleteCar(t *testing.T) {
	DeleteCar("4798c4b58ae3bfb7ae8982308861a1fa")
}
