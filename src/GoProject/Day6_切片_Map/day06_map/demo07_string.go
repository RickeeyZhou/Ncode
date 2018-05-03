package main

import "fmt"

func main()  {
	/*
	字符串：一个字符序列
		序列：有序的一串
		"hello",``
			"abc","ab","a",""
		''
	 */
	 num1:='A' // 单个的字符，使用单引号，类型本质是int32
	 num2:=65
	 num3:="A"//string
	 s1 := "hello国中"
	 fmt.Println(num1,num2,num3)
	 fmt.Printf("%T,%T\n",num1,s1) //int32  string

	 //长度：字符串中的字节的数量。
	 fmt.Println(len(s1))// 5+6  字节

	 //序列：有序-->下标，位置，从0开始
	 s2:="HelloWorld" // 字符串：
	 fmt.Println("打印第一个字母：",s2[0])
	 fmt.Println("打印第二个字母：",s2[1])
	 //遍历字符串：
	 for i:=0;i<len(s2);i++{
	 	//fmt.Println(s2[i])
	 	fmt.Printf("%c\n",s2[i])
	 }
	 /*
	 range string
	 	index,value
	  */
	 for _,char:= range s2{
	 	//fmt.Println(index,char)
	 	//fmt.Printf("%c",char)
	 	fmt.Println(char)
	 }
	//string-->byte[]
	 arr1:=[]byte(s1)
	 fmt.Println(arr1)
	 arr2:=[]byte{65,66,67,68,69,70}
	 fmt.Println(arr2)
	 s3:=string(arr2)
	 fmt.Println(s3)



}
