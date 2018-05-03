package main

import "fmt"

func main()  {
	/*
	常量：同变量类似，区别在于程序执行过程中，数值不可以改变
	定义语法：
		const 常量名 数据类型 = 常量值/常量表达式
		const 常量名 = 常量值

	常量的数值：布尔类型，数值类型(整数，浮点，复数),字符串
	常量名：编码习惯：所有的字母都大写

	注意事项：
	1.常量定义后，可以不使用，也不会报错
	2.常量的值：布尔，数值，字符串
		常量表达式：常量的计算，内置函数：len(),cap()
	3.常量组：定义了一组常量
		没有直接定义常量值，默认和上一行一致。
		常量组中定义的常量值是一组相关的数据，可以作为枚举
	4.iota，特殊的常量，理解为计数器，
		定义const，iota的值从0开始。
		常量组中定义了常量就累加1.
	 */
	 // 1.定义常量
	const cPATH string  = "http://www.baidu.com"
	const PI  = 3.14
	fmt.Println(cPATH)
	fmt.Println(PI)

	//pi = 3.67 // 不能给常量重新赋值

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = LENGTH * WIDTH
	fmt.Println(area)

	//var a = 2
	//var b = 7
	//const area2  = LENGTH * WIDTH
	//const area2  = a * b
	//fmt.Println(area2)

	const c = "memeda"
	const con2  = len(c)

	// 同时定义多个常量
	const a2, b2, c2 = 1, false, "str" //多重赋值

	// 常量组
	const (
		a3 = 10
		a4 = false
		a5 = 3.14
	)
	//定义常量组的时候，如果没有指定类型和数值，默认和上一行相同
	const (
		x uint16 = 16 // 16
		y //16
		s = "abc" // "abc"
		z // "abc"
	)
	fmt.Println(x, y, s, z)

	// 用于枚举：一年四季，一周七天
	const (
		Sunday = 0
		Monday = 1
		Tuesday = 2
		Wednesday = 3
		Thursday = 4
		Friday = 5
		Saturday = 6
	)


}
