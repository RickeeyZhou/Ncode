package main

import "fmt"

func main()  {
	/*
	数组的特点：
	有序(有下标)，可以重复，数据可以被修改，但是长度不能修改(定长)。
	切片：变长，长度可以改变。
	容器：
		len(),cap()
	 */
	 // 创建数组
	 a := [4]int{1,2,3,4}
	 fmt.Println(len(a))
	 fmt.Println("数组的：",len(a),cap(a))
	 // 创建切片1.直接创建
	 s1 := [] int{1,2,3,4,5,6,7,8}
	 fmt.Println(s1)
	 fmt.Println("切片的：",len(s1),cap(s1)) // 8 8
	 fmt.Printf("%T\n", s1) // []int
	 s2 :=[]int{} // 创建了一个空的切片
	 fmt.Println(len(s2), cap(s2)) // 0 0

	 // 2.使用make([]T, len,cap) cap如果不给数值，那么和len一致。
	 s3 := make([]int, 3, 8)
	 fmt.Println(s3) // [0 0 0]
	 fmt.Println(len(s3),cap(s3)) // 3 8

	 //3.在数组的基础上创建切片
	 b := [10]int{1,2,3,4,5,6,7,8,9,10}
	 s4 := b[2:6] // 数组名[start:end)
	 fmt.Println(s4)//[3 4 5 6]
	 fmt.Println(len(s4), cap(s4)) // 4 8
	 s5 := b[:6] // [0:6]，从头开始，切割刀下标6(不包含6)
	 fmt.Println(s5)
	 s6 :=b[6:] // [start:],从start开始，切割刀最后
	 fmt.Println(s6)
	 //s7 := b[4:12] // invalid slice index 12 (out of bounds for 10-element array)
	 //fmt.Println(s7)





}
