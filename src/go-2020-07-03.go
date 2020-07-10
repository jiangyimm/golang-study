package main

import (
	"fmt"
)

func main() {

	//有多个返回值，只想返回一个值,是通过下划线方式实现
	sum, _ := testCalc(10, 12)
	fmt.Println(sum)

	//测试管道
	testPipe()

	//天然高并发
	//在实现高并发的时候只需要在调用的函数前面加上go，就表示开启了并发
	/*
		for i := 0; i < 100; i++ {
			go testPrint(i)
		}
		time.Sleep(time.Second)
	*/

	/*
		sum := add(1, 2)
		fmt.Println(sum)
	*/
}

/*
定义一个方法
*/
func testAdd(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

/*
天然高并发
*/
func testPrint(a int) {
	fmt.Println(a)
}

/*
channel管道
*/
func testPipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	fmt.Println(len(pipe))
}

/*
多返回值
*/
func testCalc(a int, b int) (int, int) {
	sum := a + b
	avg := sum / 2
	return sum, avg
}
