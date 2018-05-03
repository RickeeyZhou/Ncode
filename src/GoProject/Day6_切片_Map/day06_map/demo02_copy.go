package main

import "fmt"

func main()  {
	/*
	copy(dst,src)
	 */
	 s1:=[]int{1,2,3,4,5}
	 s2:=[]int{7,8,9}
	 fmt.Println(s1)
	 fmt.Println(s2)
	 //copy(s1,s2) //将s2拷贝到s1中
	 //copy(s2, s1) // 将s1中的数据，拷贝到s2中。
	 //copy(s2, s1[2:4])// 将s1中3-4 的数据，拷贝到s2中
	 copy(s2[1:] , s1[2:4])
	 fmt.Println(s1)
	 fmt.Println(s2)
}
