package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	//fmt.Println("测试function中的函数")
	//t.Run("AddUser", testAddUsert)
	//	t.Run("QueryUser", testQueryUser)
}

func testAddUsert(*testing.T) {
	AddUser("mippi", "123456", "825736335@qq.com", "15974030136","深圳宝安")
}

func testQueryUser(*testing.T) {
	user := QueryUser("mili")
	fmt.Println(user)
}
