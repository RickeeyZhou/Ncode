package main

import "fmt"

func main()  {
	/*
	数据类型：基本数据类型，复合数据类型
	基本数据类型：
	1.布尔类型：bool，表示真和假
		取值范围：true，false
	2.数值型：整数，浮点，复数
	整数：
		有符号/无符号，位数，取值范围
		int8/uint8,	   8位   [-128,127]/[0,255]
		int16/uint16,  16位	 [-32768,32767]/[0,65535]
		int32/uint32,  32位  (-2147483648 到 2147483647)/ (0 到 4294967295)
		int64/uint64,  64位   (-9223372036854775808 到 9223372036854775807)/ (0 到 18446744073709551615)
		byte,同uint8
		rune，同int32
		int/uint,同操作系统相关

	浮点型：
		float32/float64, 32/64

	字符串：一个字符序列。
		字符个数可以0个，1个，多个。"", "a", "abc"
		"", ``
	 */
	 //1. 布尔类型
	 var b1 bool = false
	 b2 := true
	 fmt.Println("b1的值：", b1)
	 fmt.Println("b2的值：", b2)
	 fmt.Printf("b2的数值是：%t,类型是：%T\n", b2, b2)
	 // 2.整数类型
	 var num1 int8 = 100
	 fmt.Printf("num1的数值：%d，类型是：%T\n", num1, num1)
	 var num2 uint8 = 200
	 fmt.Printf("num2的数值：%d，类型是：%T\n", num2, num2)

	 var num3 = 400
	 fmt.Printf("num3的数值是：%d,类型是：%T\n", num3, num3)
	 num4 := 'A' // int32
	 fmt.Println(num4)
	 //%v,原型，
	 //%q,数值对应的utf8编码的字符
	 fmt.Printf("num4的数值：%d，%v,%q,%T\n", num4,num4,num4,num4)

	 // 3.浮点类型：
	 var num5 float32 = 3.14
	 var num6 float64 = 8.99
	 fmt.Printf("num5的数值是：%f，类型是：%T,num6的数值是：%f,类型：%T\n", num5,num5,num6,num6)

	 //4. 字符串
	 var s1 string = "memeda"
	 s2 := `helloworld`
	 s3 := 'a' // int32
	 s4 := "a" // string
	 s5 := "ab" // string
	 fmt.Println(s1, s2, s3, s4,s5)
	 fmt.Printf("字符串：%s,类型：%T\n", s2, s2)

	 // 练习1：练习数据类型，每个类型定义变量

}
