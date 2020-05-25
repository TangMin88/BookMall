package dao

import (
	"fmt"
	"page/modal"
	"testing"
)

func TestSession(t *testing.T) {
	//fmt.Println("测试session中的函数")
	//t.Run("AddSession", testAddSession)
	//t.Run("DeleteSession", testDeleteSession)
	//t.Run("QuerySession", testQuerySession)
}
func testAddSession(*testing.T) {
	sess := &modal.Session{
		SessionID: "d73a645ce9c42b2941d088",
		UserName:  "mili",
		UserID:    3,
	}
	AddSession(sess)
}
func testDeleteSession(*testing.T) {
	a := "e8b05d73a645ce9c42b2941d088dbcca"
	DeleteSession(a)
}
func testQuerySession(*testing.T) {
	c := QuerySession("d73a645ce9c42b2941d0888")
	fmt.Println(c)
	
}
