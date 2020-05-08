package modal

type Confrim struct {
	OrderSlice 	[]*Order //所以订单的切片
	TotalAmounts float64  //付款时的总金额
	User		*User
}