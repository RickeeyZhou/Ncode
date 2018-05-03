package main

import "fmt"

// 全局变量，不能使用:=
var m = 1000

func main()  {
	/*
	作用域：理解为可以使用的范围。
		全局变量和局部变量
		全局变量：所有的函数都可以使用，共享这一份数据。

		局部变量：函数内声明的变量，都叫局部变量
			局部变量的作用域就是定义该变量的范围内。{}
	 */
	 for i:=0;i<3;i++{// i的作用域就是for循环
	 	fmt.Println(i)
	 }
	 j := 100
	 fmt.Println(j)
	 fmt.Println(m)
	 fun7()
	 fun8()
	 fmt.Println(m)
}

func fun7()  {
	k := 200
	fmt.Println(k)
	m = 10
}
func fun8()  {
	fmt.Println(m)
}
