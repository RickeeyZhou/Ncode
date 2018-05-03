package main

import "fmt"

func main()  {
	var i = 30 // 整数int
	fmt.Printf("整数20,10进制:%d,2进制：%b,8进制：%o,十六进制：%x, 十六进制：%#X  \n", i, i, i, i, i) //


	 var flag1 = true  //定义一个变量表示布尔类型
	fmt.Println("变量flag的值：", flag1)
	fmt.Printf("变量flag的值：%t \n", flag1)
	fmt.Printf("数值是：%t,类型是：%T \n", flag1, flag1)

	var s1 = "思密达"
	fmt.Printf("数值是：%s,类型是：%T", s1, s1)
}
/*
打印格式化：Printf()
Printf("规则%...",x)

%d,整数(默认值10进制)
%b，二进制
%o，八进制
%x，十六进制
%#x，Ox 14


 */
