/*
常量快捷定义
① 自动递增常量
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)
② 按位自增常量
const (
	Readable = 1 << iota
	Writable
	Executable
)
*/
package constant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	fmt.Println(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111
	fmt.Println(Readable, Writable, Executable)
	fmt.Println(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
