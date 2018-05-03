package main

import "fmt"

func main()  {
	/*
	变量：
		概念：本质就是一小块内存，专门用于存储数据。在程序运行时，可以被改变。
		语法：
			1种：var 变量名 数据类型
				 变量名 = 数值

			2种：var 变量名 数据类型 = 数值

			3种：类型推断，omitted省略类型
			4种：省略var ，变量名 := 数值，限于函数内使用。
	 */
	 //1. 定义(声明，申请)变量
	 var age int // 定义一个变量，变量名叫age，存储int类型的整数
	 age = 100
	 fmt.Println("age的数值是：", age)// 100

	 // 更改变量的数值
	 age = 200
	 fmt.Println("更改age的数值：", age)// 200

	 // 2.定义一个变量
	 var age2 int = 30
	 fmt.Println("age2的数值：", age2)

	 // 3. 类型推断，可以省略类型
	 var num = 200
	 fmt.Printf("num的数值是：%d，类型是：%T\n", num, num)

	 // 4.省略var
	 num2 := 1000
	 fmt.Println("num2的数值是：", num2)

	 // 5，同时定义多个变量：
	 //类型一致
	 var a, b, c int
	 a = 3
	 b = 4
	 c = 5
	 fmt.Println(a, b, c)

	 var d, e, f int = 6, 7, 8
	 fmt.Println(d, e, f)

	 //类型不一致
	 var h ,i = 100, "memeda"
	 fmt.Println(h,i)
	 name, age3 := "王二狗", 28
	 fmt.Println(name, age3)

	 // 变量的集合：
	 var(
	 	name2 = "老王"
	 	age4 = 40
	 )
	 fmt.Println(name2, age4)
// 练习1：声明定义几个变量
// 练习2：变量名 := 数值，
}
