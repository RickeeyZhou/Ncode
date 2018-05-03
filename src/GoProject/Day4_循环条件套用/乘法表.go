package main

import (
	"fmt"

)

func main() {
	/*
	1*1=1
	2*1=2 2*2 =4
	3*1=3 3*2=6 3*3=9
	。。。
	9*1=9 9*2=18.。。9*9=81
	*/
	// 第一层 for 1开始 i++
	// 第二层 for j 1 开始 ，i++
	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Print(i,"*",j,"=",i*j," ")
		}
		fmt.Println()
	}
}
