package main

import "fmt"

func main() {
	/*
	数组：array
		概念：存储一组相同数据类型的数据结构。

	特点：
		有序(下标index)，数据可以重复。
		定长

	语法：
		A：数组的创建
			var variable_name [SIZE] variable_type
			var a [4] int
			var b = [4]int{}
			var c = [4] int {1,2,3,4}
			var d = [4]int{1,2}
			e := [4]int{index:value}
			f:=[...] int{}
		B：数组的访问
			主要通过下标：index，取值范围是：从0开始，到长度减1。
			index：[0,len(arr)-1]
				0,1,2,3,4,5.....
			赋值：
			数组名[index] = value
			取值：
			数组名[index]

		常用的内置函数：
		len(),长度-->容器中存储的元素的个数
		cap(),容量-->容器中存储的最多元素的个数

	注意事项：
		1.一个数组中只能存储同一种类型的数据。
	 */
	//a := 10 //仅能存储一个数据
	//b := 20
	//fmt.Println(a)
	//fmt.Println(b)
	//a = 100
	//fmt.Println(a)
	//1.定义一个数组
	var arr1 [5] int
	fmt.Println(len(arr1)) // 打印数组的长度
	fmt.Println(arr1)      // [0 0 0 0 0]

	//2. 访问数组：向数据中存储数据，从数组中获取数据
	arr1[0] = 1 // 将数值1，存储到数组的第一个位置上
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	fmt.Println(arr1)

	// 3.获取数组中的数据
	fmt.Println("数组中第一个数值：", arr1[0])
	fmt.Println("数组中的第三个数值：", arr1[2])
	fmt.Println("数组中的第四个数值：", arr1[3])
	//fmt.Println(arr1[5]) // invalid array index 5 (out of bounds for 5-element

	//4.数组的遍历：依次访问
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// 5.数组的创建的其他语法格式：
	var a [4] int // 同var a = [4] int{}
	fmt.Println(a)
	// 创建数组的时候，可以同时赋值，赋值的数量<=数组的长度。
	var b = [5] int{6,7,8,9,10}
	fmt.Println(b)
	var c = [5] int{1,2,3,4}
	fmt.Println(c)
	var d = [5] int{4:100}
	fmt.Println(d)
	var e =[5]int{1:3,3:8}
	fmt.Println(e) // 0 3 0 8 0

	f := [4] string{"Rose","王二狗","老王","三胖"}
	fmt.Println(f)

	g := [...]int{1,2,3,4,5} // 可以通过...,根据数据的实际数量，设置数组的长度
	fmt.Println(g)
	fmt.Println(len(g))

	h := [...]int{1:100,4:200,9:300}
	fmt.Println(h)
	fmt.Println(len(h))
	/*
	练习1：创建一个int类型的数组，存储4个数。分别进行赋值，并打印结果。
练习2：创建一个double类型的数组，存储2个小数。
练习3：创建一个String类型的数组，存储3个字符串。
练习4：给定一个数组，arr1 := [10] int{5,4,3,7,10,2,9,8,6,1}，
	求数组中所有数据的总和。
	 */

}
