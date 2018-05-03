package main

import "fmt"

func main()  {
	/*
	切片：
	len(),存储的实际的数量
	cap(),存储的最大的数量，自动扩容：规则：2倍
	append(slice, ele1,ele2,ele3)
	append(slice1,slice2...)
	cap:0-->1-->2-->4-->8-->16-->32
	 */
	 s1:=[]int{}
	 fmt.Println(s1)
	 fmt.Println(len(s1), cap(s1))
	 s1 = append(s1,10)
	 fmt.Println(s1)
	 fmt.Println(len(s1),cap(s1))
	 s1 = append(s1,20,30,40)
	 fmt.Println(s1)
	 fmt.Println(len(s1),cap(s1))
	 for i:=1;i<=10;i++{
	 	s1 = append(s1,i)
	 	fmt.Println("长度：",len(s1),"容量：",cap(s1),"内容：",s1)
	 }
	 arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
	 s2:= arr[:5] //[1,2,3,4,5]
	 s3 := arr[2:7] //[3,4,5,6,7]
	 fmt.Println(arr)
	 fmt.Println(s2)
	 fmt.Println(s3)
	 s2[3] = 100
	 fmt.Println(s2)
	 fmt.Println(s3)
	 fmt.Println(arr)
	 fmt.Println(len(s2), cap(s2)) // 5, 10
	 fmt.Println(len(s3), cap(s3)) // 5, 8
	 fmt.Printf("%p\n", s2) // 0xc04206a0a0
	 s2 = append(s2, 1,1,1,1,1,1,1,1,1)
	 fmt.Printf("%p\n",s2)
	 fmt.Println(s2)
	 fmt.Println(arr)
	 fmt.Println(s3)

	s4:=[]int{1,2,3,4}
	s5 := make([]int ,4,10)
	for i:=0;i<len(s5);i++{
		s5[i] = i*2+1
	}
	fmt.Println(s4)
	fmt.Println(s5)

	s4 = append(s4, s5...)
	fmt.Println(s4)

	//实现删除效果：
	index :=3 // 数组/切片[start:end]
	s4 = append(s4[:index],s4[index+1:]...)
	fmt.Println(s4)





}
