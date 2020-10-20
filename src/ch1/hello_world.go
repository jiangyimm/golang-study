/*
- 必须是main包
- package main
- 必须是main方法
- func main
- 文件名不一定是main.go
*/

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
