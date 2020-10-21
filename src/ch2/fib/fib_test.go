/*
斐波那锲数列
① 变量声明及赋值
var a int = 1
a := 1
② 多个变量同时赋值
a, b = b, a
*/
package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	for i := 0; i < 100; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
