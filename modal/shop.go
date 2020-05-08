package modal

type Shop struct{
	ID 			int64
	ShopName 	string  	//店名
	Books		[]*Book		//属于店铺的图书切片
	UserID		int64 		//店铺所属用户id
}

