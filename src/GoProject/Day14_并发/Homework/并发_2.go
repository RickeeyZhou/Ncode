package main

import (
	"fmt"
)
var Date chan bool
func main(){

	go test1()
	go test2()
	<-Date

}
func test1(){
	for i:=1;i<=10000;i++{
		fmt.Println("test1's function i:",i)
	}
	Date<-true
}
func test2(){
	for j:=1;j<=10000;j++{
		fmt.Println("test2 function j:",j)
	}

}

