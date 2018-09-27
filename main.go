package main

import (
	"fmt"
)

func main() {
	//变量声明 var + 变量名+类型 变量声明必须要用
	var a int
	var c, d int
	fmt.Println(a)
	fmt.Println(c, d)

	a = 10
	fmt.Println(a)
	//变量初始化
	var e int = 3 //初始化
	fmt.Println(e)

	f := 30
	fmt.Println(f)
}
