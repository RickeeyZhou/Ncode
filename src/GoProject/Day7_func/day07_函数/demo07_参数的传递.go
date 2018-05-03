package main

import "fmt"

func main()  {
	/*
	参数的传递的两种方式：
		值传递：传递的是数据的副本。对原来的数据没有影响
			int,string，float64,bool,数组
		引用传递：传递的是数据的地址。多个引用 指向同一个快内容。
			一个引用修改，那么数据就会改变。
			切片,map
	 */
	a := 1
	b := 2
	fmt.Printf("函数调用前a的值：%d,b的值：%d\n ",a, b) // 1,2
	fun10(a, b)//调用函数
	fmt.Printf("函数调用前a的值：%d,b的值：%d\n ",a, b)//1,2
	fmt.Println("-----------")
	a1 := [4]int {1,2,3,4} // [4]int
	fmt.Printf("%T\n",a1)
	a2:= a1 // 值传递：数组将自己的数据备份一份，传递给a2
	fmt.Println(a1)
	fmt.Println(a2)
	a2[0] = 1000
	fmt.Println(a2)
	fmt.Println(a1)

	fmt.Println("函数调用前：数值的内容：",a1) //[1 2 3 4]
	fun11(a1)
	fmt.Println("函数调用后：数组的内容；",a1)//[1,2,3,4]

	fmt.Println("-------------------")
	s1 :=[]int{1,2,3,4}
	s2 := s1 // 是引用类型，传递是地址，s1和s2指向同一块内存
	s2[0] = 10
	fmt.Println(s1) // [10, 2, 3, 4]
	fmt.Println(s2) //[10, 2, 3, 4]

	fun12(s1)
	fmt.Println(s1)
	fmt.Println(s2)
}
func fun10(m, n int)  {
	fmt.Println(m,n) // 1, 2
	m = 100
	n = 200
	fmt.Printf("函数执行过程中：m的值：%d，n的值：%d\n",m, n)//100,200
}

func fun11(b1 [4]int)  { // b1 = a1,值传递
	fmt.Println(b1) //[1 2 3 4]
	b1[0] = 100
	fmt.Println(b1) //[100 2 3 4]
}

func fun12(s3 []int){ // s3 = s1
	fmt.Println(s3)
	s3[0]= 1000
	fmt.Println(s3)
}
