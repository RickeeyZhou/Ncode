package main

import "fmt"

func main()  {
	/*
利用defer函数，将一个字符串倒序打印
 */
 	i1:='A' //int32
 	fmt.Printf("%T\n", i1)
 	s1:="helloworld" //字节的集合：byte
 	for i:=0;i<len(s1);i++{
 		//fmt.Printf("%T\n",s1[i]) //uint8,byte
 		fmt.Printf("%c",s1[i])
	}
	fmt.Println()
	for _,v :=range s1{
		defer fmt.Printf("%c",v)
		// h, e, l, l .....d
	}
}
