package main

import "fmt"

func main()  {
	/*
	算术运算符： arithmetic operator
	+， -， * ,/ , % ， ++， --
	/,取商
	%，取余，取模
	 */
	 a := 9
	 b := 4
	 sum := a + b // 求和
	 sub := a - b //做差
	 mul := a * b // 乘法
	 div := a / b // 除
	 fmt.Println(a, "+", b, "=", sum)
	 fmt.Println(sub, mul, div)
	 mod := a % b
	 fmt.Println(mod)
	 /*
	 8/5 = 1
	 8%5 = 3
	 3/4 = 0
	 3%4 = 3
	 7/3 = 2
	 7%3 = 1
	 2/1 = 2
	 2%1 = 0
	  */

	  fmt.Println("a的数值：", a)
	  a++ // 给a自增1   a = a + 1
	  fmt.Println("a的数值：", a)
	  a-- // 给a自减1
	  fmt.Println(a)
}
