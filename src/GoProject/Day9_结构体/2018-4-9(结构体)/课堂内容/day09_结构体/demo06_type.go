package main

import "fmt"

// 起别名
type my_int int

type my_fun func(int, int) int

func test1(a, b int, fun my_fun)  {
	//....
}
/*
byte：uint8
rune：int32
 */

func main()  {
	a := 10
	fmt.Printf("%T\n", a)
	var b my_int = 20
	fmt.Printf("%T\n",b)
	a = b
	fmt.Println(a, b)
}
