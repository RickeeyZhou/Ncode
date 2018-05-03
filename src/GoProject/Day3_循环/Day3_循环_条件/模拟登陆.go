package main

import "fmt"

func main() {
	fmt.Println("this is a login Programming")
	var (
		acc string
		Pas string //var acc,Pas string
	)
	fmt.Scanf("%s ,%s",&acc,&Pas)          // scanf  是把空格 当作分隔符的 遇到空格就自动跳装
	fmt.Printf("用户名：%s\n密码：%s\n",acc,Pas)
}
