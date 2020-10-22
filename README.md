# golang-study
golang code study records

#

## ch1 
### hello world project

- 必须是main包
- package main
- 必须是main方法
- func main
- 文件名不一定是main.go

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	//获取命令行参数
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}

	fmt.Println("hello world")

	//退出状态
	os.Exit(1)
}
```
#

## ch2 变量、常量
### 单元测试
- MUST 文件名：以xxx_test.go
- MUST 方法名：func Testxxx(t *testing.T)
### 斐波那锲数列
- 变量声明及赋值
```golang
var a int = 1
a := 1
```
- 多个变量同时赋值
```golang
a, b = b, a
```
### 常量快捷定义
- 自动递增常量
```golang
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)
```
- 按位自增常量
```golang
const (
	Readable = 1 << iota
	Writable
	Executable
)
```

## ch3 数据类型
### 隐式类型
- 不支持任何隐式类型转换
- 不支持别名的隐式类型转换
```golang
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	fmt.Println(a, b)
	fmt.Println(c)
}
```
### 指针
- 支持指针
- 不支持指针运算
```golang
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	fmt.Println(a, aPtr)
	fmt.Printf("%T %T", a, aPtr)
}
```
### 字符串
- 值类型
- 初始值为空字符串而不是空
```golang
func TestString(t *testing.T) {
	var s string
	fmt.Println("*" + s + "*")
	fmt.Println(len(s))
	if s == "" {
		fmt.Println(s)
	}
}
```