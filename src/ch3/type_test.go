/*
## 数据类型
### 隐式类型
- 不支持任何隐式类型转换
- 不支持别名的隐式类型转换
### 指针
- 支持指针
- 不支持指针运算
### 字符串
- 值类型
- 初始值为空字符串而不是空
*/
package type_test

import (
	"fmt"
	"testing"
)

type MyInt int64

//测试隐式类型转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	fmt.Println(a, b)
	fmt.Println(c)
}

//测试指针
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	fmt.Println(a, aPtr)
	fmt.Printf("%T %T", a, aPtr)
}

//测试字符串
func TestString(t *testing.T) {
	var s string
	fmt.Println("*" + s + "*")
	fmt.Println(len(s))
	if s == "" {
		fmt.Println(s)
	}
}
