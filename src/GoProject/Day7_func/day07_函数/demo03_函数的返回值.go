package main

import "fmt"
/*
函数的返回值：
	一个函数执行后，返回给调用处的结果。

	需要借助于return关键字：
		A：表示将return后的数据，返回给调用处
		B：结束该函数的执行。

	函数名声时需要写清楚：返回值的定义
		func 函数名 (形参列表) 返回值声明{
			//...
			return
		}
 */
func main()  {
	// 需求：1-10的和，乘以2输出
	res1 := 0
	res1 = getSum3() // 相当于res1的值，取自getSum中return后的数据
	fmt.Println(res1 * 2)
	// 理解为res1 = sum
	fmt.Println(getSum4())

	res3,res4:=rect(4.4, 6.5)
	fmt.Println("矩形的面积是：", res3,",矩形的周长是：",res4)
	// 使用空白标识符，舍弃某个数据
	res5,_ := rect2(2.3, 4.5)
	fmt.Println("面积：",res5)

}
//1-10的和
func getSum3()int{
	sum:=0 // 局部变量
	for i:=0;i<=10;i++{
		sum += i
	}
	return sum // return表示将后面的数据，返回给调用处
}
// 也可以这样写：
func getSum4()(sum int){
	for i:=0;i<=10;i++{
		sum += i
	}
	return // 函数声明时，写清楚了返回值变量，那么此处可以省略不写
}

//设计一个函数，用于求一个矩形的面积和周长。
func rect(width, length float64)(area float64,peri float64) {
	area = width * length // 求面积
	peri = 2*( width + length) // 求周长
	return
}
func rect2(wid, len float64)(float64,float64)  {
	area:= wid*len
	peri:=2*(wid + len)
	return area, peri
}
/*
练习1：定义一个函数，接收两个参数，用于比较两个数的大小，并将大的数返
回。
练习2：定义一个函数，用于接收两个参数，求和，并将结果返回。
练习3：定义一个函数，用于求圆的面积和周长，并返回结果。
	面积：PI * R * R
	周长：2 * PI * R
 */
