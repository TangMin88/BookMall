package modal

//Home 发送到主页
type Page struct {
	Judge 		bool 	
	Session 	*Session
	BooK 		[]*Book 	//每页查询出的图书存放的切片
	PageNo 		int64 		//当前页
	PageSize 	int64 		//每页显示的条数
	TotalPage 	int64		//总页数，通过计算得到
	TotalRecord int64	    //总记录数，通过查询数据库得到
	SJudge 		bool
	Shop		*Shop
	Pricemin	float64
	Pricemax	float64
} 

//判断是否有上一页
func (p *Page)IsHasPrev()bool{
	return p.PageNo > 1
}

//判断是否有下一页
func (p *Page)IsHasNext()bool{
	return p.PageNo < p.TotalPage
}

//获取上一页
func (p *Page)PreviousPage()int64{
	return p.PageNo - 1
}

//获取下一页
func (p *Page)NextPage()int64{
	return p.PageNo + 1
}
