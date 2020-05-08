package dao

import (
	"fmt"
	"page/modal"
	"testing"
)

//func TestSession(t *testing.T) {
	//fmt.Println("测试session中的函数")
	//t.Run("AddSession", testAddSession)
	//t.Run("DeleteSession", testDeleteSession)
	//t.Run("QuerySession", testQuerySession)
//}
func testAddSession(*testing.T) {
	sess := &modal.Session{
		SessionID: "d73a645ce9c42b2941d088",
		UserName:  "kk",
		UserID:    6,
	}
	AddSession(sess)
}
func testDeleteSession(*testing.T) {
	a := "e8b05d73a645ce9c42b2941d088dbcca"
	DeleteSession(a)
}
func testQuerySession(*testing.T) {
	c := QuerySession("05d73a645ce9c42b2941d088")
	fmt.Println(c)
}
