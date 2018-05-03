package main

import "fmt"

func main() {
	a := 100
	p1 := &a
	fmt.Println(a, *p1)
	*p1++ // 操作指针更改变量的数值
	fmt.Println(a, *p1)
	*p1 = 10000
}
