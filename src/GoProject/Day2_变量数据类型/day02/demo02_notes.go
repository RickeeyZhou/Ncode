package main

import "fmt"

//定义全局变量
var num6 = 100
var s2 string // ""
//num7 := 10// 该声明格式不能定义全局变量

func main()  {
	/*
	变量的注意事项：
	1.变量必须先创建才能使用，而且变量只能定义一次。
	2.go是强类型语言，变量的类型和赋值需要一致
	3.简短赋值：:=,左边的变量名至少有一个是新的。仅限于定义局部变量
	4.零值，也叫默认值。每种数据类型的默认值不同的。
	 */
	 var age int
	 age = 30
	 fmt.Println(age)

	 //age = "么么哒" // 错误类型数值
	 age, num := 1, 2
	 fmt.Println(age, num)
	 age2 ,num2 := 3, 4
	 fmt.Println(age2, num2)

	 var num3 int // 0
	 var s1 string //""
	 var num4 float64 // 0
	 var num5 bool // false
	 fmt.Println(num3, s1, num4, num5)

	_, b := 10, 20 // 空白标识符，用于舍弃数值。
	fmt.Println(b)





}
