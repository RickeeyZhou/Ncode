package main

import (
	"fmt"
	"strings"
)

func main(){
	pathName:="http://192.168.10.1:8080/Day33_Servlet/aa.jpeg"
	fmt.Println(strings.Split(pathName,"/"))
	 s :=[]string(strings.Split(pathName,"/"))
	 for index,value:=range s{
	 	fmt.Println(index,value)
	 }
	 fmt.Println(s[4])
	 //s1:=[]int{1,1}
	 //fmt.Println(s)
	 //fmt.Printf("%T",s[3])
	 //fmt.Println(s[4])
	// fmt.Println(s[3])
	//var arr1 []byte
	// arr1 =[]byte(pathName)
	// fmt.Println(arr1)

// 枚举类型 用{} 参数 函数就用（）
//数组到数组 直接 s:=arr1
// 从切片到数组只有 string（s1）
//
a:="aekjffjkJDJ294384848DKFJFJkdjfhfh2943845593nfnJRIEIFJ" //string

//s:=str
a2:=a[:]
fmt.Println(a2)
var H,h,num int
for _,value :=range a2{
	if value>=48&&value<=57{
		num++
	}
	if value >=65&&value<=90{
		H++
	}
	if value>=97&&value<=122{
		h++
	}
}
fmt.Println(H,h,num)
}
