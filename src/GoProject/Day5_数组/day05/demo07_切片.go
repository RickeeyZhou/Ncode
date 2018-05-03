package main

import "fmt"

func main()  {
	/*
	切片的类型：
	 */
	 arr1:=[4]int{1,2,3,4}  // 创建一个数组arr1
	 arr2:=arr1 // 创建一个数组arr2，将arr1数据复制一份。
	 fmt.Println(arr1)
	 fmt.Println(arr2)
	 arr1[0] = 100
	 fmt.Println(arr1)
	 fmt.Println(arr2)
	fmt.Println("-----------------")
	 //2.切片的类型：引用类型
	 s1 := [] int{5,6,7,8}
	 s2 := s1
	 fmt.Println(s1)
	 fmt.Println(s2)
	 s1[0]= 100
	 fmt.Println(s1)
	 fmt.Println(s2)
	fmt.Println("-------------")
	 // 3.操作数组
	 arr3 := [5]int{1,2,3,4,5}
	 s3 := arr3[1:4] //[2,3,4]
	 s4 := arr3[3:5] //[4,5]
	fmt.Println(arr3)
	fmt.Println(s3)
	fmt.Println(s4)
	//s3[2] = 100
	arr3[3] = 200
	fmt.Println(arr3)
	fmt.Println(s3)
	fmt.Println(s4)

}
