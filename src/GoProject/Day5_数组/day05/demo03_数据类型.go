package main

import "fmt"

func main()  {
	/*
	数组的类型：
	值类型：传递该数据的时候，传递数据的副本(备份，复制)
	引用类型：
	 */
	 a := [4]int{1,2,3,4}
	 b :=[3] string{"memeda","王二狗","李小花"}
	 fmt.Println(a)
	 fmt.Println(b)
	 fmt.Printf("数组a的的类型：%T\n", a) // [4]int
	 fmt.Printf("数组b的类型：%T\n",b) // [3]string

	 c :=[2]int{1,2}
	 fmt.Printf("数组c的类型：%T\n",c) // [2]int

	d := [2]int{5,6} // [2]int
	c = d

	 num1 := 10
	 num2 := 20
	 //num3 := "haha"
	 num1 = num2 // num1=20,num2=20
	 num2 = 100
	 fmt.Println(num1,num2)

	 //num1 = num3

	 fmt.Println("---------------------")
	 arr1 := [...]string{"Rose", "Jack","王二狗"}
	 fmt.Println(arr1)
	 arr2 := arr1 //  将arr1的数据拷贝一份，值赋值给arr2。
	 fmt.Println(arr2)
	 arr1[0] = "三胖思密达" // 更改了arr1的内容，对于arr2没有影响。因为是值类型，传递的是备份数据。
	 fmt.Println(arr1)
	 fmt.Println(arr2)




}
