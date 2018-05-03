package main

import (
	"fmt"
	"os"
)

func main()  {
	/*
	defer函数中涉及到参数：
		函数的参数在defer调用时就已经传递，只不过被延迟执行而已。
	 */
	a := 1
	defer test1(a)
	a = 100
	fmt.Println(a)
	//test1(a)
	//fmt.Println(a)

	file,err:=os.Open("/test.txt")
}
func test1(num int)  {
	fmt.Println("函数中：num的值：", num)
}
