# 一个书籍买卖的Web服务端
 数据库：MySql。
# 功能:
登录，注册，短信验证，加入购物车，查看购物车，清空购物车，删除购物项，查看订单，订单详情，更新订单状态，我的店铺，成为店主，书籍增删改查，价格区间查询书籍等基本功能。
# dao目录
连接数据库的相关操作，条件查询、联表查询、区间查询、增加、更新、删除、外键等。
## db.go
MySQL的连接句柄
## bookdao.go
对书籍的相关操作
## cardao.go
对购物车的相关操作
## cartitmdao.go
对购物车中的购物项的相关操作
## orderdao.go
对订单的相关操作
## orderitemdao.go
对订单中订单项的相关操作
## sessiondao.go
对session的相关操作
## shopdao.go
对店铺的相关操作
## userdao.go
对用户的相关操作
# htm目录
程序用到的模板文件
# jingtai目录
静态资源与书籍图片
# modal目录
程序用到的结构体
# processorfunction目录
处理器函数
## bookshop.go
店铺相关的处理器函数，成为店主，我的店铺，我的货单，发货，查看详情，增、删、改书籍。
## car.go
购物车相关的处理器函数，添加图书到购物车，获取购物车信息，清空购物车，删除购物车中的购物项，根据购物项的数量更新购物车。
## order.go
订单相关的处理器函数，查询、更新、取消订单，结账（目前没有支付功能）。
## page.go
首页相关的处理器函数，带分页的书籍首页，价格区间查询的书籍首页。
## session.go
session相关的处理器函数,创建一个session后与cookie相关联，判断是否登录，判断是否有店铺，登出等操作。
## user.go
用户相关的处理器函数，登录，注册，判断用户名是否可用，获取手机号发送验证码，更新密码。

