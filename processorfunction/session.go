package processorfunction

import (
//"fmt"
	"net/http"
	"page/dao"
	"page/modal"
)

//AddSession 创建一个session后与cookie相关联
func AddSession(user *modal.User) http.Cookie {
	uuid := modal.UniqueID()
	sess := &modal.Session{
		SessionID: uuid,
		UserName:  user.Username,
		UserID:    user.ID,
	}
	//将Session保存到数据库
	dao.AddSession(sess)
	//创建Cookie,与session管联
	cookie := http.Cookie{
		Name:     "userse",
		Value:    uuid,
		HttpOnly: true,
	}
	return cookie
}

//DeleteCookie 注销cookie并删除对应的session，登出
func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("userse")
	if cookie != nil {
		uuid := cookie.Value
		//删除对应的session
		dao.DeleteSession(uuid)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	Page(w, r)
}

//LadingCookie 根据有没有cookie确定是不是登录状态
func LadingCookie(r *http.Request) (*modal.Session, bool) {
	cookie, _ := r.Cookie("userse")
	//fmt.Println("cookie", cookie)
	if cookie != nil {
		uuid := cookie.Value
		//fmt.Println("cookieval", uuid)
		session := dao.QuerySession(uuid)
		//fmt.Println("session", session)
		if session.UserID > 0 {
			return session, true
		}

	}
	return nil, false
}
