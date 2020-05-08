package modal

type Book struct {
	ID 		int64
	Title 	string 	//书名
	Author 	string 	//作者
	Price 	float64	//价格
	Sales 	int64	//销量
	Stock 	int64	//库存
	ImgPath string 	//图片路径
	Shop   *Shop  //图书所属店铺id
}

