package modal
//import "fmt"

 
type Car struct {
	CarID 		string 		//购物车的id
	CartItms 	map[string]CartItmSlice	//购物车的所有购物项
	TotalCount 	int64		//图书的总数量，通过计算得到
	TotalAmount float64		//图书的总金额，通过计算得到
	UserID 		int64		//用户id
	
}


type CarS struct{
	Car 	*Car
	Judge   bool //
}

//GetMap 判断
func (car *Car)GetMap(c *CartItm) {
	if car.CartItms == nil{
		car.CartItms = make(map[string]CartItmSlice,1)
	}
	value,ok :=car.CartItms[c.Book.Shop.ShopName]
	if !ok {
		value = make(CartItmSlice,0,1)
	}
	value = append(value,c)
	car.CartItms[c.Book.Shop.ShopName] = value
	//fmt.Println(car.CartItms)
}
//获取图书总数量
func (car *Car)GetTotalCount() int64{	
	var totalCount int64
	//遍历购物车中的购物项
	for _,v :=range car.CartItms{
		for _,v1 :=range v{
			totalCount = totalCount + v1.Count
		}		
	}
	return totalCount
}
//获取图书总金额
func (car *Car)GetTotalAmount() float64{
	var TotalAmount float64
	for _,v :=range car.CartItms{
		for _,v1 := range v{
			TotalAmount = TotalAmount + v1.GetAmout()
		}		
	}
	return TotalAmount 
}