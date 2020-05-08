package dao

import "page/modal"

//AddSession 向数据库添加session
func AddSession(sess *modal.Session) error {
	sql := "insert into session values(?,?,?)"
	_, err := Db.Exec(sql, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 向数据库删除session
func DeleteSession(sessID string) error {
	sql := "delete from session where sessionID = ?"
	_, err := Db.Exec(sql, sessID)
	if err != nil {
		return err
	}
	return nil
}

//QuerySession 根据cookie的值查询session
func QuerySession(sessionID string) *modal.Session {
	sql := "select sessionID,user_name,user_ID from session where sessionID=?"
	row := Db.QueryRow(sql, sessionID)
	session := &modal.Session{}
	row.Scan(&session.SessionID, &session.UserName, &session.UserID)
	return session
}
