package main

import "fmt"

func main()  {
	/*
	赋值运算符：=, +=,-=,*=,/=,%=,&=,|=,^=,<<=,>>=
	=，a = b,将=右侧的数值b，赋值给=左边的变量a
	+=, a += b,相当于：a = a + b
	%=, a %= b,相当于：a = a % b
	 */
	 var a int
	 a = 4 // 将数值4赋值给变量a
	 b := 3
	 fmt.Println(a)

	 a += 2 // a=6,b=3
	 a *= b // a=18,b=3
	 b %= 3 // a=18,b=0
	 //a &= 4
	 //a /= b // 除零，runtime error: integer divide by zero

	 fmt.Println(a, b)

}
