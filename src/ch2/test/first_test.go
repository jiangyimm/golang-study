/*
单元测试
MUST 文件名：以xxx_test.go
MUST 方法名：func Testxxx(t *testing.T)
*/
package try_test

import (
	"fmt"
	"testing"
)

func TestFirstTry(t *testing.T) {
	t.Log("My First Test")
	fmt.Println("My First Test")
}
