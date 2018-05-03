package main

import (
	"fmt"
)

func main(){
	defer func(){
		if msg:=recover();msg!=nil{
			fmt.Println(msg)
		}
	}()
	testA()
	testB()
	testC()
	fmt.Println("main....over...")
}
func testA() {
	fmt.Println("函数A")
}
func testB(){
	for i:=1;i<=10;i++{
		if i==5{
			panic ("testb hanshu,panic...")
		}
		fmt.Println("函数b",i)
	}

}
func testC(){

	panic("函数c panic ...")
	fmt.Println("函数c（）。。。。。")
}