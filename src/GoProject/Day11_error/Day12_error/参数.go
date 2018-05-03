package main

import "fmt"

func main(){
	a:=1
	defer fmt.Printf("%d,%p ",a,&a)
	a=100
	fmt.Printf("%d,%p ",a,&a)
}
//func test1(num int){
//	fmt.Println("函数中：num 的值,%p",num,&num)
//}
