package modal


type CartItmSlice []*CartItm
//CartItm 购物车中的购物项
type CartItm struct {
	CartItmID 	int64 	//购物项的id
	Book 		*Book 	//购物项中的图书信息
	Count		int64 	//购物项中的图书数量
	Amout 		float64 //购物项中的金额小计，通过计算得到
	CarID 		string 	//当前购物项属于哪个购物车
}
//获取金额小计
func (cartItm *CartItm)GetAmout()float64{
	//获取当前购物项中图书的价格
	price := cartItm.Book.Price
	return float64(cartItm.Count) * price
}